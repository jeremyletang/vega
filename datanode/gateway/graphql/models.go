// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"fmt"
	"io"
	"strconv"

	"code.vegaprotocol.io/vega/protos/data-node/api/v2"
	"code.vegaprotocol.io/vega/protos/vega"
)

// One of the possible asset sources
type AssetSource interface {
	IsAssetSource()
}

// union type for wrapped events in stream PROPOSAL is mapped to governance data, something to keep in mind
type Event interface {
	IsEvent()
}

type Oracle interface {
	IsOracle()
}

type Product interface {
	IsProduct()
}

type ProposalChange interface {
	IsProposalChange()
}

type RiskModel interface {
	IsRiskModel()
}

type TransferKind interface {
	IsTransferKind()
}

// One of the possible asset sources for update assets proposals
type UpdateAssetSource interface {
	IsUpdateAssetSource()
}

type UpdateMarketRiskParameters interface {
	IsUpdateMarketRiskParameters()
}

type WithdrawalDetails interface {
	IsWithdrawalDetails()
}

// An auction duration is used to configure 3 auction periods:
// 1. `duration > 0`, `volume == 0`:
// The auction will last for at least N seconds.
// 2. `duration == 0`, `volume > 0`:
// The auction will end once the given volume will match at uncrossing.
// 3. `duration > 0`, `volume > 0`:
// The auction will take at least N seconds, but can end sooner if the market can trade a certain volume.
type AuctionDuration struct {
	// Duration of the auction in seconds
	DurationSecs int `json:"durationSecs"`
	// Target uncrossing trading volume
	Volume int `json:"volume"`
}

// A Vega builtin asset, mostly for testing purpose
type BuiltinAsset struct {
	// Maximum amount that can be requested by a party through the built-in asset faucet at a time
	MaxFaucetAmountMint string `json:"maxFaucetAmountMint"`
}

func (BuiltinAsset) IsAssetSource() {}

type BusEvent struct {
	// the ID for this event
	EventID string `json:"eventId"`
	// the block hash
	Block string `json:"block"`
	// the type of event
	Type BusEventType `json:"type"`
	// the payload - the wrapped event
	Event Event `json:"event"`
}

// A mode where Vega tries to execute orders as soon as they are received
type ContinuousTrading struct {
	// Size of an increment in price in terms of the quote currency
	TickSize string `json:"tickSize"`
}

// Frequent batch auctions trading mode
type DiscreteTrading struct {
	// Duration of the discrete trading batch in nanoseconds. Maximum 1 month.
	Duration int `json:"duration"`
	// Size of an increment in price in terms of the quote currency
	TickSize string `json:"tickSize"`
}

type DispatchStrategy struct {
	// Defines the data that will be used to compare markets so as to distribute rewards appropriately
	DispatchMetric vega.DispatchMetric `json:"dispatchMetric"`
	// The asset to use for measuring contribution to the metric
	DispatchMetricAssetID string `json:"dispatchMetricAssetId"`
	// Scope the dispatch to this market only under the metric asset
	MarketIdsInScope []string `json:"marketIdsInScope"`
}

// An asset originated from an Ethereum ERC20 Token
type Erc20 struct {
	// The address of the ERC20 contract
	ContractAddress string `json:"contractAddress"`
	// The lifetime limits deposit per address
	// Note: this is a temporary measure for alpha mainnet
	LifetimeLimit string `json:"lifetimeLimit"`
	// The maximum allowed per withdrawal
	// Note: this is a temporary measure for alpha mainnet
	WithdrawThreshold string `json:"withdrawThreshold"`
}

func (Erc20) IsAssetSource() {}

type EpochParticipation struct {
	Epoch *vega.Epoch `json:"epoch"`
	// RFC3339 timestamp
	Offline *string `json:"offline"`
	// RFC3339 timestamp
	Online       *string  `json:"online"`
	TotalRewards *float64 `json:"totalRewards"`
}

// All the data related to the approval of a withdrawal from the network
type Erc20WithdrawalApproval struct {
	// The source asset in the ethereum network
	AssetSource string `json:"assetSource"`
	// The amount to be withdrawn
	Amount string `json:"amount"`
	// Timestamp in seconds for expiry of the approval
	Expiry string `json:"expiry"`
	// The nonce to be used in the request
	Nonce string `json:"nonce"`
	// Signature aggregate from the nodes, in the following format:
	// 0x + sig1 + sig2 + ... + sigN
	Signatures string `json:"signatures"`
	// The target address which will receive the funds
	TargetAddress string `json:"targetAddress"`
	// Timestamp at which the withdrawal was created
	Creation string `json:"creation"`
}

