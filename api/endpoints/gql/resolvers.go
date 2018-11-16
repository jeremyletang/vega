package gql

import (
	"context"
	"vega/api"
	"vega/msg"
	"errors"
	"strconv"
	"time"
	"vega/log"
	"fmt"
	"vega/filters"
	"vega/vegatime"
)

type resolverRoot struct {
	orderService api.OrderService
	tradeService api.TradeService
	candleService api.CandleService
}

func NewResolverRoot(orderService api.OrderService, tradeService api.TradeService) *resolverRoot {
	return &resolverRoot{
		orderService: orderService,
		tradeService: tradeService,
	}
}

func (r *resolverRoot) Query() QueryResolver {
	return (*MyQueryResolver)(r)
}
func (r *resolverRoot) Mutation() MutationResolver {
	return (*MyMutationResolver)(r)
}
func (r *resolverRoot) Candle() CandleResolver {
	return (*MyCandleResolver)(r)
}
func (r *resolverRoot) MarketDepth() MarketDepthResolver {
	return (*MyMarketDepthResolver)(r)
}
func (r *resolverRoot) PriceLevel() PriceLevelResolver {
	return (*MyPriceLevelResolver)(r)
}
func (r *resolverRoot) Market() MarketResolver {
	return (*MyMarketResolver)(r)
}
func (r *resolverRoot) Order() OrderResolver {
	return (*MyOrderResolver)(r)
}
func (r *resolverRoot) Trade() TradeResolver {
	return (*MyTradeResolver)(r)
}
func (r *resolverRoot) Vega() VegaResolver {
	return (*MyVegaResolver)(r)
}
func (r *resolverRoot) Position() PositionResolver {
	return (*MyPositionResolver)(r)
}
func (r *resolverRoot) Party() PartyResolver {
	return (*MyPartyResolver)(r)
}
func (r *resolverRoot) Subscription() SubscriptionResolver {
	return (*MySubscriptionResolver)(r)
}



// BEGIN: Query Resolver

type MyQueryResolver resolverRoot

func (r *MyQueryResolver) Vega(ctx context.Context) (Vega, error) {
	var vega = Vega{}
	return vega, nil
}

// END: Query Resolver



// BEGIN: Root Resolver

type MyVegaResolver resolverRoot

func (r *MyVegaResolver) Markets(ctx context.Context, obj *Vega, name *string) ([]Market, error) {
	if name == nil {
		return nil, errors.New("all markets on VEGA query not implemented")
	}

	// Todo(cdm): MarketStore --> check if market exists via dedicated marketstore...
	//existing, err := r.orderService.GetMarkets(ctx)
	//if err != nil {
	//	return nil, err
	//}

	var existing = []string{"BTC/DEC18"}
	found := false
	for _, m := range existing {
		if *name == m {
		   found = true
		   break
		} 
	}
	if !found {
		return nil, errors.New(fmt.Sprintf("market %s does not exist", *name))
	}
	
	var markets = make([]Market, 0)
	var market = Market{
		Name: *name,
	}
	markets = append(markets, market)
	
	return markets, nil
}

func (r *MyVegaResolver) Parties(ctx context.Context, obj *Vega, name *string) ([]Party, error) {
	if name == nil {
		return nil, errors.New("all parties on VEGA query not implemented")
	}

	// Todo(cdm): PartyStore --> check if party exists...
	var parties = make([]Party, 0)
	var party = Party{
		Name: *name,
	}
	parties = append(parties, party)
	
	return parties, nil
}

// END: Root Resolver


// BEGIN: Market Resolver

type MyMarketResolver resolverRoot

func (r *MyMarketResolver) Orders(ctx context.Context, market *Market,
	where *OrderFilter, skip *int, first *int, last *int) ([]msg.Order, error) {

	queryFilters, err := buildOrderQueryFilters(where, skip, first, last)
	if err != nil {
		return nil, err
	}
	orders, err := r.orderService.GetByMarket(ctx, market.Name, queryFilters)
	if err != nil {
		return nil, err
	}
	valOrders := make([]msg.Order, 0)
	for _, v := range orders {
		valOrders = append(valOrders, *v)
	}
	return valOrders, nil
}

