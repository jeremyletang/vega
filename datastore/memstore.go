package datastore

import (
	"fmt"
)

// memMarket should keep track of the trades/orders operating on a Market.
type memMarket struct {
	name   string
	ordersIndex []string
	orders map[string]*memOrder
	trades map[string]*memTrade
}

// In memory order struct keeps an internal map of pointers to trades for an order.
type memOrder struct {
	order  *Order
	trades []*memTrade
}

// memOrderStore should implement OrderStore interface.
type memOrderStore struct {
	store *MemStore
}

// In memory trade struct keeps a pointer to the related order.
type memTrade struct {
	trade *Trade
	order *memOrder
}

// memTradeStore should implement TradeStore interface.
type memTradeStore struct {
	store *MemStore
}

// MemStore is a RAM based top level structure to hold information about all markets.
// It is initialised by calling NewMemStore with a list of markets.
type MemStore struct {
	markets map[string]*memMarket
}

// NewMemStore creates an instance of the ram based data store.
// This store is simply backed by maps/slices for trades and orders.
func NewMemStore(markets []string) MemStore {
	memMarkets := make(map[string]*memMarket, len(markets))
	for _, name := range markets {
		memMarket := memMarket{
			name:   name,
			orders: map[string]*memOrder{},
			trades: map[string]*memTrade{},
		}
		memMarkets[name] = &memMarket
	}
	return MemStore{
		markets: memMarkets,
	}
}

// NewTradeStore initialises a new TradeStore backed by a MemStore.
func NewTradeStore(ms *MemStore) TradeStore {
	return &memTradeStore{store: ms}
}

// NewTradeStore initialises a new OrderStore backed by a MemStore.
func NewOrderStore(ms *MemStore) OrderStore {
	return &memOrderStore{store: ms}
}

// Helper function to check if a market exists within the memory store.
func (ms *MemStore) marketExists(market string) bool {
	if _, exists := ms.markets[market]; exists {
		return true
	}
	return false
}

// GetAll retrieves a orders for a given market.
func (t *memOrderStore) GetAll(market string, params GetParams) ([]*Order, error) {
	err := t.marketExists(market)
	if err != nil {
		return nil, err
	}

	// Limit is by default descending
	pos := uint64(0)
	orders := make([]*Order, 0)
	for i := len(t.store.markets[market].ordersIndex)-1; i >= 0; i-- {
		if params.Limit > 0 && pos == params.Limit {
			break
		}
		idx := t.store.markets[market].ordersIndex[i]
		value :=t.store.markets[market].orders[idx]
		orders = append(orders, value.order)
		pos++
	}
	return orders, nil
}

// Get retrieves an order for a given market and id.
func (t *memOrderStore) Get(market string, id string) (*Order, error) {
	err := t.marketExists(market)
	if err != nil {
		return nil, err
	}
	v, ok := t.store.markets[market].orders[id]
	if !ok {
		return nil, NotFoundError{fmt.Errorf("could not find id %s", id)}
	}
	return v.order, nil
}

// Post creates a new order in the memory store.
func (t *memOrderStore) Post(or *Order) error {
	// todo validation of incoming order
	//	if err := or.Validate(); err != nil {
	//		return fmt.Errorf("cannot store record: %s", err)
	//	}
	err := t.marketExists(or.Market)
	if err != nil {
		return err
	}
	if _, exists := t.store.markets[or.Market].orders[or.Id]; exists {
		return fmt.Errorf("order exists in memstore: %s", or.Id)
	} else {
		fmt.Println("Adding new order with ID ", or.Id)
		order := &memOrder{
			trades: make([]*memTrade, 0),
			order:  or,
		}
		// Insert order struct into lookup hashtable
		t.store.markets[or.Market].orders[or.Id] = order
		// Due to go randomisation of keys, we'll need to add an index entry too for ordering
		t.store.markets[or.Market].ordersIndex = append(t.store.markets[or.Market].ordersIndex, or.Id)
	}
	return nil
}

// Put updates an existing order in the memory store.
func (t *memOrderStore) Put(or *Order) error {
	// todo validation of incoming order
	//	if err := or.Validate(); err != nil {
	//		return fmt.Errorf("cannot store record: %s", err)
	//	}
	err := t.marketExists(or.Market)
	if err != nil {
		return err
	}
	if _, exists := t.store.markets[or.Market].orders[or.Id]; exists {
		fmt.Println("Updating order with ID ", or.Id)
		t.store.markets[or.Market].orders[or.Id].order = or
	} else {
		return fmt.Errorf("order not found in memstore: %s", or.Id)
	}
	return nil
}