// Specific details for an erc20 withdrawal
type Erc20WithdrawalDetails struct {
	// The ethereum address of the receiver of the asset funds
	ReceiverAddress string `json:"receiverAddress"`
}

func (Erc20WithdrawalDetails) IsWithdrawalDetails() {}

// An Ethereum oracle
type EthereumEvent struct {
	// The ID of the ethereum contract to use (string)
	ContractID string `json:"contractId"`
	// Name of the Ethereum event to listen to. (string)
	Event string `json:"event"`
}

func (EthereumEvent) IsOracle() {}

type LedgerEntry struct {
	// Account from which the asset was taken
	FromAccount string `json:"fromAccount"`
	// Account to which the balance was transferred
	ToAccount string `json:"toAccount"`
	// The amount transferred
	Amount string `json:"amount"`
	// The transfer reference
	Reference string `json:"reference"`
	// Type of ledger entry
	Type string `json:"type"`
	// RFC3339Nano time at which the transfer was made
	Timestamp string `json:"timestamp"`
}

// Configuration of a market liquidity monitoring parameters
type LiquidityMonitoringParameters struct {
	// Specifies parameters related to target stake calculation
	TargetStakeParameters *TargetStakeParameters `json:"targetStakeParameters"`
	// Specifies the triggering ratio for entering liquidity auction
	TriggeringRatio float64 `json:"triggeringRatio"`
}

// The equity like share of liquidity fee for each liquidity provider
type LiquidityProviderFeeShare struct {
	// The liquidity provider party ID
	Party *vega.Party `json:"party"`
	// The share owned by this liquidity provider (float)
	EquityLikeShare string `json:"equityLikeShare"`
	// The average entry valuation of the liquidity provider for the market
	AverageEntryValuation string `json:"averageEntryValuation"`
}

type LossSocialization struct {
	// the market ID where loss socialization happened
	MarketID string `json:"marketId"`
	// the party that was part of the loss socialization
	PartyID string `json:"partyId"`
	// the amount lost
	Amount string `json:"amount"`
}

func (LossSocialization) IsEvent() {}

// The liquidity commitments for this market
type MarketDataCommitments struct {
	// a set of liquidity sell orders to meet the liquidity provision obligation.
	Sells []*vega.LiquidityOrderReference `json:"sells"`
	// a set of liquidity buy orders to meet the liquidity provision obligation.
	Buys []*vega.LiquidityOrderReference `json:"buys"`
}

type MarketDepthTrade struct {
	// ID of the trade for the given market (if available)
	ID string `json:"id"`
	// Price of the trade
	Price string `json:"price"`
	// Size of the trade
	Size string `json:"size"`
}

type MarketEvent struct {
	// the market ID
	MarketID string `json:"marketId"`
	// the message - market events are used for logging
	Payload string `json:"payload"`
}

func (MarketEvent) IsEvent() {}

type MarketTick struct {
	// the market ID
	MarketID string `json:"marketId"`
	// the block time
	Time string `json:"time"`
}

func (MarketTick) IsEvent() {}

// The equity like share of liquidity fee for each liquidity provider
type ObservableLiquidityProviderFeeShare struct {
	// The liquidity provider party ID
	PartyID string `json:"partyId"`
	// The share owned by this liquidity provider (float)
	EquityLikeShare string `json:"equityLikeShare"`
	// The average entry valuation of the liquidity provider for the market
	AverageEntryValuation string `json:"averageEntryValuation"`
}

type OffsetPagination struct {
	// Skip the number of records specified, default is 0
	Skip int `json:"skip"`
	// Limit the number of returned records to the value specified, default is 50
	Limit int `json:"limit"`
	// Descending reverses the order of the records returned
	// default is true, if false the results will be returned in ascending order
	Descending bool `json:"descending"`
}

