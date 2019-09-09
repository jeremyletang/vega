// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/markets.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ContinuousTrading struct {
	TickSize             uint64   `protobuf:"varint,1,opt,name=tickSize,proto3" json:"tickSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContinuousTrading) Reset()         { *m = ContinuousTrading{} }
func (m *ContinuousTrading) String() string { return proto.CompactTextString(m) }
func (*ContinuousTrading) ProtoMessage()    {}
func (*ContinuousTrading) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{0}
}

func (m *ContinuousTrading) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContinuousTrading.Unmarshal(m, b)
}
func (m *ContinuousTrading) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContinuousTrading.Marshal(b, m, deterministic)
}
func (m *ContinuousTrading) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousTrading.Merge(m, src)
}
func (m *ContinuousTrading) XXX_Size() int {
	return xxx_messageInfo_ContinuousTrading.Size(m)
}
func (m *ContinuousTrading) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousTrading.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousTrading proto.InternalMessageInfo

func (m *ContinuousTrading) GetTickSize() uint64 {
	if m != nil {
		return m.TickSize
	}
	return 0
}

type DiscreteTrading struct {
	Duration             int64    `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscreteTrading) Reset()         { *m = DiscreteTrading{} }
func (m *DiscreteTrading) String() string { return proto.CompactTextString(m) }
func (*DiscreteTrading) ProtoMessage()    {}
func (*DiscreteTrading) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{1}
}

func (m *DiscreteTrading) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscreteTrading.Unmarshal(m, b)
}
func (m *DiscreteTrading) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscreteTrading.Marshal(b, m, deterministic)
}
func (m *DiscreteTrading) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscreteTrading.Merge(m, src)
}
func (m *DiscreteTrading) XXX_Size() int {
	return xxx_messageInfo_DiscreteTrading.Size(m)
}
func (m *DiscreteTrading) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscreteTrading.DiscardUnknown(m)
}

var xxx_messageInfo_DiscreteTrading proto.InternalMessageInfo

func (m *DiscreteTrading) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type Future struct {
	Maturity string `protobuf:"bytes,1,opt,name=maturity,proto3" json:"maturity,omitempty"`
	Asset    string `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	// Types that are valid to be assigned to Oracle:
	//	*Future_EthereumEvent
	Oracle               isFuture_Oracle `protobuf_oneof:"oracle"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Future) Reset()         { *m = Future{} }
func (m *Future) String() string { return proto.CompactTextString(m) }
func (*Future) ProtoMessage()    {}
func (*Future) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{2}
}

func (m *Future) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Future.Unmarshal(m, b)
}
func (m *Future) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Future.Marshal(b, m, deterministic)
}
func (m *Future) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Future.Merge(m, src)
}
func (m *Future) XXX_Size() int {
	return xxx_messageInfo_Future.Size(m)
}
func (m *Future) XXX_DiscardUnknown() {
	xxx_messageInfo_Future.DiscardUnknown(m)
}

var xxx_messageInfo_Future proto.InternalMessageInfo

func (m *Future) GetMaturity() string {
	if m != nil {
		return m.Maturity
	}
	return ""
}

func (m *Future) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

type isFuture_Oracle interface {
	isFuture_Oracle()
}

type Future_EthereumEvent struct {
	EthereumEvent *EthereumEvent `protobuf:"bytes,100,opt,name=ethereumEvent,proto3,oneof"`
}

func (*Future_EthereumEvent) isFuture_Oracle() {}

func (m *Future) GetOracle() isFuture_Oracle {
	if m != nil {
		return m.Oracle
	}
	return nil
}

func (m *Future) GetEthereumEvent() *EthereumEvent {
	if x, ok := m.GetOracle().(*Future_EthereumEvent); ok {
		return x.EthereumEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Future) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Future_EthereumEvent)(nil),
	}
}