func (r *MyMarketResolver) Trades(ctx context.Context, market *Market,
	where *TradeFilter, skip *int, first *int, last *int) ([]msg.Trade, error) {

	queryFilters, err := buildTradeQueryFilters(where, skip, first, last)
	if err != nil {
		return nil, err
	}
	trades, err := r.tradeService.GetByMarket(ctx, market.Name, queryFilters)
	if err != nil {
		return nil, err
	}
	valTrades := make([]msg.Trade, 0)
	for _, v := range trades {
		valTrades = append(valTrades, *v)
	}
	return valTrades, nil
}

func (r *MyMarketResolver) Depth(ctx context.Context, market *Market) (msg.MarketDepth, error) {

	// Look for market depth for the given market (will validate market internally)
	// FYI: Market depth is also known as OrderBook depth within the matching-engine
	depth, err := r.orderService.GetMarketDepth(ctx, market.Name)
	if err != nil {
		return msg.MarketDepth{}, err
	}

	return *depth, nil
}

func (r *MyMarketResolver) Candles(ctx context.Context, market *Market,
	last int, interval int) ([]msg.Candle, error) {

	defaultInterval := "1m"
	defaultTime := time.Now()
	candles, err := r.candleService.GetCandles(ctx, market.Name, defaultTime, defaultInterval)
	if err != nil {
		return nil, err
	}

	valCandles := make([]msg.Candle, 0)
	for _, v := range candles {
		valCandles = append(valCandles, *v)
	}

	return valCandles, nil
}

// END: Market Resolver


// BEGIN: Party Resolver

type MyPartyResolver resolverRoot

func (r *MyPartyResolver) Orders(ctx context.Context, party *Party,
	where *OrderFilter, skip *int, first *int, last *int) ([]msg.Order, error) {

	queryFilters, err := buildOrderQueryFilters(where, skip, first, last)
	if err != nil {
		return nil, err
	}
	orders, err := r.orderService.GetByParty(ctx, party.Name, queryFilters)
	if err != nil {
		return nil, err
	}
	valOrders := make([]msg.Order, 0)
	for _, v := range orders {
		valOrders = append(valOrders, *v)
	}
	return valOrders, nil
}

func (r *MyPartyResolver) Positions(ctx context.Context, obj *Party) ([]msg.MarketPosition, error) {
	positions, err := r.tradeService.GetPositionsByParty(ctx, obj.Name)
	if err != nil {
		return nil, err
	}
	var valPositions = make([]msg.MarketPosition, 0)
	for _, v := range positions {
		valPositions = append(valPositions, *v)
	}
	return valPositions, nil
}

// END: Party Resolver


// BEGIN: Market Depth Resolver

type MyMarketDepthResolver resolverRoot

func (r *MyMarketDepthResolver) Buy(ctx context.Context, obj *msg.MarketDepth) ([]msg.PriceLevel, error) {
	valBuyLevels := make([]msg.PriceLevel, 0)
	for _, v := range obj.Buy {
		valBuyLevels = append(valBuyLevels, *v)
	}
	return valBuyLevels, nil
}
func (r *MyMarketDepthResolver) Sell(ctx context.Context, obj *msg.MarketDepth) ([]msg.PriceLevel, error) {
	valBuyLevels := make([]msg.PriceLevel, 0)
	for _, v := range obj.Sell {
		valBuyLevels = append(valBuyLevels, *v)
	}
	return valBuyLevels, nil
}

func (r *MyMarketDepthResolver) LastTrade(ctx context.Context, obj *msg.MarketDepth) (*msg.Trade, error) {
	queryFilters := &filters.TradeQueryFilters{}
	last := uint64(1)
	queryFilters.Last = &last
	trades, err := r.tradeService.GetByMarket(ctx, obj.Name, queryFilters)
	if err != nil {
		return nil, err
	}
	if trades != nil && len(trades) > 0 && trades[0] != nil {
		return trades[0], nil
	}
	// No trades found on the market yet (and no errors)
	// this can happen at the beginning of a new market
	return nil, nil
}

// END: Market Depth Resolver

// BEGIN: Order Resolver

type MyOrderResolver resolverRoot

