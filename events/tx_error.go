package events

import (
	"context"

	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
	"code.vegaprotocol.io/vega/types"
)

type TxErr struct {
	*Base
	evt *eventspb.TxErrorEvent
}

func NewTxErrEvent(ctx context.Context, err error, partyID string, tx interface{}) *TxErr {
	evt := &TxErr{
		Base: newBase(ctx, TxErrEvent),
		evt: &eventspb.TxErrorEvent{
			PartyId: partyID,
			ErrMsg:  err.Error(),
		},
	}
	switch tv := tx.(type) {
	case *types.ProposalSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_Proposal{
			Proposal: tv.IntoProto(),
		}
	case types.ProposalSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_Proposal{
			Proposal: tv.IntoProto(),
		}
	case *types.VoteSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_VoteSubmission{
			VoteSubmission: tv.IntoProto(),
		}
	case types.VoteSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_VoteSubmission{
			VoteSubmission: tv.IntoProto(),
		}
	case *types.OrderSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_OrderSubmission{
			OrderSubmission: tv.IntoProto(),
		}
	case types.OrderSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_OrderSubmission{
			OrderSubmission: tv.IntoProto(),
		}
	case *types.OrderCancellation:
		evt.evt.Transaction = &eventspb.TxErrorEvent_OrderCancellation{
			OrderCancellation: tv.IntoProto(),
		}
	case types.OrderCancellation:
		evt.evt.Transaction = &eventspb.TxErrorEvent_OrderCancellation{
			OrderCancellation: tv.IntoProto(),
		}
	case *types.OrderAmendment:
		evt.evt.Transaction = &eventspb.TxErrorEvent_OrderAmendment{
			OrderAmendment: tv.IntoProto(),
		}
	case types.OrderAmendment:
		evt.evt.Transaction = &eventspb.TxErrorEvent_OrderAmendment{
			OrderAmendment: tv.IntoProto(),
		}
	case *types.LiquidityProvisionSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_LiquidityProvisionSubmission{
			LiquidityProvisionSubmission: tv.IntoProto(),
		}
	case types.LiquidityProvisionSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_LiquidityProvisionSubmission{
			LiquidityProvisionSubmission: tv.IntoProto(),
		}
	case *types.WithdrawSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_WithdrawSubmission{
			WithdrawSubmission: tv.IntoProto(),
		}
	case types.WithdrawSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_WithdrawSubmission{
			WithdrawSubmission: tv.IntoProto(),
		}
	case *commandspb.DelegateSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_DelegateSubmission{
			DelegateSubmission: tv,
		}
	case commandspb.DelegateSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_DelegateSubmission{
			DelegateSubmission: &tv,
		}
	case *commandspb.UndelegateSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_UndelegateSubmission{
			UndelegateSubmission: tv,
		}
	case commandspb.UndelegateSubmission:
		evt.evt.Transaction = &eventspb.TxErrorEvent_UndelegateSubmission{
			UndelegateSubmission: &tv,
		}
	case *commandspb.RestoreSnapshot:
		evt.evt.Transaction = &eventspb.TxErrorEvent_RestoreSnapshot{
			RestoreSnapshot: tv,
		}
	case commandspb.RestoreSnapshot:
		evt.evt.Transaction = &eventspb.TxErrorEvent_RestoreSnapshot{
			RestoreSnapshot: &tv,
		}
	}
	return evt
}

func (t TxErr) IsParty(id string) bool {
	return t.evt.PartyId == id
}

func (t TxErr) Proto() eventspb.TxErrorEvent {
	return *t.evt
}

func (t TxErr) StreamMessage() *eventspb.BusEvent {
	return &eventspb.BusEvent{
		Id:    t.eventID(),
		Block: t.TraceID(),
		Type:  t.et.ToProto(),
		Event: &eventspb.BusEvent_TxErrEvent{
			TxErrEvent: t.evt,
		},
	}
}

func TxErrEventFromStream(ctx context.Context, be *eventspb.BusEvent) *TxErr {
	return &TxErr{
		Base: newBaseFromStream(ctx, TxErrEvent, be),
		evt:  be.GetTxErrEvent(),
	}
}
