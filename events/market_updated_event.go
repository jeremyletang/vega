package events

import (
	"context"
	"fmt"

	types "code.vegaprotocol.io/vega/proto"
)

type MarketUpdated struct {
	*Base
	m types.Market
}

func NewMarketUpdatedEvent(ctx context.Context, m types.Market) *MarketUpdated {
	cpy := m.DeepClone()
	return &MarketUpdated{
		Base: newBase(ctx, MarketUpdatedEvent),
		m:    *cpy,
	}
}

// MarketEvent -> is needs to be logged as a market event
func (m MarketUpdated) MarketEvent() string {
	return fmt.Sprintf("Market ID %s updated (%s)", m.m.Id, m.m.String())
}

func (m MarketUpdated) MarketID() string {
	return m.m.Id
}

func (m MarketUpdated) Market() types.Market {
	return m.m
}

func (m MarketUpdated) Proto() types.Market {
	return m.m
}

func (m MarketUpdated) MarketProto() types.MarketEvent {
	return types.MarketEvent{
		MarketId: m.m.Id,
		Payload:  m.MarketEvent(),
	}
}

func (m MarketUpdated) StreamMessage() *types.BusEvent {
	p := m.MarketProto()
	return &types.BusEvent{
		Id:    m.eventID(),
		Block: m.TraceID(),
		Type:  m.et.ToProto(),
		Event: &types.BusEvent_Market{
			Market: &p,
		},
	}
}

func (m MarketUpdated) StreamMarketMessage() *types.BusEvent {
	return m.StreamMessage()
}
