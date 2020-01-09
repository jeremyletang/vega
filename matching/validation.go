package matching

import (
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/metrics"
	types "code.vegaprotocol.io/vega/proto"
)

const minOrderIDLen = 22
const maxOrderIDLen = 22

func (b OrderBook) validateOrder(orderMessage *types.Order) (err error) {
	timer := metrics.NewTimeCounter(b.marketID, "matching", "validateOrder")
	if orderMessage.MarketID != b.marketID {
		b.log.Error("Market ID mismatch",
			logging.String("market", orderMessage.MarketID),
			logging.String("order-book", b.marketID),
			logging.Order(*orderMessage))
		err = types.ErrInvalidMarketID
	} else if orderMessage.CreatedAt < b.latestTimestamp {
		err = types.ErrOrderOutOfSequence
	} else if orderMessage.Remaining > 0 && orderMessage.Remaining != orderMessage.Size {
		err = types.ErrInvalidRemainingSize
	} else if orderMessage.TimeInForce == types.Order_GTT && orderMessage.ExpiresAt == 0 {
		// if order is GTT, validate timestamp and convert to block number
		err = types.ErrInvalidExpirationDatetime
	}
	timer.EngineTimeCounterAdd()
	return
}

func validateOrderID(orderID string) error {
	idLen := len(orderID)
	if idLen < minOrderIDLen && idLen > maxOrderIDLen {
		return types.ErrInvalidOrderID
	}
	return nil
}