// An estimate of the fee to be paid by the order
type OrderEstimate struct {
	// The estimated fee if the order was to trade
	Fee *TradeFee `json:"fee"`
	// The total estimated amount of fee if the order was to trade
	TotalFeeAmount string `json:"totalFeeAmount"`
	// The margin requirement for this order
	MarginLevels *vega.MarginLevels `json:"marginLevels"`
}

type PositionResolution struct {
	// The market ID where position resolution happened
	MarketID string `json:"marketId"`
	// Number of distressed parties on market
	Distressed int `json:"distressed"`
	// Number of parties closed out
	Closed int `json:"closed"`
	// The mark price at which parties were distressed/closed out
	MarkPrice string `json:"markPrice"`
}

func (PositionResolution) IsEvent() {}

// Range of valid prices and the associated price monitoring trigger
type PriceMonitoringBounds struct {
	// Minimum price that isn't currently breaching the specified price monitoring trigger
	MinValidPrice string `json:"minValidPrice"`
	// Maximum price that isn't currently breaching the specified price monitoring trigger
	MaxValidPrice string `json:"maxValidPrice"`
	// Price monitoring trigger associated with the bounds
	Trigger *PriceMonitoringTrigger `json:"trigger"`
	// Reference price used to calculate the valid price range
	ReferencePrice string `json:"referencePrice"`
}

// PriceMonitoringParameters holds a list of triggers
type PriceMonitoringParameters struct {
	// The list of triggers for this price monitoring
	Triggers []*PriceMonitoringTrigger `json:"triggers"`
}

// Configuration of a market price monitoring auctions triggers
type PriceMonitoringSettings struct {
	// Specified a set of PriceMonitoringParameters to be use for price monitoring purposes
	Parameters *PriceMonitoringParameters `json:"parameters"`
}

// PriceMonitoringParameters holds together price projection horizon τ, probability level p, and auction extension duration
type PriceMonitoringTrigger struct {
	// Price monitoring projection horizon τ in seconds (> 0).
	HorizonSecs int `json:"horizonSecs"`
	// Price monitoring probability level p. (>0 and < 1)
	Probability float64 `json:"probability"`
	// Price monitoring auction extension duration in seconds should the price
	// breach its theoretical level over the specified horizon at the specified
	// probability level (> 0)
	AuctionExtensionSecs int `json:"auctionExtensionSecs"`
}

type ProposalVote struct {
	// Cast vote
	Vote *vega.Vote `json:"vote"`
	// Proposal the vote is cast on
	ProposalID string `json:"proposalId"`
}

type ProposalVoteConnection struct {
	// The proposal votes in this connection
	Edges []*ProposalVoteEdge `json:"edges"`
	// The pagination information
	PageInfo *v2.PageInfo `json:"pageInfo"`
}

type ProposalVoteEdge struct {
	Node   *ProposalVote `json:"node"`
	Cursor *string       `json:"cursor"`
}

type ProposalVoteSide struct {
	// All votes cast for this side
	Votes []*vega.Vote `json:"votes"`
	// Total number of votes cast for this side
	TotalNumber string `json:"totalNumber"`
	// Total weight of governance token from the votes cast for this side
	TotalWeight string `json:"totalWeight"`
	// Total number of governance tokens from the votes cast for this side
	TotalTokens string `json:"totalTokens"`
}

type ProposalVotes struct {
	// Yes votes cast for this proposal
	Yes *ProposalVoteSide `json:"yes"`
	// No votes cast for this proposal
	No *ProposalVoteSide `json:"no"`
}

type RewardSummaryConnection struct {
	// List of reward summaries available for the connection
	Edges []*RewardSummaryEdge `json:"edges"`
	// Page information for the connection
	PageInfo *v2.PageInfo `json:"pageInfo"`
}

type RewardSummaryEdge struct {
	// The reward summary
	Node *vega.RewardSummary `json:"node"`
	// Cursor identifying the reward summary
	Cursor string `json:"cursor"`
}

type SettleDistressed struct {
	// the market in which a position was closed out
	MarketID string `json:"marketId"`
	// the party that was closed out
	PartyID string `json:"partyId"`
	// the margin taken from distressed party
	Margin string `json:"margin"`
	// the price at which the position was closed out
	Price string `json:"price"`
}

func (SettleDistressed) IsEvent() {}

