// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: vega/api/v1/core.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoreServiceClient is the client API for CoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreServiceClient interface {
	// Submit transaction
	//
	// Submit a signed transaction
	SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error)
	// Chain event
	//
	// Propagate a chain event
	PropagateChainEvent(ctx context.Context, in *PropagateChainEventRequest, opts ...grpc.CallOption) (*PropagateChainEventResponse, error)
	// Statistics
	//
	// Get statistics on Vega
	Statistics(ctx context.Context, in *StatisticsRequest, opts ...grpc.CallOption) (*StatisticsResponse, error)
	// Blockchain height
	//
	// Get the height of the last tendermint block
	LastBlockHeight(ctx context.Context, in *LastBlockHeightRequest, opts ...grpc.CallOption) (*LastBlockHeightResponse, error)
	// Vega time
	//
	// Get current Vega time
	GetVegaTime(ctx context.Context, in *GetVegaTimeRequest, opts ...grpc.CallOption) (*GetVegaTimeResponse, error)
	// Events subscription
	//
	// Subscribe to a stream of events from the core
	ObserveEventBus(ctx context.Context, opts ...grpc.CallOption) (CoreService_ObserveEventBusClient, error)
	// Submit raw transaction
	//
	// Submit a version agnostic signed transaction
	SubmitRawTransaction(ctx context.Context, in *SubmitRawTransactionRequest, opts ...grpc.CallOption) (*SubmitRawTransactionResponse, error)
	// Check transaction
	//
	// Check a signed transaction
	CheckTransaction(ctx context.Context, in *CheckTransactionRequest, opts ...grpc.CallOption) (*CheckTransactionResponse, error)
	// Check raw transaction
	//
	// Check a raw signed transaction
	CheckRawTransaction(ctx context.Context, in *CheckRawTransactionRequest, opts ...grpc.CallOption) (*CheckRawTransactionResponse, error)
	// Get Spam statistics
	//
	// Retrieve the spam statistics for a given party
	GetSpamStatistics(ctx context.Context, in *GetSpamStatisticsRequest, opts ...grpc.CallOption) (*GetSpamStatisticsResponse, error)
}

type coreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreServiceClient(cc grpc.ClientConnInterface) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error) {
	out := new(SubmitTransactionResponse)
	err := c.cc.Invoke(ctx, "/vega.api.v1.CoreService/SubmitTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) PropagateChainEvent(ctx context.Context, in *PropagateChainEventRequest, opts ...grpc.CallOption) (*PropagateChainEventResponse, error) {
	out := new(PropagateChainEventResponse)
	err := c.cc.Invoke(ctx, "/vega.api.v1.CoreService/PropagateChainEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) Statistics(ctx context.Context, in *StatisticsRequest, opts ...grpc.CallOption) (*StatisticsResponse, error) {
	out := new(StatisticsResponse)
	err := c.cc.Invoke(ctx, "/vega.api.v1.CoreService/Statistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) LastBlockHeight(ctx context.Context, in *LastBlockHeightRequest, opts ...grpc.CallOption) (*LastBlockHeightResponse, error) {
	out := new(LastBlockHeightResponse)
	err := c.cc.Invoke(ctx, "/vega.api.v1.CoreService/LastBlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetVegaTime(ctx context.Context, in *GetVegaTimeRequest, opts ...grpc.CallOption) (*GetVegaTimeResponse, error) {
	out := new(GetVegaTimeResponse)
	err := c.cc.Invoke(ctx, "/vega.api.v1.CoreService/GetVegaTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) ObserveEventBus(ctx context.Context, opts ...grpc.CallOption) (CoreService_ObserveEventBusClient, error) {
	stream, err := c.cc.NewStream(ctx, &CoreService_ServiceDesc.Streams[0], "/vega.api.v1.CoreService/ObserveEventBus", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreServiceObserveEventBusClient{stream}
	return x, nil
}

type CoreService_ObserveEventBusClient interface {
	Send(*ObserveEventBusRequest) error
	Recv() (*ObserveEventBusResponse, error)
	grpc.ClientStream
}

type coreServiceObserveEventBusClient struct {
	grpc.ClientStream
}

func (x *coreServiceObserveEventBusClient) Send(m *ObserveEventBusRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *coreServiceObserveEventBusClient) Recv() (*ObserveEventBusResponse, error) {
	m := new(ObserveEventBusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreServiceClient) SubmitRawTransaction(ctx context.Context, in *SubmitRawTransactionRequest, opts ...grpc.CallOption) (*SubmitRawTransactionResponse, error) {
	out := new(SubmitRawTransactionResponse)
	err := c.cc.Invoke(ctx, "/vega.api.v1.CoreService/SubmitRawTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) CheckTransaction(ctx context.Context, in *CheckTransactionRequest, opts ...grpc.CallOption) (*CheckTransactionResponse, error) {
	out := new(CheckTransactionResponse)
	err := c.cc.Invoke(ctx, "/vega.api.v1.CoreService/CheckTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) CheckRawTransaction(ctx context.Context, in *CheckRawTransactionRequest, opts ...grpc.CallOption) (*CheckRawTransactionResponse, error) {
	out := new(CheckRawTransactionResponse)
	err := c.cc.Invoke(ctx, "/vega.api.v1.CoreService/CheckRawTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetSpamStatistics(ctx context.Context, in *GetSpamStatisticsRequest, opts ...grpc.CallOption) (*GetSpamStatisticsResponse, error) {
	out := new(GetSpamStatisticsResponse)
	err := c.cc.Invoke(ctx, "/vega.api.v1.CoreService/GetSpamStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServiceServer is the server API for CoreService service.
// All implementations must embed UnimplementedCoreServiceServer
// for forward compatibility
type CoreServiceServer interface {
	// Submit transaction
	//
	// Submit a signed transaction
	SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error)
	// Chain event
	//
	// Propagate a chain event
	PropagateChainEvent(context.Context, *PropagateChainEventRequest) (*PropagateChainEventResponse, error)
	// Statistics
	//
	// Get statistics on Vega
	Statistics(context.Context, *StatisticsRequest) (*StatisticsResponse, error)
	// Blockchain height
	//
	// Get the height of the last tendermint block
	LastBlockHeight(context.Context, *LastBlockHeightRequest) (*LastBlockHeightResponse, error)
	// Vega time
	//
	// Get current Vega time
	GetVegaTime(context.Context, *GetVegaTimeRequest) (*GetVegaTimeResponse, error)
	// Events subscription
	//
	// Subscribe to a stream of events from the core
	ObserveEventBus(CoreService_ObserveEventBusServer) error
	// Submit raw transaction
	//
	// Submit a version agnostic signed transaction
	SubmitRawTransaction(context.Context, *SubmitRawTransactionRequest) (*SubmitRawTransactionResponse, error)
	// Check transaction
	//
	// Check a signed transaction
	CheckTransaction(context.Context, *CheckTransactionRequest) (*CheckTransactionResponse, error)
	// Check raw transaction
	//
	// Check a raw signed transaction
	CheckRawTransaction(context.Context, *CheckRawTransactionRequest) (*CheckRawTransactionResponse, error)
	// Get Spam statistics
	//
	// Retrieve the spam statistics for a given party
	GetSpamStatistics(context.Context, *GetSpamStatisticsRequest) (*GetSpamStatisticsResponse, error)
	mustEmbedUnimplementedCoreServiceServer()
}

// UnimplementedCoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoreServiceServer struct {
}

func (UnimplementedCoreServiceServer) SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransaction not implemented")
}
func (UnimplementedCoreServiceServer) PropagateChainEvent(context.Context, *PropagateChainEventRequest) (*PropagateChainEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PropagateChainEvent not implemented")
}
func (UnimplementedCoreServiceServer) Statistics(context.Context, *StatisticsRequest) (*StatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Statistics not implemented")
}
func (UnimplementedCoreServiceServer) LastBlockHeight(context.Context, *LastBlockHeightRequest) (*LastBlockHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastBlockHeight not implemented")
}
func (UnimplementedCoreServiceServer) GetVegaTime(context.Context, *GetVegaTimeRequest) (*GetVegaTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVegaTime not implemented")
}
func (UnimplementedCoreServiceServer) ObserveEventBus(CoreService_ObserveEventBusServer) error {
	return status.Errorf(codes.Unimplemented, "method ObserveEventBus not implemented")
}
func (UnimplementedCoreServiceServer) SubmitRawTransaction(context.Context, *SubmitRawTransactionRequest) (*SubmitRawTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitRawTransaction not implemented")
}
func (UnimplementedCoreServiceServer) CheckTransaction(context.Context, *CheckTransactionRequest) (*CheckTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTransaction not implemented")
}
func (UnimplementedCoreServiceServer) CheckRawTransaction(context.Context, *CheckRawTransactionRequest) (*CheckRawTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRawTransaction not implemented")
}
func (UnimplementedCoreServiceServer) GetSpamStatistics(context.Context, *GetSpamStatisticsRequest) (*GetSpamStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpamStatistics not implemented")
}
func (UnimplementedCoreServiceServer) mustEmbedUnimplementedCoreServiceServer() {}

// UnsafeCoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreServiceServer will
// result in compilation errors.
type UnsafeCoreServiceServer interface {
	mustEmbedUnimplementedCoreServiceServer()
}

func RegisterCoreServiceServer(s grpc.ServiceRegistrar, srv CoreServiceServer) {
	s.RegisterService(&CoreService_ServiceDesc, srv)
}

func _CoreService_SubmitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).SubmitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vega.api.v1.CoreService/SubmitTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).SubmitTransaction(ctx, req.(*SubmitTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_PropagateChainEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropagateChainEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).PropagateChainEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vega.api.v1.CoreService/PropagateChainEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).PropagateChainEvent(ctx, req.(*PropagateChainEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_Statistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).Statistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vega.api.v1.CoreService/Statistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).Statistics(ctx, req.(*StatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_LastBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastBlockHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).LastBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vega.api.v1.CoreService/LastBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).LastBlockHeight(ctx, req.(*LastBlockHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetVegaTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVegaTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetVegaTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vega.api.v1.CoreService/GetVegaTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetVegaTime(ctx, req.(*GetVegaTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_ObserveEventBus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoreServiceServer).ObserveEventBus(&coreServiceObserveEventBusServer{stream})
}

type CoreService_ObserveEventBusServer interface {
	Send(*ObserveEventBusResponse) error
	Recv() (*ObserveEventBusRequest, error)
	grpc.ServerStream
}

type coreServiceObserveEventBusServer struct {
	grpc.ServerStream
}

func (x *coreServiceObserveEventBusServer) Send(m *ObserveEventBusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *coreServiceObserveEventBusServer) Recv() (*ObserveEventBusRequest, error) {
	m := new(ObserveEventBusRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CoreService_SubmitRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRawTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).SubmitRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vega.api.v1.CoreService/SubmitRawTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).SubmitRawTransaction(ctx, req.(*SubmitRawTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_CheckTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CheckTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vega.api.v1.CoreService/CheckTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CheckTransaction(ctx, req.(*CheckTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_CheckRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRawTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CheckRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vega.api.v1.CoreService/CheckRawTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CheckRawTransaction(ctx, req.(*CheckRawTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetSpamStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpamStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetSpamStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vega.api.v1.CoreService/GetSpamStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetSpamStatistics(ctx, req.(*GetSpamStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoreService_ServiceDesc is the grpc.ServiceDesc for CoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vega.api.v1.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTransaction",
			Handler:    _CoreService_SubmitTransaction_Handler,
		},
		{
			MethodName: "PropagateChainEvent",
			Handler:    _CoreService_PropagateChainEvent_Handler,
		},
		{
			MethodName: "Statistics",
			Handler:    _CoreService_Statistics_Handler,
		},
		{
			MethodName: "LastBlockHeight",
			Handler:    _CoreService_LastBlockHeight_Handler,
		},
		{
			MethodName: "GetVegaTime",
			Handler:    _CoreService_GetVegaTime_Handler,
		},
		{
			MethodName: "SubmitRawTransaction",
			Handler:    _CoreService_SubmitRawTransaction_Handler,
		},
		{
			MethodName: "CheckTransaction",
			Handler:    _CoreService_CheckTransaction_Handler,
		},
		{
			MethodName: "CheckRawTransaction",
			Handler:    _CoreService_CheckRawTransaction_Handler,
		},
		{
			MethodName: "GetSpamStatistics",
			Handler:    _CoreService_GetSpamStatistics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ObserveEventBus",
			Handler:       _CoreService_ObserveEventBus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "vega/api/v1/core.proto",
}
