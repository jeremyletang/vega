package orders_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"code.vegaprotocol.io/vega/internal/logging"
	"code.vegaprotocol.io/vega/internal/orders"
	"code.vegaprotocol.io/vega/internal/orders/mocks"
	"code.vegaprotocol.io/vega/internal/vegatime"
	types "code.vegaprotocol.io/vega/proto"

	"github.com/golang/mock/gomock"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var (
	orderSubmission = types.OrderSubmission{
		Id:       "order_id",
		MarketId: "market_id",
		Party:    "party",
		Price:    10000,
		Size:     1,
		Side:     types.Side(1),
		Type:     types.Order_GTT,
	}
)

type testService struct {
	ctrl       *gomock.Controller
	orderStore *mocks.MockOrderStore
	timeSvc    *mocks.MockTimeService
	block      *mocks.MockBlockchain
	svc        *orders.Svc
}

type orderMatcher struct {
	e types.Order
}

func TestCreateOrder(t *testing.T) {
	t.Run("Create order - successful", testOrderSuccess)
	t.Run("Create order - expired", testOrderExpired)
	t.Run("Create order - blockchain error", testOrderBlockchainError)
}

func testOrderSuccess(t *testing.T) {
	// now
	now := vegatime.Now()
	// expires 2 hours from now
	expires := now.Add(time.Hour * 2)
	pre := &types.PreConsensusOrder{
		Accepted:  true,
		Reference: "order_reference",
	}
	order := orderSubmission
	order.ExpiresAt = expires.UnixNano()
	matcher := orderMatcher{
		e: types.Order{
			Id:        order.Id,
			Market:    order.MarketId,
			Party:     order.Party,
			Price:     order.Price,
			Size:      order.Size,
			Side:      order.Side,
			Type:      order.Type,
			ExpiresAt: expires.UnixNano(),
		},
	}
	svc := getTestService(t)
	defer svc.ctrl.Finish()

	svc.timeSvc.EXPECT().GetTimeNow().Times(1).Return(now, nil)
	svc.block.EXPECT().CreateOrder(gomock.Any(), matcher).Times(1).Return(pre, nil)
	res, err := svc.svc.CreateOrder(context.Background(), &order)
	assert.True(t, res.Accepted)
	assert.NoError(t, err)
	assert.Equal(t, pre.Reference, res.Reference)
}

func testOrderExpired(t *testing.T) {
	// now
	now := vegatime.Now()
	//expired 2 hours ago
	// expires := now.Add(time.Hour * -2)
	order := orderSubmission
	svc := getTestService(t)
	defer svc.ctrl.Finish()
	svc.timeSvc.EXPECT().GetTimeNow().Times(1).Return(now, nil)
	res, err := svc.svc.CreateOrder(context.Background(), &order)
	assert.False(t, res.Accepted)
	assert.Error(t, err)
	assert.Equal(t, "", res.Reference)
}

func testOrderBlockchainError(t *testing.T) {
	// now
	now := vegatime.Now()
	// expires 2 hours from now
	expires := now.Add(time.Hour * 2)
	bcErr := errors.New("blockchain error")
	order := orderSubmission
	order.ExpiresAt = expires.UnixNano()
	matcher := orderMatcher{
		e: types.Order{
			Id:        order.Id,
			Market:    order.MarketId,
			Party:     order.Party,
			Price:     order.Price,
			Size:      order.Size,
			Side:      order.Side,
			Type:      order.Type,
			ExpiresAt: expires.UnixNano(),
		},
	}
	svc := getTestService(t)
	defer svc.ctrl.Finish()
	svc.timeSvc.EXPECT().GetTimeNow().Times(1).Return(now, nil)
	svc.block.EXPECT().CreateOrder(gomock.Any(), matcher).Times(1).Return(&types.PreConsensusOrder{}, bcErr)
	res, err := svc.svc.CreateOrder(context.Background(), &order)
	assert.False(t, res.Accepted)
	assert.Error(t, err)
	assert.Equal(t, bcErr, err)
	assert.Equal(t, "", res.Reference)
}

func getTestService(t *testing.T) *testService {
	log := logging.NewTestLogger()
	ctrl := gomock.NewController(t)
	orderStore := mocks.NewMockOrderStore(ctrl)
	timeSvc := mocks.NewMockTimeService(ctrl)
	block := mocks.NewMockBlockchain(ctrl)
	conf := orders.NewDefaultConfig(log)
	svc, err := orders.NewService(conf, orderStore, timeSvc, block)
	if err != nil {
		t.Fatalf("Failed to get test service: %+v", err)
	}
	return &testService{
		ctrl:       ctrl,
		orderStore: orderStore,
		timeSvc:    timeSvc,
		block:      block,
		svc:        svc,
	}
}

func (m orderMatcher) String() string {
	return fmt.Sprintf("%#v", m.e)
}

func (m orderMatcher) Matches(x interface{}) bool {
	var v types.Order
	switch val := x.(type) {
	case *types.Order:
		v = *val
	case types.Order:
		v = val
	default:
		return false
	}
	if m.e.Id != v.Id && m.e.Market != v.Market {
		return false
	}
	if m.e.Party != v.Party {
		return false
	}

	return (m.e.ExpiresAt == v.ExpiresAt)
}
