package storage

import (
	"context"
	"fmt"
	"sync"
	"time"

	"code.vegaprotocol.io/vega/logging"
	types "code.vegaprotocol.io/vega/proto"
)

// NoopOrder is a package internal data struct that implements the OrderStore interface.
type NoopOrder struct {
	Config

	cfgMu        sync.Mutex
	log          *logging.Logger
	subscribers  map[uint64]chan<- []types.Order
	subscriberID uint64
	depth        map[string]*Depth
	mu           sync.Mutex
}

func NewNoopOrders(log *logging.Logger, c Config) *NoopOrder {
	log = log.Named(namedLogger)
	log.SetLevel(c.Level.Get())

	return &NoopOrder{
		log:         log,
		Config:      c,
		depth:       map[string]*Depth{},
		subscribers: map[uint64]chan<- []types.Order{},
	}
}

// ReloadConf reloads the config, watches for a changed loglevel.
func (os *NoopOrder) ReloadConf(cfg Config) {
	os.log.Info("reloading configuration")
	if os.log.GetLevel() != cfg.Level.Get() {
		os.log.Info("updating log level",
			logging.String("old", os.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		os.log.SetLevel(cfg.Level.Get())
	}

	os.cfgMu.Lock()
	os.Config = cfg
	os.cfgMu.Unlock()
}

// Subscribe to a channel of new or updated orders. The subscriber id will be returned as a uint64 value
// and must be retained for future reference and to unsubscribe.
func (os *NoopOrder) Subscribe(orders chan<- []types.Order) uint64 {
	os.mu.Lock()
	defer os.mu.Unlock()

	os.subscriberID = os.subscriberID + 1
	os.subscribers[os.subscriberID] = orders

	os.log.Debug("Orders subscriber added in order store",
		logging.Uint64("subscriber-id", os.subscriberID))

	return os.subscriberID
}

func (os *NoopOrder) SaveBatch(accs []types.Order) error {
	return nil
}

func (os *NoopOrder) Unsubscribe(id uint64) error {
	os.mu.Lock()
	defer os.mu.Unlock()

	if len(os.subscribers) == 0 {
		os.log.Debug("Un-subscribe called in order store, no subscribers connected",
			logging.Uint64("subscriber-id", id))
		return nil
	}

	if _, exists := os.subscribers[id]; exists {
		delete(os.subscribers, id)
		os.log.Debug("Un-subscribe called in order store, subscriber removed",
			logging.Uint64("subscriber-id", id))
		return nil
	}

	return fmt.Errorf("subscriber to Orders store does not exist with id: %d", id)
}

func (os *NoopOrder) Post(order types.Order) error {
	return nil
}

func (os *NoopOrder) Put(order types.Order) error {
	return nil
}

func (os *NoopOrder) Commit() (err error) {
	return
}

func (os *NoopOrder) Close() error {
	return nil
}

func (os *NoopOrder) GetByMarket(ctx context.Context, market string, skip,
	limit uint64, descending bool, open bool) ([]*types.Order, error) {
	return []*types.Order{}, nil
}

func (os *NoopOrder) GetByMarketAndID(ctx context.Context, market string, id string) (*types.Order, error) {
	var order types.Order
	return &order, nil
}

func (os *NoopOrder) GetByParty(ctx context.Context, party string, skip uint64,
	limit uint64, descending bool, open bool) ([]*types.Order, error) {

	return []*types.Order{}, nil
}

func (os *NoopOrder) GetByPartyAndID(ctx context.Context, party string, id string) (*types.Order, error) {
	var order types.Order
	return &order, nil
}

func (os *NoopOrder) GetAllVersionsByOrderID(ctx context.Context, id string,
	skip, limit uint64, descending bool) (orders []*types.Order, err error) {

	return []*types.Order{}, nil
}

func (os *NoopOrder) GetByReference(ctx context.Context, ref string) (*types.Order, error) {
	var order types.Order
	return &order, nil
}

func (os *NoopOrder) GetByOrderID(ctx context.Context, orderID string, version *uint64) (*types.Order, error) {
	var order types.Order
	return &order, nil
}

func (os *NoopOrder) GetMarketDepth(ctx context.Context, market string, limit uint64) (*types.MarketDepth, error) {

	depth, ok := os.depth[market]
	if !ok || depth == nil {
		// When a market is new with no orders there will not be any market depth/order book
		// so we do not need to try and calculate the depth cumulative volumes etc
		return &types.MarketDepth{
			MarketID: market,
			Buy:      []*types.PriceLevel{},
			Sell:     []*types.PriceLevel{},
		}, nil
	}

	// load from store
	buy := depth.BuySide(limit)
	sell := depth.SellSide(limit)

	buyPtr := make([]*types.PriceLevel, 0, len(buy))
	sellPtr := make([]*types.PriceLevel, 0, len(sell))

	ctx, cancel := context.WithTimeout(ctx, os.Config.Timeout.Duration)
	defer cancel()
	deadline, _ := ctx.Deadline()
	// 2 routines, each can push one error on here, so buffer to avoid deadlock
	errCh := make(chan error, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// recalculate accumulated volume, concurrently rather than sequentially
	// make the most of the time we have
	// --- buy side ---
	go func() {
		defer wg.Done()
		var cumulativeVolume uint64
		for i, b := range buy {
			select {
			case <-ctx.Done():
				if deadline.Before(time.Now()) {
					errCh <- ErrTimeoutReached
				}
				return
			default:
				// keep running total
				cumulativeVolume += b.Volume
				buy[i].CumulativeVolume = cumulativeVolume
				buyPtr = append(buyPtr, &buy[i].PriceLevel)
			}
		}
	}()
	// --- sell side ---
	go func() {
		defer wg.Done()
		var cumulativeVolume uint64
		for i, s := range sell {
			select {
			case <-ctx.Done():
				if deadline.Before(time.Now()) {
					errCh <- ErrTimeoutReached
				}
				return
			default:
				// keep running total
				cumulativeVolume += s.Volume
				sell[i].CumulativeVolume = cumulativeVolume
				sellPtr = append(sellPtr, &sell[i].PriceLevel)
			}
		}
	}()
	wg.Wait()
	close(errCh)
	// the second error is the same, they're both ctx.Err()
	for err := range errCh {
		return nil, err
	}

	// return new re-calculated market depth for each side of order book
	return &types.MarketDepth{
		MarketID: market,
		Buy:      buyPtr,
		Sell:     sellPtr,
	}, nil
}