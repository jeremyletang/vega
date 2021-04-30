// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands/v1/commands.proto

package v1

import (
	proto1 "code.vegaprotocol.io/vega/proto"
	_ "code.vegaprotocol.io/vega/proto/oracles/v1"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
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

// An order submission is a request to submit or create a new order on Vega
type OrderSubmission struct {
	// Unique identifier for the order (set by the system after consensus)
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Market identifier for the order, required field
	MarketId string `protobuf:"bytes,2,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// Party identifier for the order, required field
	PartyId string `protobuf:"bytes,3,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// Price for the order, the price is an integer, for example `123456` is a correctly
	// formatted price of `1.23456` assuming market configured to 5 decimal places,
	// , required field for limit orders, however it is not required for market orders
	Price uint64 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	// Size for the order, for example, in a futures market the size equals the number of contracts, cannot be negative
	Size uint64 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	// Side for the order, e.g. SIDE_BUY or SIDE_SELL, required field - See [`Side`](#vega.Side)
	Side proto1.Side `protobuf:"varint,6,opt,name=side,proto3,enum=vega.Side" json:"side,omitempty"`
	// Time in force indicates how long an order will remain active before it is executed or expires, required field
	// - See [`Order.TimeInForce`](#vega.Order.TimeInForce)
	TimeInForce proto1.Order_TimeInForce `protobuf:"varint,7,opt,name=time_in_force,json=timeInForce,proto3,enum=vega.Order_TimeInForce" json:"time_in_force,omitempty"`
	// Timestamp for when the order will expire, in nanoseconds since the epoch,
	// required field only for [`Order.TimeInForce`](#vega.Order.TimeInForce)`.TIME_IN_FORCE_GTT`
	// - See [`VegaTimeResponse`](#api.VegaTimeResponse).`timestamp`
	ExpiresAt int64 `protobuf:"varint,8,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// Type for the order, required field - See [`Order.Type`](#vega.Order.Type)
	Type proto1.Order_Type `protobuf:"varint,9,opt,name=type,proto3,enum=vega.Order_Type" json:"type,omitempty"`
	// Reference given for the order, this is typically used to retrieve an order submitted through consensus, currently
	// set internally by the node to return a unique reference identifier for the order submission
	Reference string `protobuf:"bytes,10,opt,name=reference,proto3" json:"reference,omitempty"`
	// Used to specify the details for a pegged order
	// - See [`PeggedOrder`](#vega.PeggedOrder)
	PeggedOrder          *proto1.PeggedOrder `protobuf:"bytes,11,opt,name=pegged_order,json=peggedOrder,proto3" json:"pegged_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *OrderSubmission) Reset()         { *m = OrderSubmission{} }
func (m *OrderSubmission) String() string { return proto.CompactTextString(m) }
func (*OrderSubmission) ProtoMessage()    {}
func (*OrderSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd984d1ab638cd87, []int{0}
}

func (m *OrderSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderSubmission.Unmarshal(m, b)
}
func (m *OrderSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderSubmission.Marshal(b, m, deterministic)
}
func (m *OrderSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderSubmission.Merge(m, src)
}
func (m *OrderSubmission) XXX_Size() int {
	return xxx_messageInfo_OrderSubmission.Size(m)
}
func (m *OrderSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_OrderSubmission proto.InternalMessageInfo

func (m *OrderSubmission) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrderSubmission) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *OrderSubmission) GetPartyId() string {
	if m != nil {
		return m.PartyId
	}
	return ""
}

func (m *OrderSubmission) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderSubmission) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *OrderSubmission) GetSide() proto1.Side {
	if m != nil {
		return m.Side
	}
	return proto1.Side_SIDE_UNSPECIFIED
}

func (m *OrderSubmission) GetTimeInForce() proto1.Order_TimeInForce {
	if m != nil {
		return m.TimeInForce
	}
	return proto1.Order_TIME_IN_FORCE_UNSPECIFIED
}

func (m *OrderSubmission) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *OrderSubmission) GetType() proto1.Order_Type {
	if m != nil {
		return m.Type
	}
	return proto1.Order_TYPE_UNSPECIFIED
}

func (m *OrderSubmission) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *OrderSubmission) GetPeggedOrder() *proto1.PeggedOrder {
	if m != nil {
		return m.PeggedOrder
	}
	return nil
}

// An order cancellation is a request to cancel an existing order on Vega
type OrderCancellation struct {
	// Unique identifier for the order (set by the system after consensus), required field
	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// Market identifier for the order, required field
	MarketId string `protobuf:"bytes,2,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// Party identifier for the order, required field
	PartyId              string   `protobuf:"bytes,3,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderCancellation) Reset()         { *m = OrderCancellation{} }
func (m *OrderCancellation) String() string { return proto.CompactTextString(m) }
func (*OrderCancellation) ProtoMessage()    {}
func (*OrderCancellation) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd984d1ab638cd87, []int{1}
}