type EthereumEvent struct {
	ContractID           string   `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthereumEvent) Reset()         { *m = EthereumEvent{} }
func (m *EthereumEvent) String() string { return proto.CompactTextString(m) }
func (*EthereumEvent) ProtoMessage()    {}
func (*EthereumEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{3}
}

func (m *EthereumEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthereumEvent.Unmarshal(m, b)
}
func (m *EthereumEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthereumEvent.Marshal(b, m, deterministic)
}
func (m *EthereumEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumEvent.Merge(m, src)
}
func (m *EthereumEvent) XXX_Size() int {
	return xxx_messageInfo_EthereumEvent.Size(m)
}
func (m *EthereumEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumEvent proto.InternalMessageInfo

func (m *EthereumEvent) GetContractID() string {
	if m != nil {
		return m.ContractID
	}
	return ""
}

func (m *EthereumEvent) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

type InstrumentMetadata struct {
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstrumentMetadata) Reset()         { *m = InstrumentMetadata{} }
func (m *InstrumentMetadata) String() string { return proto.CompactTextString(m) }
func (*InstrumentMetadata) ProtoMessage()    {}
func (*InstrumentMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{4}
}

func (m *InstrumentMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstrumentMetadata.Unmarshal(m, b)
}
func (m *InstrumentMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstrumentMetadata.Marshal(b, m, deterministic)
}
func (m *InstrumentMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentMetadata.Merge(m, src)
}
func (m *InstrumentMetadata) XXX_Size() int {
	return xxx_messageInfo_InstrumentMetadata.Size(m)
}
func (m *InstrumentMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentMetadata proto.InternalMessageInfo

func (m *InstrumentMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Instrument struct {
	Id        string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code      string              `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name      string              `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	BaseName  string              `protobuf:"bytes,4,opt,name=baseName,proto3" json:"baseName,omitempty"`
	QuoteName string              `protobuf:"bytes,5,opt,name=quoteName,proto3" json:"quoteName,omitempty"`
	Metadata  *InstrumentMetadata `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Types that are valid to be assigned to Product:
	//	*Instrument_Future
	Product              isInstrument_Product `protobuf_oneof:"product"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Instrument) Reset()         { *m = Instrument{} }
func (m *Instrument) String() string { return proto.CompactTextString(m) }
func (*Instrument) ProtoMessage()    {}
func (*Instrument) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{5}
}

func (m *Instrument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instrument.Unmarshal(m, b)
}
func (m *Instrument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instrument.Marshal(b, m, deterministic)
}
func (m *Instrument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instrument.Merge(m, src)
}
func (m *Instrument) XXX_Size() int {
	return xxx_messageInfo_Instrument.Size(m)
}
func (m *Instrument) XXX_DiscardUnknown() {
	xxx_messageInfo_Instrument.DiscardUnknown(m)
}

var xxx_messageInfo_Instrument proto.InternalMessageInfo

func (m *Instrument) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Instrument) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Instrument) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instrument) GetBaseName() string {
	if m != nil {
		return m.BaseName
	}
	return ""
}

func (m *Instrument) GetQuoteName() string {
	if m != nil {
		return m.QuoteName
	}
	return ""
}

func (m *Instrument) GetMetadata() *InstrumentMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type isInstrument_Product interface {
	isInstrument_Product()
}

type Instrument_Future struct {
	Future *Future `protobuf:"bytes,100,opt,name=future,proto3,oneof"`
}

func (*Instrument_Future) isInstrument_Product() {}

func (m *Instrument) GetProduct() isInstrument_Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *Instrument) GetFuture() *Future {
	if x, ok := m.GetProduct().(*Instrument_Future); ok {
		return x.Future
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Instrument) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Instrument_Future)(nil),
	}
}

type ModelParamsBS struct {
	Mu                   float64  `protobuf:"fixed64,1,opt,name=mu,proto3" json:"mu,omitempty"`
	R                    float64  `protobuf:"fixed64,2,opt,name=r,proto3" json:"r,omitempty"`
	Sigma                float64  `protobuf:"fixed64,3,opt,name=sigma,proto3" json:"sigma,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelParamsBS) Reset()         { *m = ModelParamsBS{} }
func (m *ModelParamsBS) String() string { return proto.CompactTextString(m) }
func (*ModelParamsBS) ProtoMessage()    {}
func (*ModelParamsBS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{6}
}

func (m *ModelParamsBS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelParamsBS.Unmarshal(m, b)
}
func (m *ModelParamsBS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelParamsBS.Marshal(b, m, deterministic)
}
func (m *ModelParamsBS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelParamsBS.Merge(m, src)
}
func (m *ModelParamsBS) XXX_Size() int {
	return xxx_messageInfo_ModelParamsBS.Size(m)
}
func (m *ModelParamsBS) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelParamsBS.DiscardUnknown(m)
}