func (r *MyOrderResolver) Price(ctx context.Context, obj *msg.Order) (string, error) {
	return strconv.FormatUint(obj.Price, 10), nil
}
func (r *MyOrderResolver) Type(ctx context.Context, obj *msg.Order) (OrderType, error) {
	return OrderType(obj.Type.String()), nil
}
func (r *MyOrderResolver) Side(ctx context.Context, obj *msg.Order) (Side, error) {
	return Side(obj.Side.String()), nil
}
func (r *MyOrderResolver) Market(ctx context.Context, obj *msg.Order) (Market, error) {
	return Market {
		Name: obj.Market,
	}, nil
}
func (r *MyOrderResolver) Size(ctx context.Context, obj *msg.Order) (string, error) {
	return strconv.FormatUint(obj.Size, 10), nil
}
func (r *MyOrderResolver) Remaining(ctx context.Context, obj *msg.Order) (string, error) {
	return strconv.FormatUint(obj.Remaining, 10), nil
}
func (r *MyOrderResolver) Timestamp(ctx context.Context, obj *msg.Order) (string, error) {
	return strconv.FormatUint(obj.Timestamp, 10), nil
}
func (r *MyOrderResolver) Status(ctx context.Context, obj *msg.Order) (OrderStatus, error) {
	return OrderStatus(obj.Status.String()), nil
}
func (r *MyOrderResolver) Datetime(ctx context.Context, obj *msg.Order) (string, error) {
	vegaTimestamp := vegatime.Stamp(obj.Timestamp)
	return vegaTimestamp.Rfc3339Nano(), nil
}

// END: Order Resolver

// BEGIN: Trade Resolver

type MyTradeResolver resolverRoot

func (r *MyTradeResolver) Market(ctx context.Context, obj *msg.Trade) (Market, error) {
	return Market{Name: obj.Market}, nil
}
func (r *MyTradeResolver) Aggressor(ctx context.Context, obj *msg.Trade) (Side, error) {
	return Side(obj.Aggressor.String()), nil
}
func (r *MyTradeResolver) Price(ctx context.Context, obj *msg.Trade) (string, error) {
	return strconv.FormatUint(obj.Price, 10), nil
}
func (r *MyTradeResolver) Size(ctx context.Context, obj *msg.Trade) (string, error) {
	return strconv.FormatUint(obj.Size, 10), nil
}
func (r *MyTradeResolver) Timestamp(ctx context.Context, obj *msg.Trade) (string, error) {
	return strconv.FormatUint(obj.Timestamp, 10), nil
}
func (r *MyTradeResolver) Datetime(ctx context.Context, obj *msg.Trade) (string, error) {
	vegaTimestamp := vegatime.Stamp(obj.Timestamp)
	return vegaTimestamp.Rfc3339Nano(), nil
}

// END: Trade Resolver

// BEGIN: Candle Resolver

type MyCandleResolver resolverRoot

func (r *MyCandleResolver) High(ctx context.Context, obj *msg.Candle) (string, error) {
	return strconv.FormatUint(obj.High, 10), nil
}
func (r *MyCandleResolver) Low(ctx context.Context, obj *msg.Candle) (string, error) {
	return strconv.FormatUint(obj.Low, 10), nil
}
func (r *MyCandleResolver) Open(ctx context.Context, obj *msg.Candle) (string, error) {
	return strconv.FormatUint(obj.Open, 10), nil
}
func (r *MyCandleResolver) Close(ctx context.Context, obj *msg.Candle) (string, error) {
	return strconv.FormatUint(obj.Close, 10), nil
}
func (r *MyCandleResolver) Volume(ctx context.Context, obj *msg.Candle) (string, error) {
	return strconv.FormatUint(obj.Volume, 10), nil
}
func (r *MyCandleResolver) Datetime(ctx context.Context, obj *msg.Candle) (string, error) {
	return vegatime.Stamp(obj.Timestamp).Rfc3339Nano(), nil
}
func (r *MyCandleResolver) Timestamp(ctx context.Context, obj *msg.Candle) (string, error) {
	return strconv.FormatUint(obj.Timestamp, 10), nil
}
// END: Candle Resolver

// BEGIN: Price Level Resolver

type MyPriceLevelResolver resolverRoot

func (r *MyPriceLevelResolver) Price(ctx context.Context, obj *msg.PriceLevel) (string, error) {
	return strconv.FormatUint(obj.Price, 10), nil
}

