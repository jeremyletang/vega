// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"fmt"
	"io"
	"strconv"

	"code.vegaprotocol.io/vega/proto"
)

type Oracle interface {
	IsOracle()
}

type Product interface {
	IsProduct()
}

type RiskModel interface {
	IsRiskModel()
}

type TradingMode interface {
	IsTradingMode()
}

type CheckTokenResponse struct {
	Ok bool `json:"ok"`
}

type ContinuousTrading struct {
	TickSize *int `json:"tickSize"`
}

func (ContinuousTrading) IsTradingMode() {}

type DiscreteTrading struct {
	Duration *int `json:"duration"`
}

func (DiscreteTrading) IsTradingMode() {}

type EthereumEvent struct {
	ContractID string `json:"contractId"`
	Event      string `json:"event"`
}

func (EthereumEvent) IsOracle() {}

type Forward struct {
	Lambd  float64        `json:"lambd"`
	Tau    float64        `json:"Tau"`
	Params *ModelParamsBs `json:"params"`
}

func (Forward) IsRiskModel() {}

type Future struct {
	Maturity string `json:"maturity"`
	Asset    string `json:"asset"`
	Oracle   Oracle `json:"oracle"`
}

func (Future) IsProduct() {}

type Instrument struct {
	ID        string              `json:"id"`
	Code      string              `json:"code"`
	Name      string              `json:"name"`
	BaseName  string              `json:"baseName"`
	QuoteName string              `json:"quoteName"`
	Metadata  *InstrumentMetadata `json:"metadata"`
	Product   Product             `json:"product"`
}

type InstrumentMetadata struct {
	Tags []*string `json:"tags"`
}

type Market struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	TradableInstrument *TradableInstrument `json:"tradableInstrument"`
	TradingMode        TradingMode         `json:"tradingMode"`
	DecimalPlaces      int                 `json:"decimalPlaces"`
	Orders             []*proto.Order      `json:"orders"`
	Accounts           []*proto.Account    `json:"accounts"`
	Trades             []*proto.Trade      `json:"trades"`
	Depth              *proto.MarketDepth  `json:"depth"`
	Candles            []*proto.Candle     `json:"candles"`
	OrderByReference   *proto.Order        `json:"orderByReference"`
}

type ModelParamsBs struct {
	Mu    float64 `json:"mu"`
	R     float64 `json:"r"`
	Sigma float64 `json:"sigma"`
}

type Party struct {
	ID        string                  `json:"id"`
	Orders    []*proto.Order          `json:"orders"`
	Trades    []*proto.Trade          `json:"trades"`
	Accounts  []*proto.Account        `json:"accounts"`
	Positions []*proto.MarketPosition `json:"positions"`
}

type SimpleRiskModel struct {
	Params *SimpleRiskModelParams `json:"params"`
}

func (SimpleRiskModel) IsRiskModel() {}

type SimpleRiskModelParams struct {
	FactorLong  float64 `json:"factorLong"`
	FactorShort float64 `json:"factorShort"`
}

type TradableInstrument struct {
	Instrument *Instrument `json:"instrument"`
	RiskModel  RiskModel   `json:"riskModel"`
}

type AccountType string

const (
	AccountTypeInsurance  AccountType = "Insurance"
	AccountTypeSettlement AccountType = "Settlement"
	AccountTypeMargin     AccountType = "Margin"
	AccountTypeGeneral    AccountType = "General"
)

var AllAccountType = []AccountType{
	AccountTypeInsurance,
	AccountTypeSettlement,
	AccountTypeMargin,
	AccountTypeGeneral,
}

func (e AccountType) IsValid() bool {
	switch e {
	case AccountTypeInsurance, AccountTypeSettlement, AccountTypeMargin, AccountTypeGeneral:
		return true
	}
	return false
}

func (e AccountType) String() string {
	return string(e)
}

func (e *AccountType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccountType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccountType", str)
	}
	return nil
}