var xxx_messageInfo_ModelParamsBS proto.InternalMessageInfo

func (m *ModelParamsBS) GetMu() float64 {
	if m != nil {
		return m.Mu
	}
	return 0
}

func (m *ModelParamsBS) GetR() float64 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *ModelParamsBS) GetSigma() float64 {
	if m != nil {
		return m.Sigma
	}
	return 0
}

type Forward struct {
	Lambd                float64        `protobuf:"fixed64,1,opt,name=lambd,proto3" json:"lambd,omitempty"`
	Tau                  float64        `protobuf:"fixed64,2,opt,name=tau,proto3" json:"tau,omitempty"`
	Params               *ModelParamsBS `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Forward) Reset()         { *m = Forward{} }
func (m *Forward) String() string { return proto.CompactTextString(m) }
func (*Forward) ProtoMessage()    {}
func (*Forward) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{7}
}

func (m *Forward) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forward.Unmarshal(m, b)
}
func (m *Forward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forward.Marshal(b, m, deterministic)
}
func (m *Forward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forward.Merge(m, src)
}
func (m *Forward) XXX_Size() int {
	return xxx_messageInfo_Forward.Size(m)
}
func (m *Forward) XXX_DiscardUnknown() {
	xxx_messageInfo_Forward.DiscardUnknown(m)
}

var xxx_messageInfo_Forward proto.InternalMessageInfo

func (m *Forward) GetLambd() float64 {
	if m != nil {
		return m.Lambd
	}
	return 0
}

func (m *Forward) GetTau() float64 {
	if m != nil {
		return m.Tau
	}
	return 0
}

func (m *Forward) GetParams() *ModelParamsBS {
	if m != nil {
		return m.Params
	}
	return nil
}

type ExternalRiskModel struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Socket               string            `protobuf:"bytes,2,opt,name=socket,proto3" json:"socket,omitempty"`
	Config               map[string]string `protobuf:"bytes,3,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExternalRiskModel) Reset()         { *m = ExternalRiskModel{} }
func (m *ExternalRiskModel) String() string { return proto.CompactTextString(m) }
func (*ExternalRiskModel) ProtoMessage()    {}
func (*ExternalRiskModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{8}
}

func (m *ExternalRiskModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalRiskModel.Unmarshal(m, b)
}
func (m *ExternalRiskModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalRiskModel.Marshal(b, m, deterministic)
}
func (m *ExternalRiskModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalRiskModel.Merge(m, src)
}
func (m *ExternalRiskModel) XXX_Size() int {
	return xxx_messageInfo_ExternalRiskModel.Size(m)
}
func (m *ExternalRiskModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalRiskModel.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalRiskModel proto.InternalMessageInfo

func (m *ExternalRiskModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExternalRiskModel) GetSocket() string {
	if m != nil {
		return m.Socket
	}
	return ""
}

func (m *ExternalRiskModel) GetConfig() map[string]string {
	if m != nil {
		return m.Config
	}
	return nil
}

type ScalingFactors struct {
	SearchLevel          float64  `protobuf:"fixed64,1,opt,name=searchLevel,proto3" json:"searchLevel,omitempty"`
	InitialMargin        float64  `protobuf:"fixed64,2,opt,name=initialMargin,proto3" json:"initialMargin,omitempty"`
	CollateralRelease    float64  `protobuf:"fixed64,3,opt,name=collateralRelease,proto3" json:"collateralRelease,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScalingFactors) Reset()         { *m = ScalingFactors{} }
func (m *ScalingFactors) String() string { return proto.CompactTextString(m) }
func (*ScalingFactors) ProtoMessage()    {}
func (*ScalingFactors) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{9}
}

func (m *ScalingFactors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScalingFactors.Unmarshal(m, b)
}
func (m *ScalingFactors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScalingFactors.Marshal(b, m, deterministic)
}
func (m *ScalingFactors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScalingFactors.Merge(m, src)
}
func (m *ScalingFactors) XXX_Size() int {
	return xxx_messageInfo_ScalingFactors.Size(m)
}
func (m *ScalingFactors) XXX_DiscardUnknown() {
	xxx_messageInfo_ScalingFactors.DiscardUnknown(m)
}

var xxx_messageInfo_ScalingFactors proto.InternalMessageInfo

func (m *ScalingFactors) GetSearchLevel() float64 {
	if m != nil {
		return m.SearchLevel
	}
	return 0
}

func (m *ScalingFactors) GetInitialMargin() float64 {
	if m != nil {
		return m.InitialMargin
	}
	return 0
}

func (m *ScalingFactors) GetCollateralRelease() float64 {
	if m != nil {
		return m.CollateralRelease
	}
	return 0
}

type MarginCalculator struct {
	ScalingFactors       *ScalingFactors `protobuf:"bytes,1,opt,name=scalingFactors,proto3" json:"scalingFactors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MarginCalculator) Reset()         { *m = MarginCalculator{} }
func (m *MarginCalculator) String() string { return proto.CompactTextString(m) }
func (*MarginCalculator) ProtoMessage()    {}
func (*MarginCalculator) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{10}
}

func (m *MarginCalculator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarginCalculator.Unmarshal(m, b)
}
func (m *MarginCalculator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarginCalculator.Marshal(b, m, deterministic)
}
func (m *MarginCalculator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarginCalculator.Merge(m, src)
}
func (m *MarginCalculator) XXX_Size() int {
	return xxx_messageInfo_MarginCalculator.Size(m)
}
func (m *MarginCalculator) XXX_DiscardUnknown() {
	xxx_messageInfo_MarginCalculator.DiscardUnknown(m)
}

var xxx_messageInfo_MarginCalculator proto.InternalMessageInfo

func (m *MarginCalculator) GetScalingFactors() *ScalingFactors {
	if m != nil {
		return m.ScalingFactors
	}
	return nil
}

type TradableInstrument struct {
	Instrument       *Instrument       `protobuf:"bytes,1,opt,name=instrument,proto3" json:"instrument,omitempty"`
	MarginCalculator *MarginCalculator `protobuf:"bytes,2,opt,name=marginCalculator,proto3" json:"marginCalculator,omitempty"`
	// Types that are valid to be assigned to RiskModel:
	//	*TradableInstrument_Forward
	//	*TradableInstrument_ExternalRiskModel
	RiskModel            isTradableInstrument_RiskModel `protobuf_oneof:"riskModel"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TradableInstrument) Reset()         { *m = TradableInstrument{} }
func (m *TradableInstrument) String() string { return proto.CompactTextString(m) }
func (*TradableInstrument) ProtoMessage()    {}
func (*TradableInstrument) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{11}
}

func (m *TradableInstrument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradableInstrument.Unmarshal(m, b)
}
func (m *TradableInstrument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradableInstrument.Marshal(b, m, deterministic)
}
func (m *TradableInstrument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradableInstrument.Merge(m, src)
}
func (m *TradableInstrument) XXX_Size() int {
	return xxx_messageInfo_TradableInstrument.Size(m)
}
func (m *TradableInstrument) XXX_DiscardUnknown() {
	xxx_messageInfo_TradableInstrument.DiscardUnknown(m)
}

var xxx_messageInfo_TradableInstrument proto.InternalMessageInfo

func (m *TradableInstrument) GetInstrument() *Instrument {
	if m != nil {
		return m.Instrument
	}
	return nil
}

func (m *TradableInstrument) GetMarginCalculator() *MarginCalculator {
	if m != nil {
		return m.MarginCalculator
	}
	return nil
}

type isTradableInstrument_RiskModel interface {
	isTradableInstrument_RiskModel()
}

type TradableInstrument_Forward struct {
	Forward *Forward `protobuf:"bytes,100,opt,name=forward,proto3,oneof"`
}

type TradableInstrument_ExternalRiskModel struct {
	ExternalRiskModel *ExternalRiskModel `protobuf:"bytes,101,opt,name=externalRiskModel,proto3,oneof"`
}

func (*TradableInstrument_Forward) isTradableInstrument_RiskModel() {}

func (*TradableInstrument_ExternalRiskModel) isTradableInstrument_RiskModel() {}

func (m *TradableInstrument) GetRiskModel() isTradableInstrument_RiskModel {
	if m != nil {
		return m.RiskModel
	}
	return nil
}

func (m *TradableInstrument) GetForward() *Forward {
	if x, ok := m.GetRiskModel().(*TradableInstrument_Forward); ok {
		return x.Forward
	}
	return nil
}

func (m *TradableInstrument) GetExternalRiskModel() *ExternalRiskModel {
	if x, ok := m.GetRiskModel().(*TradableInstrument_ExternalRiskModel); ok {
		return x.ExternalRiskModel
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TradableInstrument) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TradableInstrument_Forward)(nil),
		(*TradableInstrument_ExternalRiskModel)(nil),
	}
}

type Market struct {
	Id                 string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TradableInstrument *TradableInstrument `protobuf:"bytes,3,opt,name=tradableInstrument,proto3" json:"tradableInstrument,omitempty"`
	DecimalPlaces      uint64              `protobuf:"varint,4,opt,name=decimalPlaces,proto3" json:"decimalPlaces,omitempty"`
	// Types that are valid to be assigned to TradingMode:
	//	*Market_Continuous
	//	*Market_Discrete
	TradingMode          isMarket_TradingMode `protobuf_oneof:"tradingMode"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{12}
}

func (m *Market) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Market.Unmarshal(m, b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Market.Marshal(b, m, deterministic)
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return xxx_messageInfo_Market.Size(m)
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

func (m *Market) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Market) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Market) GetTradableInstrument() *TradableInstrument {
	if m != nil {
		return m.TradableInstrument
	}
	return nil
}

func (m *Market) GetDecimalPlaces() uint64 {
	if m != nil {
		return m.DecimalPlaces
	}
	return 0
}

type isMarket_TradingMode interface {
	isMarket_TradingMode()
}

type Market_Continuous struct {
	Continuous *ContinuousTrading `protobuf:"bytes,100,opt,name=continuous,proto3,oneof"`
}

type Market_Discrete struct {
	Discrete *DiscreteTrading `protobuf:"bytes,101,opt,name=discrete,proto3,oneof"`
}

func (*Market_Continuous) isMarket_TradingMode() {}

func (*Market_Discrete) isMarket_TradingMode() {}

func (m *Market) GetTradingMode() isMarket_TradingMode {
	if m != nil {
		return m.TradingMode
	}
	return nil
}

func (m *Market) GetContinuous() *ContinuousTrading {
	if x, ok := m.GetTradingMode().(*Market_Continuous); ok {
		return x.Continuous
	}
	return nil
}

func (m *Market) GetDiscrete() *DiscreteTrading {
	if x, ok := m.GetTradingMode().(*Market_Discrete); ok {
		return x.Discrete
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Market) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Market_Continuous)(nil),
		(*Market_Discrete)(nil),
	}
}

