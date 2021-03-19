package steps

import (
	"context"
	"time"

	"github.com/cucumber/godog/gherkin"
	uuid "github.com/satori/go.uuid"

	"code.vegaprotocol.io/vega/execution"
	types "code.vegaprotocol.io/vega/proto"
)

func TradersPlaceOrders(
	exec *execution.Engine,
	table *gherkin.DataTable,
) error {
	for _, row := range TableWrapper(*table).Parse() {
		trader := row.Str("trader")
		marketID := row.Str("market id")
		side := row.Side("side")
		volume := row.U64("volume")
		price := row.U64("price")
		oty := row.OrderType("type")
		tif := row.TIF("tif")
		reference := row.Str("reference")

		var expiresAt int64
		if oty != types.Order_TYPE_MARKET {
			expiresAt = time.Now().Add(24 * time.Hour).UnixNano()
		}

		order := types.Order{
			Status:      types.Order_STATUS_ACTIVE,
			Id:          uuid.NewV4().String(),
			MarketId:    marketID,
			PartyId:     trader,
			Side:        side,
			Price:       price,
			Size:        volume,
			Remaining:   volume,
			ExpiresAt:   expiresAt,
			Type:        oty,
			TimeInForce: tif,
			CreatedAt:   time.Now().UnixNano(),
			Reference:   reference,
		}
		_, err := exec.SubmitOrder(context.Background(), &order)
		if err != nil {
			errCh <- err
		}
		//if err != nil {
		//	return errUnableToPlaceOrder(trader, reference, err)
		//}
		//
		//resultingTrades := row.U64("resulting trades")
		//if len(result.Trades) != int(resultingTrades) {
		//	return errWrongNumberOfTrades(resultingTrades, result)
		//}
	}
	return nil
}
