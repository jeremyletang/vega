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

package governance

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"code.vegaprotocol.io/vega/core/netparams"
	"code.vegaprotocol.io/vega/core/oracles"
	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/libs/num"
	proto "code.vegaprotocol.io/vega/protos/vega"
	oraclespb "code.vegaprotocol.io/vega/protos/vega/oracles/v1"
)

var (
	// ErrMissingProduct is returned if selected product is nil.
	ErrMissingProduct = errors.New("missing product")
	// ErrUnsupportedProduct is returned if selected product is not supported.
	ErrUnsupportedProduct = errors.New("product type is not supported")
	// ErrUnsupportedRiskParameters is returned if risk parameters supplied via governance are not yet supported.
	ErrUnsupportedRiskParameters = errors.New("risk model parameters are not supported")
	// ErrMissingRiskParameters ...
	ErrMissingRiskParameters = errors.New("missing risk parameters")
	// ErrMissingOracleSpecBinding is returned when the oracle spec binding is absent.
	ErrMissingOracleSpecBinding = errors.New("missing oracle spec binding")
	// ErrMissingOracleSpecForSettlementPrice is returned when the oracle spec for settlement price is absent.
	ErrMissingOracleSpecForSettlementPrice = errors.New("missing oracle spec for settlement price")
	// ErrMissingOracleSpecForTradingTermination is returned when the oracle spec for trading termination is absent.
	ErrMissingOracleSpecForTradingTermination = errors.New("missing oracle spec for trading termination")
	// ErrOracleSpecTerminationTimeBeforeEnactment is returned when termination time is before enactment
	// for time triggered termination condition.
	ErrOracleSpecTerminationTimeBeforeEnactment = errors.New("oracle spec termination time before enactment")
	// ErrMissingFutureProduct is returned when future product is absent from the instrument.
	ErrMissingFutureProduct = errors.New("missing future product")
	// ErrInvalidRiskParameter ...
	ErrInvalidRiskParameter = errors.New("invalid risk parameter")
)

func assignProduct(
	source *types.InstrumentConfiguration,
	target *types.Instrument,
) (proto.ProposalError, error) {
	switch product := source.Product.(type) {
	case *types.InstrumentConfigurationFuture:
		if product.Future == nil {
			return types.ProposalErrorInvalidFutureProduct, ErrMissingFutureProduct
		}
		if product.Future.OracleSpecForSettlementPrice == nil {
			return types.ProposalErrorInvalidFutureProduct, ErrMissingOracleSpecForSettlementPrice
		}
		if product.Future.OracleSpecForTradingTermination == nil {
			return types.ProposalErrorInvalidFutureProduct, ErrMissingOracleSpecForTradingTermination
		}
		if product.Future.OracleSpecBinding == nil {
			return types.ProposalErrorInvalidFutureProduct, ErrMissingOracleSpecBinding
		}

		target.Product = &types.InstrumentFuture{
			Future: &types.Future{
				SettlementAsset:                 product.Future.SettlementAsset,
				QuoteName:                       product.Future.QuoteName,
				OracleSpecForSettlementPrice:    product.Future.OracleSpecForSettlementPrice.ToOracleSpec(),
				OracleSpecForTradingTermination: product.Future.OracleSpecForTradingTermination.ToOracleSpec(),
				SettlementDataDecimals:          product.Future.SettlementDataDecimalPlaces,
				OracleSpecBinding:               product.Future.OracleSpecBinding,
			},
		}
	default:
		return types.ProposalErrorUnsupportedProduct, ErrUnsupportedProduct
	}
	return types.ProposalErrorUnspecified, nil
}

func createInstrument(
	input *types.InstrumentConfiguration,
	tags []string,
) (*types.Instrument, types.ProposalError, error) {
	result := &types.Instrument{
		Name: input.Name,
		Code: input.Code,
		Metadata: &types.InstrumentMetadata{
			Tags: tags,
		},
	}

	if perr, err := assignProduct(input, result); err != nil {
		return nil, perr, err
	}
	return result, types.ProposalErrorUnspecified, nil
}