func init() {
	proto.RegisterType((*ContinuousTrading)(nil), "vega.ContinuousTrading")
	proto.RegisterType((*DiscreteTrading)(nil), "vega.DiscreteTrading")
	proto.RegisterType((*Future)(nil), "vega.Future")
	proto.RegisterType((*EthereumEvent)(nil), "vega.EthereumEvent")
	proto.RegisterType((*InstrumentMetadata)(nil), "vega.InstrumentMetadata")
	proto.RegisterType((*Instrument)(nil), "vega.Instrument")
	proto.RegisterType((*ModelParamsBS)(nil), "vega.ModelParamsBS")
	proto.RegisterType((*Forward)(nil), "vega.Forward")
	proto.RegisterType((*ExternalRiskModel)(nil), "vega.ExternalRiskModel")
	proto.RegisterMapType((map[string]string)(nil), "vega.ExternalRiskModel.ConfigEntry")
	proto.RegisterType((*ScalingFactors)(nil), "vega.ScalingFactors")
	proto.RegisterType((*MarginCalculator)(nil), "vega.MarginCalculator")
	proto.RegisterType((*TradableInstrument)(nil), "vega.TradableInstrument")
	proto.RegisterType((*Market)(nil), "vega.Market")
}

func init() { proto.RegisterFile("proto/markets.proto", fileDescriptor_ef38c4b9a7594dbd) }

