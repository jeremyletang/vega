// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.VEGA file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package supplied

import (
	"context"
	"errors"

	"code.vegaprotocol.io/vega/core/risk"
	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/core/types/statevar"
	"code.vegaprotocol.io/vega/libs/num"
	"code.vegaprotocol.io/vega/logging"
)

// ErrNoValidOrders informs that there weren't any valid orders to cover the liquidity obligation with.
// This could happen when for a given side (buy or sell) limit orders don't supply enough liquidity and there aren't any
// valid pegged orders (all the prives are invalid) to cover it with.
var (
	ErrNoValidOrders = errors.New("no valid orders to cover the liquidity obligation with")
)

// LiquidityOrder contains information required to compute volume required to fullfil liquidity obligation per set of liquidity provision orders for one side of the order book.
type LiquidityOrder struct {
	OrderID string

	Price      *num.Uint
	Proportion uint64
	Peg        *types.PeggedOrder

	LiquidityImpliedVolume uint64
}

// RiskModel allows calculation of min/max price range and a probability of trading.
//
//go:generate go run github.com/golang/mock/mockgen -destination mocks/risk_model_mock.go -package mocks code.vegaprotocol.io/vega/core/liquidity/supplied RiskModel
type RiskModel interface {
	ProbabilityOfTrading(currentPrice, orderPrice, minPrice, maxPrice num.Decimal, yearFraction num.Decimal, isBid, applyMinMax bool) num.Decimal
	GetProjectionHorizon() num.Decimal
}

// PriceMonitor provides the range of valid prices, that is prices that wouldn't trade the current trading mode
//
//go:generate go run github.com/golang/mock/mockgen -destination mocks/price_monitor_mock.go -package mocks code.vegaprotocol.io/vega/core/liquidity/supplied PriceMonitor
type PriceMonitor interface {
	GetValidPriceRange() (num.WrappedDecimal, num.WrappedDecimal)
}

type StateVarEngine interface {
	RegisterStateVariable(asset, market, name string, converter statevar.Converter, startCalculation func(string, statevar.FinaliseCalculation), trigger []statevar.EventType, result func(context.Context, statevar.StateVariableResult) error) error
}

// Engine provides functionality related to supplied liquidity.
type Engine struct {
	rm                             RiskModel
	pm                             PriceMonitor
	marketID                       string
	horizon                        num.Decimal // projection horizon used in probability calculations
	probabilityOfTradingTauScaling num.Decimal
	minProbabilityOfTrading        num.Decimal
	pot                            *probabilityOfTrading
	potInitialised                 bool

	getBestStaticPrices func() (num.Decimal, num.Decimal, error)
	log                 *logging.Logger
	positionFactor      num.Decimal
}

// NewEngine returns a reference to a new supplied liquidity calculation engine.
func NewEngine(riskModel RiskModel, priceMonitor PriceMonitor, asset, marketID string, stateVarEngine StateVarEngine, log *logging.Logger, positionFactor num.Decimal) *Engine {
	e := &Engine{
		rm:                             riskModel,
		pm:                             priceMonitor,
		marketID:                       marketID,
		horizon:                        riskModel.GetProjectionHorizon(),
		probabilityOfTradingTauScaling: num.DecimalFromInt64(1), // this is the same as the default in the netparams
		minProbabilityOfTrading:        defaultMinimumProbabilityOfTrading,
		pot:                            &probabilityOfTrading{},
		potInitialised:                 false,
		log:                            log,
		positionFactor:                 positionFactor,
	}

	stateVarEngine.RegisterStateVariable(asset, marketID, "probability_of_trading", probabilityOfTradingConverter{}, e.startCalcProbOfTrading, []statevar.EventType{statevar.EventTypeTimeTrigger, statevar.EventTypeAuctionEnded, statevar.EventTypeOpeningAuctionFirstUncrossingPrice}, e.updateProbabilities)
	return e
}

func (e *Engine) UpdateMarketConfig(riskModel risk.Model, monitor PriceMonitor) {
	e.rm = riskModel
	e.pm = monitor
	e.horizon = riskModel.GetProjectionHorizon()
	e.potInitialised = false
}

func (e *Engine) SetGetStaticPricesFunc(f func() (num.Decimal, num.Decimal, error)) {
	e.getBestStaticPrices = f
}

func (e *Engine) OnMinProbabilityOfTradingLPOrdersUpdate(v num.Decimal) {
	e.minProbabilityOfTrading = v
}

func (e *Engine) OnProbabilityOfTradingTauScalingUpdate(v num.Decimal) {
	e.probabilityOfTradingTauScaling = v
}

// CalculateSuppliedLiquidity returns the current supplied liquidity per specified current mark price and order set.
func (e *Engine) CalculateSuppliedLiquidity(
	bestBidPrice, bestAskPrice num.Decimal,
	orders []*types.Order,
) *num.Uint {
	bLiq, sLiq := e.calculateBuySellLiquidityWithMinMax(bestBidPrice, bestAskPrice, orders)

	return num.Min(bLiq, sLiq)
}

