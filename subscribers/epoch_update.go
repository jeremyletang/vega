package subscribers

import (
	"context"

	"code.vegaprotocol.io/data-node/logging"
	types "code.vegaprotocol.io/protos/vega"
	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
	"code.vegaprotocol.io/vega/events"
)

type EpochUpdateEvent interface {
	events.Event
	Proto() eventspb.EpochEvent
}

type EpochStore interface {
	AddEpoch(seq uint64, startTime int64, endTime int64)
	AddDelegation(types.Delegation)
}

type EpochUpdateSub struct {
	*Base

	epochStore EpochStore

	log *logging.Logger
}

func NewEpochUpdateSub(ctx context.Context, epochStore EpochStore, log *logging.Logger, ack bool) *EpochUpdateSub {
	sub := &EpochUpdateSub{
		Base:       NewBase(ctx, 10, ack),
		epochStore: epochStore,
		log:        log,
	}

	if sub.isRunning() {
		go sub.loop(ctx)
	}

	return sub
}

func (vu *EpochUpdateSub) loop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			vu.Halt()
			return
		case e := <-vu.ch:
			if vu.isRunning() {
				vu.Push(e...)
			}
		}
	}
}

func (vu *EpochUpdateSub) Push(evts ...events.Event) {
	if len(evts) == 0 {
		return
	}

	for _, e := range evts {
		switch et := e.(type) {
		case EpochUpdateEvent:
			eu := et.Proto()
			vu.epochStore.AddEpoch(eu.GetSeq(), eu.GetStartTime(), eu.GetEndTime())
		default:
			vu.log.Panic("Unknown event type in candles subscriber", logging.String("Type", et.Type().String()))
		}
	}
}

func (vu *EpochUpdateSub) Types() []events.Type {
	return []events.Type{
		events.EpochUpdate,
	}
}