// Delete removes an order from the memory store.
func (t *memOrderStore) Delete(or *Order) error {
	err := t.marketExists(or.Market)
	if err != nil {
		return err
	}

	// Remove value
	delete(t.store.markets[or.Market].orders, or.Id)

	// Remove index
	pos := uint64(0)
	for p, value := range t.store.markets[or.Market].ordersIndex {
		if value == or.Id {
			pos = uint64(p)
			break
		}
	}
	t.store.markets[or.Market].ordersIndex = append(t.store.markets[or.Market].ordersIndex[:pos], t.store.markets[or.Market].ordersIndex[pos+1:]...)
	return nil
}

// Checks to see if we have a market on the related memory store with given identifier.
// Returns an error if the market cannot be found and nil otherwise.
func (t *memOrderStore) marketExists(market string) error {
	if !t.store.marketExists(market) {
		return NotFoundError{fmt.Errorf("could not find market %s", market)}
	}
	return nil
}

// GetAll retrieves all trades for a given market.
func (t *memTradeStore) GetAll(market string, params GetParams) ([]*Trade, error) {
	err := t.marketExists(market)
	if err != nil {
		return nil, err
	}
	pos := uint64(0)
	trades := make([]*Trade, 0)
	for _, value := range t.store.markets[market].trades {
		trades = append(trades, value.trade)
		if params.Limit > 0 && pos == params.Limit {
			break
		}
		pos++
	}
	return trades, nil
}

// Get retrieves a trade for a given id.
func (t *memTradeStore) Get(market string, id string) (*Trade, error) {
	err := t.marketExists(market)
	if err != nil {
		return nil, err
	}
	v, ok := t.store.markets[market].trades[id]
	if !ok {
		return nil, NotFoundError{fmt.Errorf("could not find id %s", id)}
	}
	return v.trade, nil
}


// GetByOrderId retrieves all trades for a given order id.
func (t *memTradeStore) GetByOrderId(market string, orderId string, params GetParams) ([]*Trade, error) {
	err := t.marketExists(market)
	if err != nil {
		return nil, err
	}
	order := t.store.markets[market].orders[orderId]
	if order == nil {
		return nil, fmt.Errorf("order not found in memstore: %s", orderId)
	} else {
		pos := uint64(0)
		trades := make([]*Trade, 0)
		for _, v := range order.trades {
			trades = append(trades, v.trade)
			if params.Limit > 0 && pos == params.Limit {
				break
			}
			pos++
		}
		return trades, nil
	}
}

// Post creates a new trade in the memory store.
func (t *memTradeStore) Post(tr *Trade) error {
	//todo validation of incoming trade
	// if err := tr.Validate(); err != nil {
	//		return fmt.Errorf("cannot store record: %s", err)
	//	}
	err := t.marketExists(tr.Market)
	if err != nil {
		return err
	}
	if o, exists := t.store.markets[tr.Market].orders[tr.OrderId]; exists {
		trade := &memTrade{
			trade: tr,
			order: o,
		}
		if _, exists := t.store.markets[tr.Market].trades[tr.Id]; exists {
			return fmt.Errorf("trade exists in memstore: %s", tr.Id)
		} else {
			// Map new trade to memstore and append trade to order
			t.store.markets[tr.Market].trades[tr.Id] = trade
			o.trades = append(o.trades, trade)
		}
		return nil
	} else {
		return fmt.Errorf("related order for trade not found in memstore: %s", tr.OrderId)
	}
}

// Put updates an existing trade in the store.
func (t *memTradeStore) Put(tr *Trade) error {
	//todo validation of incoming trade
	// if err := tr.Validate(); err != nil {
	//		return fmt.Errorf("cannot store record: %s", err)
	//	}
	err := t.marketExists(tr.Market)
	if err != nil {
		return err
	}
	if o, exists := t.store.markets[tr.Market].orders[tr.OrderId]; exists {
		trade := &memTrade{
			trade: tr,
			order: o,
		}
		if _, exists := t.store.markets[tr.Market].trades[tr.Id]; exists {
			// Perform the update
			t.store.markets[tr.Market].trades[tr.Id] = trade
		} else {
			return fmt.Errorf("trade not found in memstore: %s", tr.Id)
		}
		//o.trades = append(o.trades, trade)
		return nil
	} else {
		return fmt.Errorf("related order for trade not found in memstore: %s", tr.OrderId)
	}
}

// Removes an order from the store.
func (t *memTradeStore) Delete(tr *Trade) error {
	err := t.marketExists(tr.Market)
	if err != nil {
		return err
	}
	delete(t.store.markets[tr.Market].trades, tr.Id)
	return nil
}

// Checks to see if we have a market on the related memory store with given identifier.
// Returns an error if the market cannot be found and nil otherwise.
func (t *memTradeStore) marketExists(market string) error {
	if !t.store.marketExists(market) {
		return NotFoundError{fmt.Errorf("could not find market %s", market)}
	}
	return nil
}