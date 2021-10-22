package staking

import (
	"context"

	"code.vegaprotocol.io/vega/libs/crypto"
	"code.vegaprotocol.io/vega/types"

	"github.com/golang/protobuf/proto"
)

var (
	depositedKey = (&types.PayloadStakeVerifierDeposited{}).Key()
	removedKey   = (&types.PayloadStakeVerifierRemoved{}).Key()

	hashKeys = []string{
		depositedKey,
		removedKey,
	}
)

type stakeVerifierSnapshotState struct {
	hash       map[string][]byte
	serialised map[string][]byte
	changed    map[string]bool
}

func (s *StakeVerifier) serialisePendingSD() ([]byte, error) {
	deposited := make([]*types.StakeDeposited, 0, len(s.pendingSDs))

	for _, p := range s.pendingSDs {
		deposited = append(deposited, p.StakeDeposited)
	}

	pl := types.Payload{
		Data: &types.PayloadStakeVerifierDeposited{
			StakeVerifierDeposited: deposited,
		},
	}
	return proto.Marshal(pl.IntoProto())
}

func (s *StakeVerifier) serialisePendingSR() ([]byte, error) {
	removed := make([]*types.StakeRemoved, 0, len(s.pendingSRs))

	for _, p := range s.pendingSRs {
		removed = append(removed, p.StakeRemoved)
	}

	pl := types.Payload{
		Data: &types.PayloadStakeVerifierRemoved{
			StakeVerifierRemoved: removed,
		},
	}
	return proto.Marshal(pl.IntoProto())
}

// get the serialised form and hash of the given key.
func (s *StakeVerifier) getSerialisedAndHash(k string) ([]byte, []byte, error) {
	if _, ok := s.keyToSerialiser[k]; !ok {
		return nil, nil, types.ErrSnapshotKeyDoesNotExist
	}

	if !s.svss.changed[k] {
		return s.svss.serialised[k], s.svss.hash[k], nil
	}

	data, err := s.keyToSerialiser[k]()
	if err != nil {
		return nil, nil, err
	}

	hash := crypto.Hash(data)
	s.svss.serialised[k] = data
	s.svss.hash[k] = hash
	s.svss.changed[k] = false
	return data, hash, nil
}

func (s *StakeVerifier) Namespace() types.SnapshotNamespace {
	return types.StakeVerifierSnapshot
}

func (s *StakeVerifier) Keys() []string {
	return hashKeys
}

func (s *StakeVerifier) GetHash(k string) ([]byte, error) {
	_, hash, err := s.getSerialisedAndHash(k)
	return hash, err
}

func (s *StakeVerifier) GetState(k string) ([]byte, error) {
	data, _, err := s.getSerialisedAndHash(k)
	return data, err
}

func (s *StakeVerifier) LoadState(ctx context.Context, payload *types.Payload) ([]types.StateProvider, error) {
	if s.Namespace() != payload.Data.Namespace() {
		return nil, types.ErrInvalidSnapshotNamespace
	}

	switch pl := payload.Data.(type) {
	case *types.PayloadStakeVerifierDeposited:
		return nil, s.restorePendingSD(ctx, pl.StakeVerifierDeposited)
	case *types.PayloadStakeVerifierRemoved:
		return nil, s.restorePendingSR(ctx, pl.StakeVerifierRemoved)
	default:
		return nil, types.ErrUnknownSnapshotType
	}
}

func (s *StakeVerifier) restorePendingSD(ctx context.Context, deposited []*types.StakeDeposited) error {
	s.pendingSDs = make([]*pendingSD, 0, len(deposited))

	for _, d := range deposited {
		s.ensureNotDuplicate(d.ID, d.Hash())

		pending := &pendingSD{
			StakeDeposited: d,
			check:          func() error { return s.ocv.CheckStakeDeposited(d) },
		}

		s.pendingSDs = append(s.pendingSDs, pending)
		s.witness.RestoreResource(pending, s.onEventVerified)
	}
	s.svss.changed[depositedKey] = true
	return nil
}

func (s *StakeVerifier) restorePendingSR(ctx context.Context, removed []*types.StakeRemoved) error {
	s.pendingSRs = make([]*pendingSR, 0, len(removed))

	for _, r := range removed {
		s.ensureNotDuplicate(r.ID, r.Hash())

		pending := &pendingSR{
			StakeRemoved: r,
			check:        func() error { return s.ocv.CheckStakeRemoved(r) },
		}

		s.pendingSRs = append(s.pendingSRs, pending)
		s.witness.RestoreResource(pending, s.onEventVerified)
	}

	s.svss.changed[removedKey] = true
	return nil
}