type SettlePosition struct {
	// the market in which a position was settled
	MarketID string `json:"marketId"`
	// the party who settled a position
	PartyID string `json:"partyId"`
	// the settle price
	Price string `json:"price"`
	// the trades that were settled to close the overall position
	TradeSettlements []*TradeSettlement `json:"tradeSettlements"`
}

func (SettlePosition) IsEvent() {}

// All staking information related to a Party.
// Contains the current recognised balance by the network and
// all the StakeLink/Unlink seen by the network
type StakingSummary struct {
	// The stake currently available for the party
	CurrentStakeAvailable string `json:"currentStakeAvailable"`
	// The list of all stake link/unlink for the party
	Linkings *v2.StakesConnection `json:"linkings"`
}

// TargetStakeParameters contains parameters used in target stake calculation
type TargetStakeParameters struct {
	// Specifies length of time window expressed in seconds for target stake calculation
	TimeWindow int `json:"timeWindow"`
	// Specifies scaling factors used in target stake calculation
	ScalingFactor float64 `json:"scalingFactor"`
}

type TimeUpdate struct {
	// RFC3339Nano time of new block time
	Timestamp string `json:"timestamp"`
}

func (TimeUpdate) IsEvent() {}

// The fee paid by the party when a trade occurs
type TradeFee struct {
	// The maker fee, paid by the aggressive party to the other party (the one who had an order in the book)
	MakerFee string `json:"makerFee"`
	// The infrastructure fee, a fee paid to the validators to maintain the Vega network
	InfrastructureFee string `json:"infrastructureFee"`
	// The fee paid to the liquidity providers that committed liquidity to the market
	LiquidityFee string `json:"liquidityFee"`
}

type TradeSettlement struct {
	// the size of the trade
	Size int `json:"size"`
	// the price of the trade
	Price string `json:"price"`
}

type TransactionSubmitted struct {
	Success bool `json:"success"`
}

type TransferBalance struct {
	// Account involved in transfer
	Account *vega.Account `json:"account"`
	// The new balance of the account
	Balance string `json:"balance"`
}

type TransferResponse struct {
	// The ledger entries and balances resulting from a transfer request
	Transfers []*LedgerEntry `json:"transfers"`
	// The balances of accounts involved in the transfer
	Balances []*TransferBalance `json:"balances"`
}

type TransferResponses struct {
	// A group of transfer responses - events from core
	Responses []*TransferResponse `json:"responses"`
}

func (TransferResponses) IsEvent() {}

// An asset originated from an Ethereum ERC20 Token
type UpdateErc20 struct {
	// The lifetime limits deposit per address
	// Note: this is a temporary measure for alpha mainnet
	LifetimeLimit string `json:"lifetimeLimit"`
	// The maximum allowed per withdrawal
	// Note: this is a temporary measure for alpha mainnet
	WithdrawThreshold string `json:"withdrawThreshold"`
}

func (UpdateErc20) IsUpdateAssetSource() {}

type UpdateInstrumentConfiguration struct {
	Code    string                    `json:"code"`
	Product *vega.UpdateFutureProduct `json:"product"`
}

type BusEventType string

