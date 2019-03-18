package storage

import (
	"code.vegaprotocol.io/vega/internal/logging"
	types "code.vegaprotocol.io/vega/proto"

	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

// Store provides the data storage contract for markets.
//go:generate go run github.com/golang/mock/mockgen -destination newmocks/market_store_mock.go -package newmocks code.vegaprotocol.io/vega/internal/storage MarketStore
type MarketStore interface {
	//Subscribe(markets chan<- []types.Market) uint64
	//Unsubscribe(id uint64) error

	// Post adds a market to the store, this adds
	// to queue the operation to be committed later.
	Post(party *types.Market) error

	// Commit typically saves any operations that are queued to underlying storage,
	// if supported by underlying storage implementation.
	Commit() error

	// Close can be called to clean up and close any storage
	// connections held by the underlying storage mechanism.
	Close() error

	// GetByID searches for the given market by id in the underlying store.
	GetByID(name string) (*types.Market, error)

	// GetAll returns all markets in the underlying store.
	GetAll() ([]*types.Market, error)
}

// memMarketStore is used for memory/RAM based markets storage.
type badgerMarketStore struct {
	*Config
	badger *badgerStore
}

// NewMarketStore returns a concrete implementation of MarketStore.
func NewMarketStore(c *Config) (MarketStore, error) {
	err := InitStoreDirectory(c.MarketStoreDirPath)
	if err != nil {
		return nil, errors.Wrap(err, "error on init badger database for candles storage")
	}
	db, err := badger.Open(customBadgerOptions(c.MarketStoreDirPath, c.GetLogger()))
	if err != nil {
		return nil, errors.Wrap(err, "error opening badger database for candles storage")
	}
	bs := badgerStore{db: db}
	return &badgerMarketStore{
		Config: c,
		badger: &bs,
	}, nil
}

// Post saves a given market to the mem-store.
func (ms *badgerMarketStore) Post(market *types.Market) error {
	buf, err := proto.Marshal(market)
	if err != nil {
		ms.log.Error("unable to marshal market",
			logging.Error(err),
			logging.String("market-id", market.Id),
		)
		return err
	}
	marketKey := ms.badger.marketKey(market.Id)
	err = ms.badger.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(marketKey, buf)
		if err != nil {
			ms.log.Error("unable to save market in badger",
				logging.Error(err),
				logging.String("market-id", market.Id),
			)
			return err
		}
		return nil
	})

	return err
}

// GetByID searches for the given market by id in the mem-store.
func (ms *badgerMarketStore) GetByID(id string) (*types.Market, error) {
	market := types.Market{}
	var buf []byte
	marketKey := ms.badger.marketKey(id)
	err := ms.badger.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(marketKey)
		if err != nil {
			return err
		}
		// fine to use value copy here, only one ID to get
		buf, err = item.ValueCopy(nil)
		return err
	})

	if err != nil {
		ms.log.Error("unable to get market from badger store",
			logging.Error(err),
			logging.String("market-id", id),
		)
		return nil, err
	}

	err = proto.Unmarshal(buf, &market)
	if err != nil {
		ms.log.Error("unable to unmarshal market from badger store",
			logging.Error(err),
			logging.String("market-id", id),
		)
		return nil, err
	}
	return &market, nil
}

// GetAll returns all markets in the mem-store.
func (ms *badgerMarketStore) GetAll() ([]*types.Market, error) {
	out := []*types.Market{}
	return out, nil
}

// Commit typically saves any operations that are queued to underlying storage,
// if supported by underlying storage implementation.
func (ms *badgerMarketStore) Commit() error {
	// Not required with a mem-store implementation.
	return nil
}

// Close can be called to clean up and close any storage
// connections held by the underlying storage mechanism.
func (ms *badgerMarketStore) Close() error {
	return ms.badger.db.Close()
}