func (m *OrderCancellation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderCancellation.Unmarshal(m, b)
}
func (m *OrderCancellation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderCancellation.Marshal(b, m, deterministic)
}
func (m *OrderCancellation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderCancellation.Merge(m, src)
}
func (m *OrderCancellation) XXX_Size() int {
	return xxx_messageInfo_OrderCancellation.Size(m)
}
func (m *OrderCancellation) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderCancellation.DiscardUnknown(m)
}

var xxx_messageInfo_OrderCancellation proto.InternalMessageInfo

func (m *OrderCancellation) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderCancellation) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *OrderCancellation) GetPartyId() string {
	if m != nil {
		return m.PartyId
	}
	return ""
}

// An order amendment is a request to amend or update an existing order on Vega
type OrderAmendment struct {
	// Order identifier, this is required to find the order and will not be updated, required field
	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// Party identifier, this is required to find the order and will not be updated, required field
	PartyId string `protobuf:"bytes,2,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// Market identifier, this is required to find the order and will not be updated
	MarketId string `protobuf:"bytes,3,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// Amend the price for the order, if the Price value is set, otherwise price will remain unchanged - See [`Price`](#vega.Price)
	Price *proto1.Price `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	// Amend the size for the order by the delta specified:
	// - To reduce the size from the current value set a negative integer value
	// - To increase the size from the current value, set a positive integer value
	// - To leave the size unchanged set a value of zero
	SizeDelta int64 `protobuf:"varint,5,opt,name=size_delta,json=sizeDelta,proto3" json:"size_delta,omitempty"`
	// Amend the expiry time for the order, if the Timestamp value is set, otherwise expiry time will remain unchanged
	// - See [`VegaTimeResponse`](#api.VegaTimeResponse).`timestamp`
	ExpiresAt *proto1.Timestamp `protobuf:"bytes,6,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// Amend the time in force for the order, set to TIME_IN_FORCE_UNSPECIFIED to remain unchanged
	// - See [`TimeInForce`](#api.VegaTimeResponse).`timestamp`
	TimeInForce proto1.Order_TimeInForce `protobuf:"varint,7,opt,name=time_in_force,json=timeInForce,proto3,enum=vega.Order_TimeInForce" json:"time_in_force,omitempty"`
	// Amend the pegged order offset for the order
	PeggedOffset *wrappers.Int64Value `protobuf:"bytes,8,opt,name=pegged_offset,json=peggedOffset,proto3" json:"pegged_offset,omitempty"`
	// Amend the pegged order reference for the order
	// - See [`PeggedReference`](#vega.PeggedReference)
	PeggedReference      proto1.PeggedReference `protobuf:"varint,9,opt,name=pegged_reference,json=peggedReference,proto3,enum=vega.PeggedReference" json:"pegged_reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *OrderAmendment) Reset()         { *m = OrderAmendment{} }
func (m *OrderAmendment) String() string { return proto.CompactTextString(m) }
func (*OrderAmendment) ProtoMessage()    {}
func (*OrderAmendment) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd984d1ab638cd87, []int{2}
}

func (m *OrderAmendment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderAmendment.Unmarshal(m, b)
}
func (m *OrderAmendment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderAmendment.Marshal(b, m, deterministic)
}
func (m *OrderAmendment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderAmendment.Merge(m, src)
}
func (m *OrderAmendment) XXX_Size() int {
	return xxx_messageInfo_OrderAmendment.Size(m)
}
func (m *OrderAmendment) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderAmendment.DiscardUnknown(m)
}

var xxx_messageInfo_OrderAmendment proto.InternalMessageInfo

func (m *OrderAmendment) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderAmendment) GetPartyId() string {
	if m != nil {
		return m.PartyId
	}
	return ""
}

func (m *OrderAmendment) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *OrderAmendment) GetPrice() *proto1.Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *OrderAmendment) GetSizeDelta() int64 {
	if m != nil {
		return m.SizeDelta
	}
	return 0
}

func (m *OrderAmendment) GetExpiresAt() *proto1.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

func (m *OrderAmendment) GetTimeInForce() proto1.Order_TimeInForce {
	if m != nil {
		return m.TimeInForce
	}
	return proto1.Order_TIME_IN_FORCE_UNSPECIFIED
}

func (m *OrderAmendment) GetPeggedOffset() *wrappers.Int64Value {
	if m != nil {
		return m.PeggedOffset
	}
	return nil
}

func (m *OrderAmendment) GetPeggedReference() proto1.PeggedReference {
	if m != nil {
		return m.PeggedReference
	}
	return proto1.PeggedReference_PEGGED_REFERENCE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*OrderSubmission)(nil), "vega.commands.v1.OrderSubmission")
	proto.RegisterType((*OrderCancellation)(nil), "vega.commands.v1.OrderCancellation")
	proto.RegisterType((*OrderAmendment)(nil), "vega.commands.v1.OrderAmendment")
}

func init() { proto.RegisterFile("commands/v1/commands.proto", fileDescriptor_bd984d1ab638cd87) }

var fileDescriptor_bd984d1ab638cd87 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0xc5, 0x49, 0xc8, 0xcf, 0x98, 0x9f, 0x30, 0xfa, 0x7e, 0xdc, 0x00, 0x25, 0xa5, 0x8b, 0x46,
	0xaa, 0x62, 0x8b, 0x14, 0xb1, 0xe9, 0x06, 0x28, 0xaa, 0x94, 0x15, 0x95, 0xa9, 0xaa, 0xaa, 0x1b,
	0x6b, 0xe2, 0xb9, 0x71, 0x47, 0xd8, 0x1e, 0x6b, 0x3c, 0x09, 0xa5, 0x4f, 0xd0, 0x27, 0xe9, 0x13,
	0xf4, 0x7d, 0x90, 0x78, 0x92, 0x6a, 0x66, 0x62, 0x6c, 0x8b, 0x76, 0xc1, 0xee, 0xde, 0x73, 0xef,
	0x39, 0x77, 0x3c, 0xf7, 0x78, 0xd0, 0x20, 0xe4, 0x49, 0x42, 0x52, 0x9a, 0x7b, 0xcb, 0x23, 0xaf,
	0x88, 0xdd, 0x4c, 0x70, 0xc9, 0x71, 0x7f, 0x09, 0x11, 0x71, 0x1f, 0xc0, 0xe5, 0xd1, 0x60, 0x33,
	0x21, 0xe2, 0x1a, 0xe4, 0xaa, 0x61, 0xb0, 0x41, 0xf2, 0xbc, 0xcc, 0xfa, 0x11, 0x5f, 0x82, 0x48,
	0x49, 0x1a, 0xc2, 0x0a, 0x41, 0x5a, 0xc0, 0xc4, 0x7b, 0x5c, 0x90, 0x30, 0x06, 0x3d, 0xc7, 0x84,
	0x41, 0x9e, 0x41, 0xf8, 0xf7, 0x2a, 0x25, 0xb2, 0xe0, 0x3e, 0x8f, 0x38, 0x8f, 0x62, 0xf0, 0x74,
	0x36, 0x5b, 0xcc, 0xbd, 0x1b, 0x41, 0xb2, 0x0c, 0x44, 0x31, 0xf9, 0x24, 0x62, 0xf2, 0xeb, 0x62,
	0xa6, 0x8e, 0xea, 0x25, 0x37, 0x4c, 0x5e, 0xf3, 0x1b, 0x2f, 0xe2, 0x63, 0x5d, 0x1c, 0x2f, 0x49,
	0xcc, 0x28, 0x91, 0x5c, 0xe4, 0xde, 0x43, 0x68, 0x78, 0x87, 0xbf, 0x9a, 0x68, 0xfb, 0x52, 0x50,
	0x10, 0x57, 0x8b, 0x59, 0xc2, 0xf2, 0x9c, 0xf1, 0x14, 0xff, 0x87, 0x1a, 0x8c, 0x3a, 0xd6, 0xd0,
	0x1a, 0xf5, 0xce, 0xdb, 0xf7, 0x77, 0x07, 0x8d, 0xcf, 0x6b, 0x7e, 0x83, 0x51, 0xfc, 0x12, 0xf5,
	0xcc, 0xc7, 0x07, 0x8c, 0x3a, 0x8d, 0x4a, 0xd9, 0xf2, 0xbb, 0xa6, 0x30, 0xa5, 0xf8, 0x05, 0xea,
	0x66, 0x44, 0xc8, 0x5b, 0xd5, 0xd3, 0xac, 0xf5, 0x74, 0x34, 0x3e, 0xa5, 0xf8, 0x1f, 0xb4, 0x9e,
	0x09, 0x16, 0x82, 0xd3, 0x1a, 0x5a, 0xa3, 0x96, 0x6f, 0x12, 0x3c, 0x40, 0xad, 0x9c, 0x7d, 0x07,
	0x67, 0x5d, 0x81, 0x86, 0xd4, 0x5f, 0xf3, 0x35, 0x86, 0x5f, 0xa9, 0x1a, 0x05, 0xa7, 0x3d, 0xb4,
	0x46, 0x5b, 0x13, 0xe4, 0xea, 0x4b, 0xbd, 0x62, 0x14, 0xce, 0x3b, 0xf7, 0x77, 0x07, 0xcd, 0x1f,
	0x96, 0xe5, 0xeb, 0x06, 0x7c, 0x81, 0x36, 0x25, 0x4b, 0x20, 0x60, 0x69, 0x30, 0xe7, 0x22, 0x04,
	0xa7, 0xa3, 0x19, 0xff, 0x1b, 0x86, 0xfe, 0x50, 0xf7, 0x23, 0x4b, 0x60, 0x9a, 0xbe, 0x57, 0xe5,
	0x92, 0x6e, 0xcb, 0x12, 0xc5, 0xfb, 0x08, 0xc1, 0xb7, 0x8c, 0x09, 0xc8, 0x03, 0x22, 0x9d, 0xee,
	0xd0, 0x1a, 0x35, 0xfd, 0xde, 0x0a, 0x39, 0x93, 0x78, 0x8c, 0x5a, 0xf2, 0x36, 0x03, 0xa7, 0xa7,
	0xb5, 0xfb, 0x35, 0xed, 0xdb, 0xac, 0x7a, 0x26, 0xd5, 0x86, 0xf7, 0x50, 0x4f, 0xc0, 0x1c, 0x04,
	0xa4, 0x21, 0x38, 0x48, 0x5d, 0x89, 0x5f, 0x02, 0xf8, 0x18, 0x6d, 0x64, 0x10, 0x45, 0x40, 0x03,
	0xae, 0x14, 0x1c, 0x7b, 0x68, 0x8d, 0xec, 0xc9, 0x8e, 0x11, 0xfd, 0xa0, 0x2b, 0x5a, 0xda, 0xb7,
	0xb3, 0x32, 0x39, 0x4c, 0xd1, 0x8e, 0x0e, 0xde, 0x29, 0xab, 0xc5, 0x31, 0x91, 0x6a, 0x6f, 0xcf,
	0x50, 0x57, 0x6b, 0x04, 0xc5, 0xf6, 0xfc, 0x8e, 0xce, 0xa7, 0x14, 0xef, 0x3e, 0x5a, 0xdd, 0x93,
	0x56, 0x76, 0xf8, 0xb3, 0x89, 0xb6, 0xf4, 0xc0, 0xb3, 0x04, 0x52, 0x9a, 0x40, 0x2a, 0x15, 0xab,
	0x3e, 0xad, 0x64, 0x15, 0x53, 0xab, 0xc2, 0x8d, 0x3f, 0x7b, 0xa1, 0x76, 0xb0, 0xe6, 0xa3, 0x83,
	0x55, 0x8c, 0x62, 0x4f, 0xec, 0xd5, 0xa5, 0x28, 0xa8, 0x70, 0xcd, 0x3e, 0x42, 0xca, 0x21, 0x01,
	0x85, 0x58, 0x12, 0xed, 0x9d, 0xa6, 0xdf, 0x53, 0xc8, 0x85, 0x02, 0xb0, 0x5b, 0xdb, 0x64, 0x5b,
	0xcb, 0x6c, 0x1b, 0x19, 0x65, 0x83, 0x5c, 0x92, 0x24, 0xab, 0xae, 0xf6, 0xed, 0xd3, 0xfc, 0x53,
	0xb7, 0xcd, 0x29, 0xda, 0x2c, 0x56, 0x39, 0x9f, 0xe7, 0x60, 0x9c, 0x63, 0x4f, 0x76, 0x5d, 0xf3,
	0xef, 0xba, 0xc5, 0xbf, 0xeb, 0x4e, 0x53, 0x79, 0x72, 0xfc, 0x89, 0xc4, 0x0b, 0xf0, 0x57, 0xcb,
	0xbf, 0xd4, 0x04, 0x7c, 0x8a, 0xfa, 0x2b, 0x85, 0xd2, 0x31, 0xc6, 0x65, 0xff, 0x56, 0x0d, 0xe1,
	0x17, 0x45, 0x7f, 0x3b, 0xab, 0x03, 0xe7, 0xe3, 0x2f, 0xaf, 0x43, 0x4e, 0x41, 0x77, 0xeb, 0x79,
	0x21, 0x8f, 0x5d, 0xc6, 0x3d, 0x95, 0x9b, 0xc7, 0xc3, 0xab, 0x3c, 0x77, 0xb3, 0xb6, 0x86, 0xde,
	0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xda, 0xd7, 0x61, 0xad, 0x04, 0x05, 0x00, 0x00,
}
