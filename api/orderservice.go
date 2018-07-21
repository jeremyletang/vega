package api

import (
	"context"
	"vega/blockchain"
	"vega/core"
	"vega/datastore"
	"vega/msg"
)

type OrderService interface {
	Init(vega *core.Vega, orderStore datastore.OrderStore)
	CreateOrder(ctx context.Context, order *msg.Order) (success bool, err error)
	CancelOrder(ctx context.Context, order *msg.Order) (success bool, err error)
	GetByMarket(ctx context.Context, market string, limit uint64) (orders []*msg.Order, err error)
	GetByParty(ctx context.Context, party string, limit uint64) (orders []*msg.Order, err error)
	GetByMarketAndId(ctx context.Context, market string, id string) (order *msg.Order, err error)
	GetByPartyAndId(ctx context.Context, market string, id string) (order *msg.Order, err error)
	GetMarkets(ctx context.Context) ([]string, error)
	GetOrderBookDepth(ctx context.Context, market string) (orderBookDepth *msg.OrderBookDepth, err error)
}

type orderService struct {
	app        *core.Vega
	orderStore datastore.OrderStore
	blockchain blockchain.Client
}

func NewOrderService() OrderService {
	return &orderService{}
}

func (p *orderService) Init(app *core.Vega, orderStore datastore.OrderStore) {
	p.app = app
	p.orderStore = orderStore
	p.blockchain = blockchain.NewClient()
}

func (p *orderService) CreateOrder(ctx context.Context, order *msg.Order) (success bool, err error) {
	order.Remaining = order.Size
	// TODO validate
	// Call out to the blockchain package/layer and use internal client to gain consensus
	return p.blockchain.CreateOrder(ctx, order)
}

func (p *orderService) CancelOrder(ctx context.Context, order *msg.Order) (success bool, err error) {
	// Cancel by ID, market, other fields not required
	// TODO validate
	return p.blockchain.CancelOrder(ctx, order)
}

func (p *orderService) GetByMarket(ctx context.Context, market string, limit uint64) (orders []*msg.Order, err error) {
	o, err := p.orderStore.GetByMarket(market, datastore.GetParams{Limit: limit})
	if err != nil {
		return nil, err
	}
	result := make([]*msg.Order, 0)
	for _, order := range o {
		//if order.Remaining == 0 {
		//	continue
		//}
		o := &msg.Order{
			Id:        order.Id,
			Market:    order.Market,
			Party:     order.Party,
			Side:      order.Side,
			Price:     order.Price,
			Size:      order.Timestamp,
			Remaining: order.Remaining,
			Timestamp: order.Timestamp,
			Type:      order.Type,
		}
		result = append(result, o)
	}
	return result, err
}

func (p *orderService) GetByParty(ctx context.Context, party string, limit uint64) (orders []*msg.Order, err error) {
	o, err := p.orderStore.GetByParty(party, datastore.GetParams{Limit: limit})
	if err != nil {
		return nil, err
	}
	result := make([]*msg.Order, 0)
	for _, order := range o {
		//if order.Remaining == 0 {
		//	continue
		//}
		o := &msg.Order{
			Id:        order.Id,
			Market:    order.Market,
			Party:     order.Party,
			Side:      order.Side,
			Price:     order.Price,
			Size:      order.Timestamp,
			Remaining: order.Remaining,
			Timestamp: order.Timestamp,
			Type:      order.Type,
		}
		result = append(result, o)
	}
	return result, err
}

func (p *orderService) GetByMarketAndId(ctx context.Context, market string, id string) (order *msg.Order, err error) {
	o, err := p.orderStore.GetByMarketAndId(market, id)
	if err != nil {
		return &msg.Order{}, err
	}
	return o.ToProtoMessage(), err
}

func (p *orderService) GetByPartyAndId(ctx context.Context, market string, id string) (order *msg.Order, err error) {
	o, err := p.orderStore.GetByPartyAndId(market, id)
	if err != nil {
		return &msg.Order{}, err
	}
	return o.ToProtoMessage(), err
}

func (p *orderService) GetMarkets(ctx context.Context) ([]string, error) {
	markets, err := p.orderStore.GetMarkets()
	if err != nil {
		return []string{}, err
	}
	return markets, err
}

func (p *orderService) GetOrderBookDepth(ctx context.Context, marketName string) (orderBookDepth *msg.OrderBookDepth, err error) {
	return p.orderStore.GetOrderBookDepth(marketName)
}

//
//func getClient() (*rpc.Client, error) {
//	mux.Lock()
//	if len(clients) == 0 {
//		mux.Unlock()
//		client := rpc.Client{
//		}
//		if err := client.Connect(); err != nil {
//			return nil, err
//		}
//		return &client, nil
//	}
//	client := clients[0]
//	clients = clients[1:]
//	mux.Unlock()
//	return client, nil
//}
//
//func releaseClient(c *rpc.Client) {
//	mux.Lock()
//	clients = append(clients, c)
//	mux.Unlock()
//}
