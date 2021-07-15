package events

import (
	"context"

	"code.vegaprotocol.io/data-node/proto"
	eventspb "code.vegaprotocol.io/data-node/proto/events/v1"
)

type Party struct {
	*Base
	p proto.Party
}

func NewPartyEvent(ctx context.Context, p proto.Party) *Party {
	cpy := p.DeepClone()
	return &Party{
		Base: newBase(ctx, PartyEvent),
		p:    *cpy,
	}
}

func (p Party) IsParty(id string) bool {
	return p.p.Id == id
}

func (p *Party) Party() proto.Party {
	return p.p
}

func (p Party) Proto() proto.Party {
	return p.p
}

func (p Party) StreamMessage() *eventspb.BusEvent {
	return &eventspb.BusEvent{
		Id:    p.eventID(),
		Block: p.TraceID(),
		Type:  p.et.ToProto(),
		Event: &eventspb.BusEvent_Party{
			Party: &p.p,
		},
	}
}
