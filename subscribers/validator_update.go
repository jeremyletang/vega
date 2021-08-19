package subscribers

import (
	"context"

	"code.vegaprotocol.io/data-node/logging"
	types "code.vegaprotocol.io/protos/vega"
	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
	"code.vegaprotocol.io/vega/events"
)

type ValidatorUpdateEvent interface {
	events.Event
	Proto() eventspb.ValidatorUpdate
}

type NodeStore interface {
	AddNode(types.Node)
	AddDelegation(types.Delegation)
}

type ValidatorUpdateSub struct {
	*Base
	nodeStore NodeStore

	log *logging.Logger
}

func NewValidatorUpdateSub(ctx context.Context, nodeStore NodeStore, log *logging.Logger, ack bool) *ValidatorUpdateSub {
	sub := &ValidatorUpdateSub{
		Base:      NewBase(ctx, 10, ack),
		nodeStore: nodeStore,
		log:       log,
	}

	if sub.isRunning() {
		go sub.loop(ctx)
	}

	return sub
}

func (vu *ValidatorUpdateSub) loop(ctx context.Context) {
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

func (vu *ValidatorUpdateSub) Push(evts ...events.Event) {
	if len(evts) == 0 {
		return
	}

	for _, e := range evts {
		switch et := e.(type) {
		case ValidatorUpdateEvent:
			vue := et.Proto()

			vu.nodeStore.AddNode(types.Node{
				// @TODO - add Id
				// Id: vue.GetInfoUrl(),
				PubKey:   vue.GetTmPubKey(),
				InfoUrl:  vue.GetInfoUrl(),
				Location: vue.GetCountry(),
				Status:   types.NodeStatus_NODE_STATUS_NON_VALIDATOR,
			})
		default:
			vu.log.Panic("Unknown event type in candles subscriber", logging.String("Type", et.Type().String()))
		}
	}
}

func (vu *ValidatorUpdateSub) Types() []events.Type {
	return []events.Type{
		events.ValidatorUpdateEvent,
	}
}
