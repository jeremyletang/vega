package steps

import (
	"context"

	"code.vegaprotocol.io/vega/execution"
	"code.vegaprotocol.io/vega/integration/helpers"
	"code.vegaprotocol.io/vega/integration/stubs"
	types "code.vegaprotocol.io/vega/proto"
	commandspb "code.vegaprotocol.io/vega/proto/commands/v1"

	"github.com/cucumber/godog/gherkin"
)

func TradersCancelTheFollowingOrders(
	broker *stubs.BrokerStub,
	exec *execution.Engine,
	errorHandler *helpers.ErrorHandler,
	orders *gherkin.DataTable,
) error {
	for _, row := range TableWrapper(*orders).Parse() {
		trader := row.MustStr("trader")
		reference := row.Str("reference")
		marketID := row.Str("market id")

		var orders []types.Order
		switch {
		case marketID != "":
			orders = broker.GetOrdersByPartyAndMarket(trader, marketID)
		default:
			o, err := broker.GetByReference(trader, reference)
			if err != nil {
				return errOrderNotFound(trader, reference, err)
			}
			orders = append(orders, o)
		}

		for _, o := range orders {
			cancel := commandspb.OrderCancellation{
				OrderId:  o.Id,
				MarketId: o.MarketId,
			}
			reference = o.Reference
			cancelOrder(exec, errorHandler, cancel, trader, reference)
		}

	}

	return nil
}

func cancelOrder(exec *execution.Engine, errHandler *helpers.ErrorHandler, cancel commandspb.OrderCancellation, party string, ref string) {
	if _, err := exec.CancelOrder(context.Background(), &cancel, party); err != nil {
		errHandler.HandleError(CancelOrderError{
			reference: ref,
			request:   cancel,
			Err:       err,
		})
	}
}
