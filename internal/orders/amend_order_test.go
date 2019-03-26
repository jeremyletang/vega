package orders

import (
	"context"
	"fmt"
	"testing"
	"time"

	"code.vegaprotocol.io/vega/internal/vegatime"
	"code.vegaprotocol.io/vega/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	amend = proto.OrderAmendment{
		Id:    "order_id",
		Party: "party",
		Price: 10000,
		Size:  1,
	}
)

type amendMatcher struct {
	e proto.OrderAmendment
}

func TestAmendOrder(t *testing.T) {
	t.Run("Amend order - success", testAmendOrderSuccess)
	t.Run("Amend order - expired", testAmendOrderExpired)
	t.Run("Amend order - not active", testAmendOrderNotActive)
	t.Run("Amend order - invalid payload", testAmendOrderInvalidPayload)
}

func testAmendOrderSuccess(t *testing.T) {
	now := time.Now()
	expires := now.Add(2 * time.Hour)
	arg := amend
	arg.ExpirationDatetime = expires.Format(time.RFC3339)
	arg.ExpirationTimestamp = uint64(time.Duration(expires.Unix()) * time.Second)
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	order := proto.Order{
		Id:     arg.Id,
		Market: "market",
		Party:  arg.Party,
		Status: proto.Order_Active,
	}
	svc.orderStore.EXPECT().GetByPartyAndId(gomock.Any(), arg.Party, arg.Id).Times(1).Return(&order, nil)
	svc.timeSvc.EXPECT().GetTimeNow().Times(1).Return(vegatime.Stamp(now.UnixNano()), now, nil)
	svc.block.EXPECT().AmendOrder(gomock.Any(), amendMatcher{e: arg}).Times(1).Return(true, nil)

	success, err := svc.svc.AmendOrder(context.Background(), &arg)
	assert.True(t, success)
	assert.NoError(t, err)
}

func testAmendOrderExpired(t *testing.T) {
	now := time.Now()
	expires := now.Add(-2 * time.Hour)
	arg := amend
	arg.ExpirationDatetime = expires.Format(time.RFC3339)
	arg.ExpirationTimestamp = uint64(time.Duration(expires.Unix()) * time.Second)
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	order := proto.Order{
		Id:     arg.Id,
		Market: "market",
		Party:  arg.Party,
		Status: proto.Order_Active,
	}
	svc.orderStore.EXPECT().GetByPartyAndId(gomock.Any(), arg.Party, arg.Id).Times(1).Return(&order, nil)
	svc.timeSvc.EXPECT().GetTimeNow().Times(1).Return(vegatime.Stamp(now.UnixNano()), now, nil)

	success, err := svc.svc.AmendOrder(context.Background(), &arg)
	assert.False(t, success)
	assert.Error(t, err)
}

func testAmendOrderNotActive(t *testing.T) {
	now := time.Now()
	expires := now.Add(2 * time.Hour)
	arg := amend
	arg.ExpirationDatetime = expires.Format(time.RFC3339)
	arg.ExpirationTimestamp = uint64(time.Duration(expires.Unix()) * time.Second)
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	order := proto.Order{
		Id:     arg.Id,
		Market: "market",
		Party:  arg.Party,
		Status: proto.Order_Expired,
	}
	svc.orderStore.EXPECT().GetByPartyAndId(gomock.Any(), arg.Party, arg.Id).Times(1).Return(&order, nil)

	success, err := svc.svc.AmendOrder(context.Background(), &arg)
	assert.False(t, success)
	assert.Error(t, err)
}

func testAmendOrderInvalidPayload(t *testing.T) {
	arg := amend
	arg.Size = 0
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	success, err := svc.svc.AmendOrder(context.Background(), &arg)
	assert.False(t, success)
	assert.Error(t, err)
}

func (m amendMatcher) String() string {
	return fmt.Sprintf("%#v", m.e)
}

func (m amendMatcher) Matches(x interface{}) bool {
	var v proto.OrderAmendment
	switch val := x.(type) {
	case *proto.OrderAmendment:
		v = *val
	case proto.OrderAmendment:
		v = val
	default:
		return false
	}
	return (m.e.Id == v.Id && m.e.Party == v.Party && m.e.Price == v.Price && m.e.Size == v.Size &&
		m.e.ExpirationDatetime == v.ExpirationDatetime && m.e.ExpirationTimestamp == v.ExpirationTimestamp)
}
