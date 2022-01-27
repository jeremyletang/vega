package events

import (
	"context"

	proto "code.vegaprotocol.io/protos/vega"
	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
)

type NetworkLimits struct {
	*Base
	nl *proto.NetworkLimits
}

func NewNetworkLimitsEvent(ctx context.Context, limits *proto.NetworkLimits) *NetworkLimits {
	limitsCopy := limits.DeepClone()
	return &NetworkLimits{
		Base: newBase(ctx, NetworkLimitsEvent),
		nl:   limitsCopy,
	}
}

func (n *NetworkLimits) NetworkLimits() *proto.NetworkLimits {
	return n.nl
}

func (n NetworkLimits) Proto() *proto.NetworkLimits {
	return n.nl
}

func (n NetworkLimits) StreamMessage() *eventspb.BusEvent {
	busEvent := newBusEventFromBase(n.Base)
	busEvent.Event = &eventspb.BusEvent_NetworkLimits{
		NetworkLimits: n.nl,
	}

	return busEvent
}

func NetworkLimitsEventFromStream(ctx context.Context, be *eventspb.BusEvent) *NetworkLimits {
	return &NetworkLimits{
		Base: newBaseFromBusEvent(ctx, NetworkLimitsEvent, be),
		nl:   be.GetNetworkLimits(),
	}
}