func (r *MyPriceLevelResolver) Volume(ctx context.Context, obj *msg.PriceLevel) (string, error) {
	return strconv.FormatUint(obj.Volume, 10), nil
}

func (r *MyPriceLevelResolver) NumberOfOrders(ctx context.Context, obj *msg.PriceLevel) (string, error) {
	return strconv.FormatUint(obj.Price, 10), nil
}

func (r *MyPriceLevelResolver) CumulativeVolume(ctx context.Context, obj *msg.PriceLevel) (string, error) {
	return strconv.FormatUint(obj.CumulativeVolume, 10), nil
}

// END: Price Level Resolver


// BEGIN: Position Resolver

type MyPositionResolver resolverRoot

func (r *MyPositionResolver) RealisedVolume(ctx context.Context, obj *msg.MarketPosition) (string, error) {
	return strconv.FormatInt(obj.RealisedVolume, 10), nil
}

func (r *MyPositionResolver) RealisedProfitValue(ctx context.Context, obj *msg.MarketPosition) (string, error) {
	return r.absInt64Str(obj.RealisedPNL), nil
}

func (r *MyPositionResolver) RealisedProfitDirection(ctx context.Context, obj *msg.MarketPosition) (ValueDirection, error) {
	return r.direction(obj.RealisedPNL), nil
}

func (r *MyPositionResolver) UnrealisedVolume(ctx context.Context, obj *msg.MarketPosition) (string, error) {
	return strconv.FormatInt(obj.UnrealisedVolume, 10), nil
}

func (r *MyPositionResolver) UnrealisedProfitValue(ctx context.Context, obj *msg.MarketPosition) (string, error) {
	return r.absInt64Str(obj.UnrealisedPNL), nil
}

func (r *MyPositionResolver) UnrealisedProfitDirection(ctx context.Context, obj *msg.MarketPosition) (ValueDirection, error) {
	return r.direction(obj.UnrealisedPNL), nil
}

func (r *MyPositionResolver) AverageEntryPrice(ctx context.Context, obj *msg.MarketPosition) (string, error)  {
	return strconv.FormatUint(obj.AverageEntryPrice, 10), nil
}

func (r *MyPositionResolver) MinimumMargin(ctx context.Context, obj *msg.MarketPosition) (string, error)  {
	return strconv.FormatInt(obj.MinimumMargin, 10), nil
}

func (r *MyPositionResolver) Market(ctx context.Context, obj *msg.MarketPosition) (Market, error) {
	return Market{Name: obj.Market}, nil
}

func (r *MyPositionResolver) absInt64Str(val int64) (string) {
	if val < 0 {
		return strconv.FormatInt(val * -1, 10)
	}
	return strconv.FormatInt(val, 10)
}

func (r *MyPositionResolver) direction(val int64) (ValueDirection) {
	if val < 0 {
		return ValueDirectionNegative
	}
	return ValueDirectionPositive
}

// END: Position Resolver


// BEGIN: Mutation Resolver

type MyMutationResolver resolverRoot

func (r *MyMutationResolver) OrderCreate(ctx context.Context, market string, party string, price string,
	size string, side Side, type_ OrderType, expiration *string) (PreConsensus, error) {
	order := &msg.Order{}
	res := PreConsensus{}

	// We need to convert strings to uint64 (JS doesn't yet support uint64)
	p, err := safeStringUint64(price)
	if err != nil {
		return res, err
	}
	order.Price = p
	s, err := safeStringUint64(size)
	if err != nil {
		return res, err
	}
	order.Size = s
	if len(market) == 0 {
		return res, errors.New("market missing or empty")
	}
	order.Market = market
	if len(party) == 0 {
		return res, errors.New("party missing or empty")
	}
	order.Party = party
	order.Type, err = parseOrderType(&type_)
	if err != nil {
		return res, err
	}
	order.Side, err = parseSide(&side)
	if err != nil {
		return res, err
	}

	// GTT must have an expiration value
	if order.Type == msg.Order_GTT && expiration != nil {

		// Validate RFC3339 with no milli or nanosecond (@matt has chosen this strategy, good enough until unix epoch timestamp)
		layout := "2006-01-02T15:04:05Z"
		_, err := time.Parse(layout, *expiration)
		if err != nil {
			return res, errors.New(fmt.Sprintf("cannot parse expiration time: %s - invalid format sent to create order (example: 2018-01-02T15:04:05Z)", *expiration))
		}

		// move to pure timestamps or convert an RFC format shortly
		order.ExpirationDatetime = *expiration
	}

	// Pass the order over for consensus (service layer will use RPC client internally and handle errors etc)
	accepted, reference, err := r.orderService.CreateOrder(ctx, order)
	if err != nil {
		log.Errorf("error creating order via rpc client and graph-ql", err)
		return res, err
	}

	res.Accepted = accepted
	res.Reference = reference
	return res, nil
}