func assignRiskModel(definition *types.NewMarketConfiguration, target *types.TradableInstrument) error {
	switch parameters := definition.RiskParameters.(type) {
	case *types.NewMarketConfigurationSimple:
		target.RiskModel = &types.TradableInstrumentSimpleRiskModel{
			SimpleRiskModel: &types.SimpleRiskModel{
				Params: parameters.Simple,
			},
		}
	case *types.NewMarketConfigurationLogNormal:
		target.RiskModel = &types.TradableInstrumentLogNormalRiskModel{
			LogNormalRiskModel: parameters.LogNormal,
		}
	default:
		return ErrUnsupportedRiskParameters
	}
	return nil
}

func buildMarketFromProposal(
	marketID string,
	definition *types.NewMarket,
	netp NetParams,
	openingAuctionDuration time.Duration,
) (*types.Market, types.ProposalError, error) {
	instrument, perr, err := createInstrument(definition.Changes.Instrument, definition.Changes.Metadata)
	if err != nil {
		return nil, perr, err
	}

	// get factors for the market
	makerFee, _ := netp.Get(netparams.MarketFeeFactorsMakerFee)
	infraFee, _ := netp.Get(netparams.MarketFeeFactorsInfrastructureFee)
	// get the margin scaling factors
	scalingFactors := proto.ScalingFactors{}
	_ = netp.GetJSONStruct(netparams.MarketMarginScalingFactors, &scalingFactors)
	// get price monitoring parameters
	if definition.Changes.PriceMonitoringParameters == nil {
		pmParams := &proto.PriceMonitoringParameters{}
		_ = netp.GetJSONStruct(netparams.MarketPriceMonitoringDefaultParameters, pmParams)
		definition.Changes.PriceMonitoringParameters = types.PriceMonitoringParametersFromProto(pmParams)
	}

	if definition.Changes.LiquidityMonitoringParameters == nil ||
		definition.Changes.LiquidityMonitoringParameters.TargetStakeParameters == nil {
		// get target stake parameters
		tsTimeWindow, _ := netp.GetDuration(netparams.MarketTargetStakeTimeWindow)
		tsScalingFactor, _ := netp.GetDecimal(netparams.MarketTargetStakeScalingFactor)
		// get triggering ratio
		triggeringRatio, _ := netp.GetDecimal(netparams.MarketLiquidityTargetStakeTriggeringRatio)

		params := &types.TargetStakeParameters{
			TimeWindow:    int64(tsTimeWindow.Seconds()),
			ScalingFactor: tsScalingFactor,
		}

		if definition.Changes.LiquidityMonitoringParameters == nil {
			definition.Changes.LiquidityMonitoringParameters = &types.LiquidityMonitoringParameters{
				TargetStakeParameters: params,
				TriggeringRatio:       triggeringRatio,
			}
		} else {
			definition.Changes.LiquidityMonitoringParameters.TargetStakeParameters = params
		}
	}

	makerFeeDec, _ := num.DecimalFromString(makerFee)
	infraFeeDec, _ := num.DecimalFromString(infraFee)
	market := &types.Market{
		ID:                    marketID,
		DecimalPlaces:         definition.Changes.DecimalPlaces,
		PositionDecimalPlaces: definition.Changes.PositionDecimalPlaces,
		Fees: &types.Fees{
			Factors: &types.FeeFactors{
				MakerFee:          makerFeeDec,
				InfrastructureFee: infraFeeDec,
			},
		},
		OpeningAuction: &types.AuctionDuration{
			Duration: int64(openingAuctionDuration.Seconds()),
		},
		TradableInstrument: &types.TradableInstrument{
			Instrument: instrument,
			MarginCalculator: &types.MarginCalculator{
				ScalingFactors: types.ScalingFactorsFromProto(&scalingFactors),
			},
		},
		PriceMonitoringSettings: &types.PriceMonitoringSettings{
			Parameters: definition.Changes.PriceMonitoringParameters,
		},
		LiquidityMonitoringParameters: definition.Changes.LiquidityMonitoringParameters,
	}
	if err := assignRiskModel(definition.Changes, market.TradableInstrument); err != nil {
		return nil, types.ProposalErrorUnspecified, err
	}
	return market, types.ProposalErrorUnspecified, nil
}