// CalculateLiquidityImpliedVolumes updates the LiquidityImpliedSize fields in LiquidityOrderReference so that the liquidity commitment is met.
// Current market price, liquidity obligation, and orders must be specified.
// Note that due to integer order size the actual liquidity provided will be more than or equal to the commitment amount.
func (e *Engine) CalculateLiquidityImpliedVolumes(
	bestBidPrice, bestAskPrice num.Decimal,
	liquidityObligation *num.Uint,
	orders []*types.Order,
	buyShapes, sellShapes []*LiquidityOrder,
) error {
	buySupplied, sellSupplied := e.calculateBuySellLiquidityWithMinMax(bestBidPrice, bestAskPrice, orders)

	buyRemaining := liquidityObligation.Clone()
	buyRemaining.Sub(buyRemaining, buySupplied)
	if err := e.updateSizes(buyRemaining, bestBidPrice, bestAskPrice, buyShapes, true); err != nil {
		return err
	}

	sellRemaining := liquidityObligation.Clone()
	sellRemaining.Sub(sellRemaining, sellSupplied)
	if err := e.updateSizes(sellRemaining, bestBidPrice, bestAskPrice, sellShapes, false); err != nil {
		return err
	}

	return nil
}

// calculateBuySellLiquidityWithMinMax returns the current supplied liquidity per market specified in the constructor.
func (e *Engine) calculateBuySellLiquidityWithMinMax(bestBidPrice, bestAskPrice num.Decimal, orders []*types.Order) (*num.Uint, *num.Uint) {
	bLiq := num.DecimalZero()
	sLiq := num.DecimalZero()
	min, max := e.pm.GetValidPriceRange()
	for _, o := range orders {
		if o.Side == types.SideBuy {
			// float64(o.Price.Uint64()) * float64(o.Remaining) * prob
			prob := getProbabilityOfTrading(bestBidPrice, bestAskPrice, min.Original(), max.Original(), e.pot, o.Price.ToDecimal(), true, e.minProbabilityOfTrading)
			if e.log.GetLevel() <= logging.DebugLevel {
				e.log.Debug("probability of trading", logging.Decimal("order-price", o.Price.ToDecimal()), logging.Decimal("prob", prob))
			}
			d := prob.Mul(num.DecimalFromUint(num.NewUint(o.Remaining)))
			d = d.Mul(num.DecimalFromUint(o.Price))
			bLiq = bLiq.Add(d)
		}
		if o.Side == types.SideSell {
			// float64(o.Price.Uint64()) * float64(o.Remaining) * prob
			prob := getProbabilityOfTrading(bestBidPrice, bestAskPrice, min.Original(), max.Original(), e.pot, o.Price.ToDecimal(), false, e.minProbabilityOfTrading)
			if e.log.GetLevel() <= logging.DebugLevel {
				e.log.Debug("probability of trading", logging.Decimal("order-price", o.Price.ToDecimal()), logging.Decimal("prob", prob))
			}
			d := prob.Mul(num.DecimalFromUint(num.NewUint(o.Remaining)))
			d = d.Mul(num.DecimalFromUint(o.Price))
			sLiq = sLiq.Add(d)
		}
	}

	// descale provided liquidity by 10^pdp
	bl, _ := num.UintFromDecimal(bLiq.Div(e.positionFactor))
	sl, _ := num.UintFromDecimal(sLiq.Div(e.positionFactor))
	return bl, sl
}

func (e *Engine) updateSizes(liquidityObligation *num.Uint, bestBidPrice, bestAskPrice num.Decimal, orders []*LiquidityOrder, isBid bool) error {
	if liquidityObligation.IsZero() || liquidityObligation.IsNegative() {
		setSizesTo0(orders)
		return nil
	}
	min, max := e.pm.GetValidPriceRange()
	sum := num.DecimalZero()
	probs := make([]num.Decimal, 0, len(orders))
	validatedProportions := make([]num.Decimal, 0, len(orders))

	for _, o := range orders {
		proportion := num.DecimalFromUint(num.NewUint(o.Proportion))

		prob := getProbabilityOfTrading(bestBidPrice, bestAskPrice, min.Original(), max.Original(), e.pot, o.Price.ToDecimal(), isBid, e.minProbabilityOfTrading)

		if e.log.GetLevel() <= logging.DebugLevel {
			e.log.Debug("$probability of trading$",
				logging.String("market-id", e.marketID),
				logging.String("best-bid", bestBidPrice.String()),
				logging.String("best-ask", bestAskPrice.String()),
				logging.String("min", min.Original().String()),
				logging.String("max", max.Original().String()),
				logging.String("order-price", o.Price.String()),
				logging.Bool("is-bid", isBid),
				logging.String("probability", prob.String()))
		}
		if prob.IsZero() || prob.IsNegative() {
			proportion = num.DecimalZero()
		}

		sum = sum.Add(proportion)
		validatedProportions = append(validatedProportions, proportion)
		probs = append(probs, prob)
	}
	if sum.IsZero() {
		return ErrNoValidOrders
	}

	for i, o := range orders {
		scaling := num.DecimalZero()
		if prob := probs[i]; !prob.IsZero() {
			fraction := validatedProportions[i].Div(sum)
			scaling = fraction.Div(prob)
		}
		// uint64(math.Ceil(liquidityObligation * scaling / float64(o.Price.Uint64())))
		d := num.DecimalFromUint(liquidityObligation)
		d = d.Mul(scaling)
		// scale the volume by 10^pdp BEFORE dividing by price for better precision.
		liv, _ := num.UintFromDecimal(d.Mul(e.positionFactor).Div(num.DecimalFromUint(o.Price)).Ceil())
		o.LiquidityImpliedVolume = liv.Uint64()
	}
	return nil
}

func setSizesTo0(orders []*LiquidityOrder) {
	for _, o := range orders {
		o.LiquidityImpliedVolume = 0
	}
}