func (r *MyMutationResolver) OrderCancel(ctx context.Context, id string, market string, party string) (PreConsensus, error) {
	order := &msg.Order{}
	res := PreConsensus{}

	// Cancellation currently only requires ID and Market to be set, all other fields will be added
	if len(market) == 0 {
		return res, errors.New("market missing or empty")
	}
	order.Market = market
	if len(id) == 0 {
		return res, errors.New("id missing or empty")
	}
	order.Id = id
	if len(party) == 0 {
		return res, errors.New("party missing or empty")
	}
	order.Party = party

	// Pass the cancellation over for consensus (service layer will use RPC client internally and handle errors etc)
	accepted, err := r.orderService.CancelOrder(ctx, order)
	if err != nil {
		return res, err
	}

	res.Accepted = accepted
	return res, nil
}

// END: Mutation Resolver


// BEGIN: Subscription Resolver

type MySubscriptionResolver resolverRoot

func (r *MySubscriptionResolver) Orders(ctx context.Context, market *string, party *string) (<-chan []msg.Order, error) {
	// Validate market, and todo future Party (when party store exists)
	err := r.validateMarket(ctx, market)
	if err != nil {
		return nil, err
	}
	c, ref := r.orderService.ObserveOrders(ctx, market, party)
	log.Debugf("GraphQL Orders -> New subscriber: %d", ref)
	return c, nil
}

func (r *MySubscriptionResolver) Trades(ctx context.Context, market *string, party *string) (<-chan []msg.Trade, error) {
	// Validate market, and todo future Party (when party store exists)
	err := r.validateMarket(ctx, market)
	if err != nil {
		return nil, err
	}
	c, ref := r.tradeService.ObserveTrades(ctx, market, party)
	log.Debugf("GraphQL Trades -> New subscriber: %d", ref)
	return c, nil
}

func (r *MySubscriptionResolver) Positions(ctx context.Context, party string) (<-chan msg.MarketPosition, error) {
	c, ref := r.tradeService.ObservePositions(ctx, party)
	log.Debugf("GraphQL Positions -> New subscriber: %d", ref)
	return c, nil
}

func (r *MySubscriptionResolver) MarketDepth(ctx context.Context, market string) (<-chan msg.MarketDepth, error) {
	// Validate market
	err := r.validateMarket(ctx, &market)
	if err != nil {
		return nil, err
	}
	c, ref := r.orderService.ObserveMarketDepth(ctx, market)
	log.Debugf("GraphQL Market Depth -> New subscriber: %d", ref)
	return c, nil
}

func (r *MySubscriptionResolver) Candles(ctx context.Context, market string, interval int) (<-chan msg.Candle, error) {
	// Validate market
	err := r.validateMarket(ctx, &market)
	if err != nil {
		return nil, err
	}

	// Validate interval
	defaultInterval := "1m"

	c, ref := r.candleService.ObserveCandles(ctx, &market, &defaultInterval)
	log.Debugf("GraphQL Candle Interval %s -> New subscriber: %d", defaultInterval, ref)
	return c, nil
}

func (r *MySubscriptionResolver) validateMarket(ctx context.Context, market *string) error {
	// todo(cdm): change this when we have a marketservice/marketstore
	if market != nil {
		if len(*market) == 0 {
			return errors.New("market must not be empty")
		}
		//markets, err := r.orderService.GetMarkets(ctx)
		//if err != nil {
		//	return err
		//}
		var markets = []string{"BTC/DEC18"}
		// Scan all markets for a match
		found := false
		for _, v := range markets {
			if v == *market {
				found = true
				break
			}
		}
		if !found {
			return errors.New(fmt.Sprintf("market %s not found", *market))
		}
	}
	return nil
}

// END: Subscription Resolver