func validateAsset(assetID string, decimals uint64, assets Assets, deepCheck bool) (types.ProposalError, error) {
	if len(assetID) <= 0 {
		return types.ProposalErrorInvalidAsset, errors.New("missing asset ID")
	}

	if !deepCheck {
		return types.ProposalErrorUnspecified, nil
	}

	asset, err := assets.Get(assetID)
	if err != nil {
		return types.ProposalErrorInvalidAsset, err
	}
	if !assets.IsEnabled(assetID) {
		return types.ProposalErrorInvalidAsset,
			fmt.Errorf("assets is not enabled %v", assetID)
	}
	// decimal places asset less than market -> invalid.
	// @TODO add a specific error for this validation?
	if asset.DecimalPlaces() < decimals {
		return types.ProposalErrorTooManyMarketDecimalPlaces, errors.New("market cannot have more decimal places than assets")
	}

	return types.ProposalErrorUnspecified, nil
}

func validateFuture(future *types.FutureProduct, decimals uint64, assets Assets, et *enactmentTime, deepCheck bool) (types.ProposalError, error) {
	if future.OracleSpecForSettlementPrice == nil {
		return types.ProposalErrorInvalidFutureProduct, ErrMissingOracleSpecForSettlementPrice
	}

	if future.OracleSpecForTradingTermination == nil {
		return types.ProposalErrorInvalidFutureProduct, ErrMissingOracleSpecForTradingTermination
	}

	if !et.shouldNotVerify {
		for i, f := range future.OracleSpecForTradingTermination.ToOracleSpec().Filters {
			if f.Key.Type == oraclespb.PropertyKey_TYPE_TIMESTAMP {
				for j, cond := range f.Conditions {
					v, err := strconv.ParseInt(cond.Value, 10, 64)
					if err != nil {
						return types.ProposalErrorInvalidFutureProduct, err
					}

					future.OracleSpecForTradingTermination.Filters[i].Conditions[j].Value = strconv.FormatInt(v, 10)
					if v <= et.current {
						return types.ProposalErrorInvalidFutureProduct, ErrOracleSpecTerminationTimeBeforeEnactment
					}
				}
			}
		}
	}

	if future.OracleSpecBinding == nil {
		return types.ProposalErrorInvalidFutureProduct, ErrMissingOracleSpecBinding
	}

	// ensure the oracle spec for settlement price can be constructed
	ospec, err := oracles.NewOracleSpec(*future.OracleSpecForSettlementPrice.ToOracleSpec())
	if err != nil {
		return types.ProposalErrorInvalidFutureProduct, err
	}
	if err := ospec.EnsureBoundableProperty(future.OracleSpecBinding.SettlementPriceProperty, oraclespb.PropertyKey_TYPE_INTEGER); err != nil {
		return types.ProposalErrorInvalidFutureProduct, fmt.Errorf("invalid oracle spec binding for settlement price: %w", err)
	}

	ospec, err = oracles.NewOracleSpec(*future.OracleSpecForTradingTermination.ToOracleSpec())
	if err != nil {
		return types.ProposalErrorInvalidFutureProduct, err
	}

	switch future.OracleSpecBinding.TradingTerminationProperty {
	case oracles.BuiltinOracleTimestamp:
		if err := ospec.EnsureBoundableProperty(future.OracleSpecBinding.TradingTerminationProperty, oraclespb.PropertyKey_TYPE_TIMESTAMP); err != nil {
			return types.ProposalErrorInvalidFutureProduct, fmt.Errorf("invalid oracle spec binding for trading termination: %w", err)
		}
	default:
		if err := ospec.EnsureBoundableProperty(future.OracleSpecBinding.TradingTerminationProperty, oraclespb.PropertyKey_TYPE_BOOLEAN); err != nil {
			return types.ProposalErrorInvalidFutureProduct, fmt.Errorf("invalid oracle spec binding for trading termination: %w", err)
		}
	}

	return validateAsset(future.SettlementAsset, decimals, assets, deepCheck)
}