var fileDescriptor_ef38c4b9a7594dbd = []byte{
	// 836 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xcd, 0x92, 0x1b, 0x35,
	0x10, 0xf6, 0xd8, 0xce, 0xec, 0xba, 0x1d, 0x2f, 0xbb, 0x4a, 0x08, 0x53, 0x29, 0x0a, 0xcc, 0x40,
	0x51, 0xa6, 0x00, 0x2f, 0xb5, 0xe1, 0x40, 0x08, 0x27, 0x6f, 0xbc, 0x38, 0x55, 0x2c, 0xb5, 0xa5,
	0xe5, 0xc4, 0x81, 0xaa, 0x5e, 0x8d, 0xd6, 0x51, 0x59, 0x33, 0x5a, 0x24, 0x8d, 0x21, 0xdc, 0xb8,
	0x70, 0xe1, 0x11, 0x78, 0x10, 0x1e, 0x89, 0xd7, 0xa0, 0xf4, 0xe3, 0xc9, 0xd8, 0x4e, 0x4e, 0xa3,
	0xaf, 0xfb, 0x6b, 0x4d, 0xf7, 0xa7, 0x56, 0x0b, 0x1e, 0xdc, 0x69, 0x65, 0xd5, 0x69, 0x89, 0x7a,
	0xc5, 0xad, 0x99, 0x7a, 0x44, 0xfa, 0x6b, 0xbe, 0xc4, 0xfc, 0x14, 0x4e, 0xce, 0x55, 0x65, 0x45,
	0x55, 0xab, 0xda, 0xfc, 0xa4, 0xb1, 0x10, 0xd5, 0x92, 0x3c, 0x86, 0x43, 0x2b, 0xd8, 0xea, 0x5a,
	0xfc, 0xc1, 0xb3, 0x64, 0x9c, 0x4c, 0xfa, 0xb4, 0xc1, 0xf9, 0x97, 0xf0, 0xce, 0x73, 0x61, 0x98,
	0xe6, 0x96, 0xb7, 0xe8, 0x45, 0xad, 0xd1, 0x0a, 0x55, 0x79, 0x7a, 0x8f, 0x36, 0x38, 0xff, 0x33,
	0x81, 0xf4, 0xa2, 0xb6, 0xb5, 0xe6, 0x8e, 0x56, 0xa2, 0xad, 0xb5, 0xb0, 0xaf, 0x3c, 0x6d, 0x40,
	0x1b, 0x4c, 0x1e, 0xc2, 0x3d, 0x34, 0x86, 0xdb, 0xac, 0xeb, 0x1d, 0x01, 0x90, 0x67, 0x30, 0xe2,
	0xf6, 0x25, 0xd7, 0xbc, 0x2e, 0xe7, 0x6b, 0x5e, 0xd9, 0xac, 0x18, 0x27, 0x93, 0xe1, 0xd9, 0x83,
	0xa9, 0x4b, 0x7d, 0x3a, 0x6f, 0xbb, 0x16, 0x1d, 0xba, 0xcd, 0x9d, 0x1d, 0x42, 0xaa, 0x34, 0x32,
	0xc9, 0xf3, 0x39, 0x8c, 0xb6, 0xb8, 0xe4, 0x03, 0x00, 0xa6, 0x2a, 0xab, 0x91, 0xd9, 0x17, 0xcf,
	0x63, 0x2e, 0x2d, 0x8b, 0xcb, 0x86, 0xfb, 0xff, 0xc5, 0x6c, 0x3c, 0xc8, 0x27, 0x40, 0x5e, 0x54,
	0xc6, 0xea, 0xba, 0xe4, 0x95, 0xbd, 0xe4, 0x16, 0x0b, 0xb4, 0x48, 0x08, 0xf4, 0x2d, 0x2e, 0x4d,
	0x96, 0x8c, 0x7b, 0x93, 0x01, 0xf5, 0xeb, 0xfc, 0xbf, 0x04, 0xe0, 0x35, 0x95, 0x1c, 0x41, 0x57,
	0x14, 0xf1, 0x37, 0x5d, 0x51, 0xb8, 0x10, 0xa6, 0x0a, 0x1e, 0x77, 0xf7, 0x6b, 0x67, 0xab, 0xb0,
	0xe4, 0x59, 0x2f, 0xd8, 0xdc, 0xda, 0x09, 0x76, 0x83, 0x86, 0xff, 0xe8, 0xec, 0xfd, 0x20, 0xd8,
	0x06, 0x93, 0xf7, 0x61, 0xf0, 0x6b, 0xad, 0x6c, 0x70, 0xde, 0xf3, 0xce, 0xd7, 0x06, 0xf2, 0x35,
	0x1c, 0x96, 0x31, 0xc1, 0x2c, 0xf5, 0x9a, 0x65, 0x41, 0xb3, 0xfd, 0x02, 0x68, 0xc3, 0x24, 0x9f,
	0x42, 0x7a, 0xeb, 0x8f, 0x2a, 0xea, 0x7c, 0x3f, 0xc4, 0x84, 0xe3, 0x5b, 0x74, 0x68, 0xf4, 0xce,
	0x06, 0x70, 0x70, 0xa7, 0x55, 0x51, 0x33, 0x9b, 0x9f, 0xc3, 0xe8, 0x52, 0x15, 0x5c, 0x5e, 0xa1,
	0xc6, 0xd2, 0xcc, 0xae, 0x5d, 0xad, 0x65, 0xed, 0x6b, 0x4d, 0x68, 0xb7, 0xac, 0xc9, 0x7d, 0x48,
	0xb4, 0x2f, 0x34, 0xa1, 0x89, 0x76, 0xc2, 0x1a, 0xb1, 0x2c, 0xd1, 0x97, 0x99, 0xd0, 0x00, 0xf2,
	0x5f, 0xe0, 0xe0, 0x42, 0xe9, 0xdf, 0x50, 0x17, 0x8e, 0x20, 0xb1, 0xbc, 0x29, 0xe2, 0x0e, 0x01,
	0x90, 0x63, 0xe8, 0x59, 0xac, 0xe3, 0x36, 0x6e, 0x49, 0x3e, 0x87, 0xf4, 0xce, 0xff, 0xd2, 0xef,
	0xd4, 0xb4, 0xc4, 0x56, 0x2e, 0x34, 0x52, 0xf2, 0x7f, 0x13, 0x38, 0x99, 0xff, 0x6e, 0xb9, 0xae,
	0x50, 0x52, 0x61, 0x56, 0x9e, 0xd5, 0x28, 0x9e, 0xb4, 0x14, 0x7f, 0x04, 0xa9, 0x51, 0x6c, 0xd5,
	0xf4, 0x61, 0x44, 0xe4, 0x19, 0xa4, 0x4c, 0x55, 0xb7, 0x62, 0x99, 0xf5, 0xc6, 0xbd, 0xc9, 0xf0,
	0xec, 0xe3, 0xd8, 0x81, 0xbb, 0x9b, 0x4e, 0xcf, 0x3d, 0x6b, 0x5e, 0x59, 0xfd, 0x8a, 0xc6, 0x90,
	0xc7, 0x4f, 0x61, 0xd8, 0x32, 0xbb, 0x62, 0x56, 0x7c, 0x73, 0x03, 0xdc, 0xd2, 0x15, 0xbd, 0x46,
	0x59, 0x6f, 0x1a, 0x22, 0x80, 0x6f, 0xbb, 0xdf, 0x24, 0xf9, 0x5f, 0x09, 0x1c, 0x5d, 0x33, 0x94,
	0xa2, 0x5a, 0x5e, 0x20, 0xb3, 0x4a, 0x1b, 0x32, 0x86, 0xa1, 0xe1, 0xa8, 0xd9, 0xcb, 0x1f, 0xf8,
	0x9a, 0xcb, 0xa8, 0x53, 0xdb, 0x44, 0x3e, 0x81, 0x91, 0xa8, 0x84, 0x15, 0x28, 0x2f, 0x51, 0x2f,
	0x45, 0x15, 0x75, 0xdb, 0x36, 0x92, 0x2f, 0xe0, 0x84, 0x29, 0x29, 0xd1, 0x72, 0x8d, 0x92, 0x72,
	0xc9, 0xd1, 0xf0, 0x78, 0x2c, 0xfb, 0x8e, 0xfc, 0x0a, 0x8e, 0x43, 0xdc, 0x39, 0x4a, 0x56, 0x4b,
	0xb4, 0x4a, 0x93, 0xef, 0xe0, 0xc8, 0x6c, 0xe5, 0xe6, 0x93, 0x19, 0x9e, 0x3d, 0x0c, 0xe2, 0x6c,
	0xe7, 0x4d, 0x77, 0xb8, 0xf9, 0xdf, 0x5d, 0x20, 0x6e, 0x80, 0xe0, 0x8d, 0xe4, 0xad, 0xbb, 0xf2,
	0x15, 0x80, 0x68, 0x50, 0xdc, 0xf0, 0x78, 0xb7, 0x77, 0x69, 0x8b, 0x43, 0x66, 0x70, 0x5c, 0xee,
	0xa4, 0xe6, 0x2b, 0x1e, 0x9e, 0x3d, 0x8a, 0x4d, 0xb1, 0xe3, 0xa5, 0x7b, 0x7c, 0xf2, 0x19, 0x1c,
	0xdc, 0x86, 0x0e, 0x8c, 0xad, 0x3f, 0x8a, 0xad, 0x1f, 0x8c, 0x8b, 0x0e, 0xdd, 0xf8, 0xc9, 0xf7,
	0x70, 0xc2, 0x77, 0x8f, 0x3d, 0xe3, 0x3e, 0xe8, 0xbd, 0xb7, 0x74, 0xc5, 0xa2, 0x43, 0xf7, 0x63,
	0x66, 0x43, 0x18, 0xe8, 0x0d, 0xc8, 0xff, 0xe9, 0x42, 0x7a, 0xe9, 0xc7, 0xf3, 0x9b, 0xa6, 0x85,
	0xef, 0xd3, 0x6e, 0xab, 0x4f, 0x17, 0x40, 0xec, 0x9e, 0x76, 0xf1, 0x2a, 0xc4, 0x9b, 0xbe, 0xaf,
	0x2d, 0x7d, 0x43, 0x8c, 0x6b, 0x96, 0x82, 0x33, 0x51, 0xa2, 0xbc, 0x92, 0xc8, 0xb8, 0xf1, 0x83,
	0xa6, 0x4f, 0xb7, 0x8d, 0xe4, 0x69, 0x18, 0x98, 0xe1, 0x95, 0x88, 0x12, 0xc5, 0x6a, 0xf7, 0x5e,
	0x8f, 0x45, 0x87, 0xb6, 0xc8, 0xe4, 0x09, 0x1c, 0x16, 0xf1, 0xbd, 0x88, 0x32, 0xbd, 0x1b, 0x02,
	0x77, 0x5e, 0x91, 0x45, 0x87, 0x36, 0xc4, 0xd9, 0x08, 0x86, 0x36, 0x98, 0x9d, 0x3c, 0xb3, 0x8f,
	0x7e, 0xfe, 0xd0, 0x0d, 0x49, 0x1f, 0xe7, 0x1f, 0x2f, 0xa6, 0xe4, 0x54, 0xa8, 0x53, 0x87, 0x4f,
	0xbd, 0xe1, 0x26, 0xf5, 0x9f, 0x27, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x90, 0xa0, 0xa4, 0xd1,
	0xeb, 0x06, 0x00, 0x00,
}
