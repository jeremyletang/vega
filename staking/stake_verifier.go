package staking

import (
	"context"
	"errors"
	"time"

	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/validators"
)

var (
	ErrNoStakeDepositedEventFound   = errors.New("no stake deposited event found")
	ErrNoStakeRemovedEventFound     = errors.New("no stake removed event found")
	ErrNotAnEthereumConfig          = errors.New("not an ethereum config")
	ErrMissingConfirmations         = errors.New("missing confirmations")
	ErrInvalidStakeRemovedEventID   = errors.New("invalid stake removed event ID")
	ErrInvalidStakeDepositedEventID = errors.New("invalid stake deposited event ID")
)

// Witness provide foreign chain resources validations
//go:generate go run github.com/golang/mock/mockgen -destination mocks/time_ticker_mock.go -package mocks code.vegaprotocol.io/vega/staking TimeTicker
type TimeTicker interface {
	NotifyOnTick(func(context.Context, time.Time))
}

//go:generate go run github.com/golang/mock/mockgen -destination mocks/eth_confirmations_mock.go -package mocks code.vegaprotocol.io/vega/staking EthConfirmations
type EthConfirmations interface {
	Check(uint64) error
}

//go:generate go run github.com/golang/mock/mockgen -destination mocks/eth_on_chain_verifier_mock.go -package mocks code.vegaprotocol.io/vega/staking EthOnChainVerifier
type EthOnChainVerifier interface {
	CheckStakeDeposited(*types.StakeDeposited) error
	CheckStakeRemoved(*types.StakeRemoved) error
}

// Witness provide foreign chain resources validations
//go:generate go run github.com/golang/mock/mockgen -destination mocks/witness_mock.go -package mocks code.vegaprotocol.io/vega/staking Witness
type Witness interface {
	StartCheck(validators.Resource, func(interface{}, bool), time.Time) error
}

type StakeVerifier struct {
	log *logging.Logger
	cfg Config

	accs    *Accounting
	witness Witness
	broker  Broker

	ocv EthOnChainVerifier

	currentTime time.Time

	pendingSDs      []*pendingSD
	pendingSRs      []*pendingSR
	finalizedEvents []*types.StakingEvent
}

type pendingSD struct {
	*types.StakeDeposited
	check func() error
}

func (p pendingSD) GetID() string { return p.ID }
func (p *pendingSD) Check() error { return p.check() }

type pendingSR struct {
	*types.StakeRemoved
	check func() error
}

func (p pendingSR) GetID() string { return p.ID }
func (p *pendingSR) Check() error { return p.check() }

func NewStakeVerifier(
	log *logging.Logger,
	cfg Config,
	accs *Accounting,
	tt TimeTicker,
	witness Witness,
	broker Broker,
	onChainVerifier EthOnChainVerifier,
) (sv *StakeVerifier) {
	defer func() {
		tt.NotifyOnTick(sv.onTick)
	}()

	return &StakeVerifier{
		log:     log,
		cfg:     cfg,
		accs:    accs,
		witness: witness,
		ocv:     onChainVerifier,
	}
}

func (s *StakeVerifier) ProcessStakeRemoved(
	ctx context.Context, event *types.StakeRemoved) error {
	pending := &pendingSR{
		StakeRemoved: event,
		check:        func() error { return s.ocv.CheckStakeRemoved(event) },
	}

	s.pendingSRs = append(s.pendingSRs, pending)
	evt := pending.IntoStakingEvent()
	evt.Status = types.StakingEventStatusPending
	s.broker.Send(events.NewStakingEvent(ctx, *evt))

	return s.witness.StartCheck(
		pending, s.onEventVerified, s.currentTime.Add(2*time.Hour))
}

func (s *StakeVerifier) ProcessStakeDeposited(
	ctx context.Context, event *types.StakeDeposited) error {
	pending := &pendingSD{
		StakeDeposited: event,
		check:          func() error { return s.ocv.CheckStakeDeposited(event) },
	}

	s.pendingSDs = append(s.pendingSDs, pending)

	evt := pending.IntoStakingEvent()
	evt.Status = types.StakingEventStatusPending
	s.broker.Send(events.NewStakingEvent(ctx, *evt))

	return s.witness.StartCheck(
		pending, s.onEventVerified, s.currentTime.Add(2*time.Hour))
}

func (s *StakeVerifier) removePendingStakeDeposited(id string) error {
	for i, v := range s.pendingSDs {
		if v.ID == id {
			s.pendingSDs = s.pendingSDs[:i+copy(s.pendingSDs[i:], s.pendingSDs[i+1:])]
			return nil
		}
	}
	return ErrInvalidStakeDepositedEventID
}

func (s *StakeVerifier) removePendingStakeRemoved(id string) error {
	for i, v := range s.pendingSRs {
		if v.ID == id {
			s.pendingSRs = s.pendingSRs[:i+copy(s.pendingSRs[i:], s.pendingSRs[i+1:])]
			return nil
		}
	}
	return ErrInvalidStakeRemovedEventID
}

func (s *StakeVerifier) onEventVerified(event interface{}, ok bool) {
	var evt *types.StakingEvent
	switch pending := event.(type) {
	case *pendingSD:
		evt = pending.IntoStakingEvent()
		if err := s.removePendingStakeDeposited(evt.ID); err != nil {
			s.log.Error("could not remove pending stake deposited event", logging.Error(err))
		}
	case *pendingSR:
		evt = pending.IntoStakingEvent()
		if err := s.removePendingStakeRemoved(evt.ID); err != nil {
			s.log.Error("could not remove pending stake removed event", logging.Error(err))
		}
	default:
		s.log.Error("stake verifier received invalid event")
		return
	}

	evt.Status = types.StakingEventStatusRejected
	if ok {
		evt.Status = types.StakingEventStatusAccepted
	}
	evt.FinalizedAt = s.currentTime.UnixNano()
	s.finalizedEvents = append(s.finalizedEvents, evt)
}

func (s *StakeVerifier) onTick(ctx context.Context, t time.Time) {
	s.currentTime = t

	for _, evt := range s.finalizedEvents {
		if evt.Status == types.StakingEventStatusAccepted {
			s.accs.AddEvent(ctx, evt)
		}
		s.broker.Send(events.NewStakingEvent(ctx, *evt))
	}

	s.finalizedEvents = nil
}