func validateNewInstrument(instrument *types.InstrumentConfiguration, decimals uint64, assets Assets, et *enactmentTime, deepCheck bool) (types.ProposalError, error) {
	switch product := instrument.Product.(type) {
	case nil:
		return types.ProposalErrorNoProduct, ErrMissingProduct
	case *types.InstrumentConfigurationFuture:
		return validateFuture(product.Future, decimals, assets, et, deepCheck)
	default:
		return types.ProposalErrorUnsupportedProduct, ErrUnsupportedProduct
	}
}

func validateRiskParameters(rp interface{}) (types.ProposalError, error) {
	switch r := rp.(type) {
	case *types.NewMarketConfigurationSimple:
		return types.ProposalErrorUnspecified, nil
	case *types.UpdateMarketConfigurationSimple:
		return types.ProposalErrorUnspecified, nil
	case *types.NewMarketConfigurationLogNormal:
		if r.LogNormal.Params == nil {
			return types.ProposalErrorInvalidRiskParameter, ErrInvalidRiskParameter
		}
		return types.ProposalErrorUnspecified, nil
	case *types.UpdateMarketConfigurationLogNormal:
		if r.LogNormal.Params == nil {
			return types.ProposalErrorInvalidRiskParameter, ErrInvalidRiskParameter
		}
		return types.ProposalErrorUnspecified, nil
	case nil:
		return types.ProposalErrorNoRiskParameters, ErrMissingRiskParameters
	default:
		return types.ProposalErrorUnknownRiskParameterType, ErrUnsupportedRiskParameters
	}
}

func validateAuctionDuration(proposedDuration time.Duration, netp NetParams) (types.ProposalError, error) {
	minAuctionDuration, _ := netp.GetDuration(netparams.MarketAuctionMinimumDuration)
	if proposedDuration != 0 && proposedDuration < minAuctionDuration {
		// Auction duration is too small
		return types.ProposalErrorOpeningAuctionDurationTooSmall,
			fmt.Errorf("proposal opening auction duration is too short, expected > %v, got %v", minAuctionDuration, proposedDuration)
	}
	maxAuctionDuration, _ := netp.GetDuration(netparams.MarketAuctionMaximumDuration)
	if proposedDuration > maxAuctionDuration {
		// Auction duration is too large
		return types.ProposalErrorOpeningAuctionDurationTooLarge,
			fmt.Errorf("proposal opening auction duration is too long, expected < %v, got %v", maxAuctionDuration, proposedDuration)
	}
	return types.ProposalErrorUnspecified, nil
}

// ValidateNewMarket checks new market proposal terms.
func validateNewMarketChange(
	terms *types.NewMarket,
	assets Assets,
	deepCheck bool,
	netp NetParams,
	openingAuctionDuration time.Duration,
	etu *enactmentTime,
) (types.ProposalError, error) {
	if perr, err := validateNewInstrument(terms.Changes.Instrument, terms.Changes.DecimalPlaces, assets, etu, deepCheck); err != nil {
		return perr, err
	}
	if perr, err := validateRiskParameters(terms.Changes.RiskParameters); err != nil {
		return perr, err
	}
	if perr, err := validateAuctionDuration(openingAuctionDuration, netp); err != nil {
		return perr, err
	}
	if terms.Changes.PriceMonitoringParameters != nil && len(terms.Changes.PriceMonitoringParameters.Triggers) > 5 {
		return types.ProposalErrorTooManyPriceMonitoringTriggers,
			fmt.Errorf("%v price monitoring triggers set, maximum allowed is 5", len(terms.Changes.PriceMonitoringParameters.Triggers) > 5)
	}

	return types.ProposalErrorUnspecified, nil
}