func (e AccountType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Interval string

const (
	IntervalI1m  Interval = "I1M"
	IntervalI5m  Interval = "I5M"
	IntervalI15m Interval = "I15M"
	IntervalI1h  Interval = "I1H"
	IntervalI6h  Interval = "I6H"
	IntervalI1d  Interval = "I1D"
)

var AllInterval = []Interval{
	IntervalI1m,
	IntervalI5m,
	IntervalI15m,
	IntervalI1h,
	IntervalI6h,
	IntervalI1d,
}

func (e Interval) IsValid() bool {
	switch e {
	case IntervalI1m, IntervalI5m, IntervalI15m, IntervalI1h, IntervalI6h, IntervalI1d:
		return true
	}
	return false
}

func (e Interval) String() string {
	return string(e)
}

func (e *Interval) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Interval(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Interval", str)
	}
	return nil
}

func (e Interval) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderStatus string

const (
	OrderStatusActive    OrderStatus = "Active"
	OrderStatusCancelled OrderStatus = "Cancelled"
	OrderStatusExpired   OrderStatus = "Expired"
	OrderStatusStopped   OrderStatus = "Stopped"
	OrderStatusFilled    OrderStatus = "Filled"
)

var AllOrderStatus = []OrderStatus{
	OrderStatusActive,
	OrderStatusCancelled,
	OrderStatusExpired,
	OrderStatusStopped,
	OrderStatusFilled,
}

func (e OrderStatus) IsValid() bool {
	switch e {
	case OrderStatusActive, OrderStatusCancelled, OrderStatusExpired, OrderStatusStopped, OrderStatusFilled:
		return true
	}
	return false
}

func (e OrderStatus) String() string {
	return string(e)
}

func (e *OrderStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderStatus", str)
	}
	return nil
}

func (e OrderStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderTimeInForce string

const (
	OrderTimeInForceFok OrderTimeInForce = "FOK"
	OrderTimeInForceIoc OrderTimeInForce = "IOC"
	OrderTimeInForceGtc OrderTimeInForce = "GTC"
	OrderTimeInForceGtt OrderTimeInForce = "GTT"
)

var AllOrderTimeInForce = []OrderTimeInForce{
	OrderTimeInForceFok,
	OrderTimeInForceIoc,
	OrderTimeInForceGtc,
	OrderTimeInForceGtt,
}

func (e OrderTimeInForce) IsValid() bool {
	switch e {
	case OrderTimeInForceFok, OrderTimeInForceIoc, OrderTimeInForceGtc, OrderTimeInForceGtt:
		return true
	}
	return false
}

func (e OrderTimeInForce) String() string {
	return string(e)
}

func (e *OrderTimeInForce) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderTimeInForce(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderTimeInForce", str)
	}
	return nil
}

func (e OrderTimeInForce) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderType string

const (
	OrderTypeMarket  OrderType = "MARKET"
	OrderTypeLimit   OrderType = "LIMIT"
	OrderTypeNetwork OrderType = "NETWORK"
)

var AllOrderType = []OrderType{
	OrderTypeMarket,
	OrderTypeLimit,
	OrderTypeNetwork,
}

func (e OrderType) IsValid() bool {
	switch e {
	case OrderTypeMarket, OrderTypeLimit, OrderTypeNetwork:
		return true
	}
	return false
}

func (e OrderType) String() string {
	return string(e)
}

func (e *OrderType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderType", str)
	}
	return nil
}

func (e OrderType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Side string

const (
	SideBuy  Side = "Buy"
	SideSell Side = "Sell"
)

var AllSide = []Side{
	SideBuy,
	SideSell,
}

func (e Side) IsValid() bool {
	switch e {
	case SideBuy, SideSell:
		return true
	}
	return false
}

func (e Side) String() string {
	return string(e)
}

func (e *Side) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Side(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Side", str)
	}
	return nil
}

func (e Side) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ValueDirection string

const (
	ValueDirectionPositive ValueDirection = "Positive"
	ValueDirectionNegative ValueDirection = "Negative"
)

var AllValueDirection = []ValueDirection{
	ValueDirectionPositive,
	ValueDirectionNegative,
}

func (e ValueDirection) IsValid() bool {
	switch e {
	case ValueDirectionPositive, ValueDirectionNegative:
		return true
	}
	return false
}

func (e ValueDirection) String() string {
	return string(e)
}

func (e *ValueDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ValueDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ValueDirection", str)
	}
	return nil
}

func (e ValueDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