const (
	// Vega Time has changed
	BusEventTypeTimeUpdate BusEventType = "TimeUpdate"
	// A balance has been transferred between accounts
	BusEventTypeTransferResponses BusEventType = "TransferResponses"
	// A position resolution event has occurred
	BusEventTypePositionResolution BusEventType = "PositionResolution"
	// An order has been created or updated
	BusEventTypeOrder BusEventType = "Order"
	// An account has been updated
	BusEventTypeAccount BusEventType = "Account"
	// A party has been updated
	BusEventTypeParty BusEventType = "Party"
	// A trade has been created
	BusEventTypeTrade BusEventType = "Trade"
	// Margin levels have changed for a position
	BusEventTypeMarginLevels BusEventType = "MarginLevels"
	// A governance proposal has been created or updated
	BusEventTypeProposal BusEventType = "Proposal"
	// A vote has been placed on a governance proposal
	BusEventTypeVote BusEventType = "Vote"
	// Market data has been updated
	BusEventTypeMarketData BusEventType = "MarketData"
	// Validator node signatures for an event
	BusEventTypeNodeSignature BusEventType = "NodeSignature"
	// A position has been closed without sufficient insurance pool balance to cover it
	BusEventTypeLossSocialization BusEventType = "LossSocialization"
	// A position has been settled
	BusEventTypeSettlePosition BusEventType = "SettlePosition"
	// A distressed position has been settled
	BusEventTypeSettleDistressed BusEventType = "SettleDistressed"
	// A new market has been created
	BusEventTypeMarketCreated BusEventType = "MarketCreated"
	// A market has been updated
	BusEventTypeMarketUpdated BusEventType = "MarketUpdated"
	// An asset has been created or update
	BusEventTypeAsset BusEventType = "Asset"
	// A market has progressed by one tick
	BusEventTypeMarketTick BusEventType = "MarketTick"
	// A market has either entered or exited auction
	BusEventTypeAuction BusEventType = "Auction"
	// A risk factor adjustment was made
	BusEventTypeRiskFactor BusEventType = "RiskFactor"
	// A liquidity commitment change occurred
	BusEventTypeLiquidityProvision BusEventType = "LiquidityProvision"
	// Collateral has deposited in to this Vega network via the bridge
	BusEventTypeDeposit BusEventType = "Deposit"
	// Collateral has been withdrawn from this Vega network via the bridge
	BusEventTypeWithdrawal BusEventType = "Withdrawal"
	// An oracle spec has been registered
	BusEventTypeOracleSpec BusEventType = "OracleSpec"
	// constant for market events - mainly used for logging
	BusEventTypeMarket BusEventType = "Market"
)

var AllBusEventType = []BusEventType{
	BusEventTypeTimeUpdate,
	BusEventTypeTransferResponses,
	BusEventTypePositionResolution,
	BusEventTypeOrder,
	BusEventTypeAccount,
	BusEventTypeParty,
	BusEventTypeTrade,
	BusEventTypeMarginLevels,
	BusEventTypeProposal,
	BusEventTypeVote,
	BusEventTypeMarketData,
	BusEventTypeNodeSignature,
	BusEventTypeLossSocialization,
	BusEventTypeSettlePosition,
	BusEventTypeSettleDistressed,
	BusEventTypeMarketCreated,
	BusEventTypeMarketUpdated,
	BusEventTypeAsset,
	BusEventTypeMarketTick,
	BusEventTypeAuction,
	BusEventTypeRiskFactor,
	BusEventTypeLiquidityProvision,
	BusEventTypeDeposit,
	BusEventTypeWithdrawal,
	BusEventTypeOracleSpec,
	BusEventTypeMarket,
}

func (e BusEventType) IsValid() bool {
	switch e {
	case BusEventTypeTimeUpdate, BusEventTypeTransferResponses, BusEventTypePositionResolution, BusEventTypeOrder, BusEventTypeAccount, BusEventTypeParty, BusEventTypeTrade, BusEventTypeMarginLevels, BusEventTypeProposal, BusEventTypeVote, BusEventTypeMarketData, BusEventTypeNodeSignature, BusEventTypeLossSocialization, BusEventTypeSettlePosition, BusEventTypeSettleDistressed, BusEventTypeMarketCreated, BusEventTypeMarketUpdated, BusEventTypeAsset, BusEventTypeMarketTick, BusEventTypeAuction, BusEventTypeRiskFactor, BusEventTypeLiquidityProvision, BusEventTypeDeposit, BusEventTypeWithdrawal, BusEventTypeOracleSpec, BusEventTypeMarket:
		return true
	}
	return false
}

func (e BusEventType) String() string {
	return string(e)
}

func (e *BusEventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BusEventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BusEventType", str)
	}
	return nil
}

func (e BusEventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TransferDirection string

const (
	TransferDirectionTo       TransferDirection = "To"
	TransferDirectionFrom     TransferDirection = "From"
	TransferDirectionToOrFrom TransferDirection = "ToOrFrom"
)

var AllTransferDirection = []TransferDirection{
	TransferDirectionTo,
	TransferDirectionFrom,
	TransferDirectionToOrFrom,
}

func (e TransferDirection) IsValid() bool {
	switch e {
	case TransferDirectionTo, TransferDirectionFrom, TransferDirectionToOrFrom:
		return true
	}
	return false
}

func (e TransferDirection) String() string {
	return string(e)
}

func (e *TransferDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransferDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransferDirection", str)
	}
	return nil
}

func (e TransferDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