// validateUpdateMarketChange checks market update proposal terms.
func validateUpdateMarketChange(terms *types.UpdateMarket, etu *enactmentTime) (types.ProposalError, error) {
	if perr, err := validateUpdateInstrument(terms.Changes.Instrument, etu); err != nil {
		return perr, err
	}
	if perr, err := validateRiskParameters(terms.Changes.RiskParameters); err != nil {
		return perr, err
	}

	return types.ProposalErrorUnspecified, nil
}

func validateUpdateInstrument(instrument *types.UpdateInstrumentConfiguration, et *enactmentTime) (types.ProposalError, error) {
	switch product := instrument.Product.(type) {
	case nil:
		return types.ProposalErrorNoProduct, ErrMissingProduct
	case *types.UpdateInstrumentConfigurationFuture:
		return validateUpdateFuture(product.Future, et)
	default:
		return types.ProposalErrorUnsupportedProduct, ErrUnsupportedProduct
	}
}

func validateUpdateFuture(future *types.UpdateFutureProduct, et *enactmentTime) (types.ProposalError, error) {
	if future.OracleSpecForSettlementPrice == nil {
		return types.ProposalErrorInvalidFutureProduct, ErrMissingOracleSpecForSettlementPrice
	}

	if future.OracleSpecForTradingTermination == nil {
		return types.ProposalErrorInvalidFutureProduct, ErrMissingOracleSpecForTradingTermination
	}

	if !et.shouldNotVerify {
		for i, f := range future.OracleSpecForTradingTermination.ToOracleSpec().Filters {
			if f.Key.Type == oraclespb.PropertyKey_TYPE_TIMESTAMP {
				for j, cond := range f.Conditions {
					v, err := strconv.ParseInt(cond.Value, 10, 64)
					if err != nil {
						return types.ProposalErrorInvalidFutureProduct, err
					}

					future.OracleSpecForTradingTermination.Filters[i].Conditions[j].Value = strconv.FormatInt(v, 10)
					if v <= et.current {
						return types.ProposalErrorInvalidFutureProduct, ErrOracleSpecTerminationTimeBeforeEnactment
					}
				}
			}
		}
	}

	if future.OracleSpecBinding == nil {
		return types.ProposalErrorInvalidFutureProduct, ErrMissingOracleSpecBinding
	}

	// ensure the oracle spec for settlement price can be constructed
	ospec, err := oracles.NewOracleSpec(*future.OracleSpecForSettlementPrice.ToOracleSpec())
	if err != nil {
		return types.ProposalErrorInvalidFutureProduct, err
	}
	if err := ospec.EnsureBoundableProperty(future.OracleSpecBinding.SettlementPriceProperty, oraclespb.PropertyKey_TYPE_INTEGER); err != nil {
		return types.ProposalErrorInvalidFutureProduct, fmt.Errorf("invalid oracle spec binding for settlement price: %w", err)
	}

	ospec, err = oracles.NewOracleSpec(*future.OracleSpecForTradingTermination.ToOracleSpec())
	if err != nil {
		return types.ProposalErrorInvalidFutureProduct, err
	}

	switch future.OracleSpecBinding.TradingTerminationProperty {
	case oracles.BuiltinOracleTimestamp:
		if err := ospec.EnsureBoundableProperty(future.OracleSpecBinding.TradingTerminationProperty, oraclespb.PropertyKey_TYPE_TIMESTAMP); err != nil {
			return types.ProposalErrorInvalidFutureProduct, fmt.Errorf("invalid oracle spec binding for trading termination: %w", err)
		}
	default:
		if err := ospec.EnsureBoundableProperty(future.OracleSpecBinding.TradingTerminationProperty, oraclespb.PropertyKey_TYPE_BOOLEAN); err != nil {
			return types.ProposalErrorInvalidFutureProduct, fmt.Errorf("invalid oracle spec binding for trading termination: %w", err)
		}
	}

	return types.ProposalErrorUnspecified, nil
}
