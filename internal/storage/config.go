package storage

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"code.vegaprotocol.io/vega/internal/config/encoding"
	cfgencoding "code.vegaprotocol.io/vega/internal/config/encoding"
	"code.vegaprotocol.io/vega/internal/logging"

	"github.com/pkg/errors"
)

const (
	AccountStoreDataPath = "accountstore"
	CandleStoreDataPath  = "candlestore"
	MarketStoreDataPath  = "marketstore"
	OrderStoreDataPath   = "orderstore"
	TradeStoreDataPath   = "tradestore"

	// namedLogger is the identifier for package and should ideally match the package name
	// this is simply emitted as a hierarchical label e.g. 'api.grpc'.
	namedLogger = "storage"

	defaultStorageAccessTimeout = 5 * time.Second
)

// Config provides package level settings, configuration and logging.
type Config struct {
	BadgerOptions BadgerOptions
	Level         encoding.LogLevel

	AccountStoreDirPath string
	OrderStoreDirPath   string
	TradeStoreDirPath   string
	CandleStoreDirPath  string
	MarketStoreDirPath  string
	//LogPartyStoreDebug    bool
	//LogOrderStoreDebug    bool
	//LogCandleStoreDebug   bool
	LogPositionStoreDebug bool
	Timeout               cfgencoding.Duration
}

// NewConfig constructs a new Config instance with default parameters.
// This constructor is used by the vega application code. Logger is a
// pointer to a logging instance and defaultStoreDirPath is the root directory
// where all storage directories are to be read from and written to.
func NewDefaultConfig(defaultStoreDirPath string) Config {
	return Config{
		BadgerOptions:         DefaultBadgerOptions(),
		Level:                 encoding.LogLevel{Level: logging.InfoLevel},
		AccountStoreDirPath:   filepath.Join(defaultStoreDirPath, AccountStoreDataPath),
		OrderStoreDirPath:     filepath.Join(defaultStoreDirPath, OrderStoreDataPath),
		TradeStoreDirPath:     filepath.Join(defaultStoreDirPath, TradeStoreDataPath),
		CandleStoreDirPath:    filepath.Join(defaultStoreDirPath, CandleStoreDataPath),
		MarketStoreDirPath:    filepath.Join(defaultStoreDirPath, MarketStoreDataPath),
		LogPositionStoreDebug: false,
		Timeout:               cfgencoding.Duration{Duration: defaultStorageAccessTimeout},
	}
}

// FlushStores will remove/clear the badger key and value files (i.e. databases)
// from disk at the locations specified by the given storage.Config. This is
// currently used within unit and integration tests to clear between runs.
func FlushStores(log *logging.Logger, c Config) {
	paths := map[string]string{
		"account": c.AccountStoreDirPath,
		"order":   c.OrderStoreDirPath,
		"trade":   c.TradeStoreDirPath,
		"candle":  c.CandleStoreDirPath,
		"market":  c.MarketStoreDirPath,
	}
	for name, path := range paths {
		if err := os.RemoveAll(path); err != nil {
			log.Error(
				fmt.Sprintf("Failed to flush the %s path", name),
				logging.String("path", path),
				logging.Error(err),
			)
		}
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err = os.MkdirAll(path, os.ModePerm); err != nil {
				log.Error(
					fmt.Sprintf("Failed to create the %s store", name),
					logging.String("path", path),
					logging.Error(err),
				)
			}
		}
	}
}

func InitStoreDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return errors.Wrap(err, fmt.Sprintf("could not create directory path for badger data store: %s", path))
		}
	}
	return nil
}

// NewTestConfig constructs a new Config instance with test parameters.
// This constructor is exclusively used in unit tests/integration tests
func NewTestConfig() (Config, error) {
	// Test configuration for badger stores
	cfg := Config{
		BadgerOptions:         DefaultBadgerOptions(),
		OrderStoreDirPath:     fmt.Sprintf("/tmp/vegatests/orderstore-%v", randSeq(5)),
		TradeStoreDirPath:     fmt.Sprintf("/tmp/vegatests/tradestore-%v", randSeq(5)),
		CandleStoreDirPath:    fmt.Sprintf("/tmp/vegatests/candlestore-%v", randSeq(5)),
		MarketStoreDirPath:    fmt.Sprintf("/tmp/vegatests/marketstore-%v", randSeq(5)),
		LogPositionStoreDebug: true,
		Timeout:               cfgencoding.Duration{Duration: defaultStorageAccessTimeout},
	}

	if err := ensureDir(cfg.CandleStoreDirPath); err != nil {
		return Config{}, err
	}
	if err := ensureDir(cfg.OrderStoreDirPath); err != nil {
		return Config{}, err
	}
	if err := ensureDir(cfg.TradeStoreDirPath); err != nil {
		return Config{}, err
	}

	if err := ensureDir(cfg.MarketStoreDirPath); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func ensureDir(path string) error {
	const (
		dirPerms = 0700
	)

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(path, dirPerms)
		}
		return err
	}
	return nil
}

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
