package steps

import (
	"fmt"

	types "code.vegaprotocol.io/protos/vega"
	"github.com/cucumber/godog"
)

func TheLiquidityProviderFeeSharesForTheMarketShouldBe(engine Execution, marketID string, table *godog.Table) error {
	marketData, err := engine.GetMarketData(marketID)
	if err != nil {
		return errMarketDataNotFound(marketID, err)
	}

	for _, row := range parseLiquidityFeeSharesTable(table) {
		expected := types.LiquidityProviderFeeShare{
			Party:                 row.MustStr("party"),
			EquityLikeShare:       row.MustStr("equity like share"),
			AverageEntryValuation: row.MustStr("average entry valuation"),
		}

		var found bool
		var got []types.LiquidityProviderFeeShare
		for _, v := range marketData.LiquidityProviderFeeShare {
			got = append(got, *v)
			if v.Party == expected.Party &&
				v.EquityLikeShare == expected.EquityLikeShare &&
				v.AverageEntryValuation == expected.AverageEntryValuation {
				found = true
			}
		}

		if !found {
			return errMissingLPFeeShare(marketID, expected, got)
		}
	}

	return nil
}

func errMissingLPFeeShare(market string, expected types.LiquidityProviderFeeShare, got []types.LiquidityProviderFeeShare) error {
	return fmt.Errorf("missing fee share for market %s got %#v, want %#v", market, expected, got)
}

func parseLiquidityFeeSharesTable(table *godog.Table) []RowWrapper {
	return StrictParseTable(table, []string{
		"party",
		"equity like share",
		"average entry valuation",
	}, []string{})
}
