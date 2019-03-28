// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/api/endpoints/gql (interfaces: TradeService)

// Package mocks is a generated GoMock package.
package mocks

import (
	filtering "code.vegaprotocol.io/vega/internal/filtering"
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTradeService is a mock of TradeService interface
type MockTradeService struct {
	ctrl     *gomock.Controller
	recorder *MockTradeServiceMockRecorder
}

// MockTradeServiceMockRecorder is the mock recorder for MockTradeService
type MockTradeServiceMockRecorder struct {
	mock *MockTradeService
}

// NewMockTradeService creates a new mock instance
func NewMockTradeService(ctrl *gomock.Controller) *MockTradeService {
	mock := &MockTradeService{ctrl: ctrl}
	mock.recorder = &MockTradeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTradeService) EXPECT() *MockTradeServiceMockRecorder {
	return m.recorder
}

// GetByMarket mocks base method
func (m *MockTradeService) GetByMarket(arg0 context.Context, arg1 string, arg2 *filtering.TradeQueryFilters) ([]*proto.Trade, error) {
	ret := m.ctrl.Call(m, "GetByMarket", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*proto.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarket indicates an expected call of GetByMarket
func (mr *MockTradeServiceMockRecorder) GetByMarket(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarket", reflect.TypeOf((*MockTradeService)(nil).GetByMarket), arg0, arg1, arg2)
}

// GetByOrderId mocks base method
func (m *MockTradeService) GetByOrderId(arg0 context.Context, arg1 string, arg2 *filtering.TradeQueryFilters) ([]*proto.Trade, error) {
	ret := m.ctrl.Call(m, "GetByOrderId", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*proto.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByOrderId indicates an expected call of GetByOrderId
func (mr *MockTradeServiceMockRecorder) GetByOrderId(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByOrderId", reflect.TypeOf((*MockTradeService)(nil).GetByOrderId), arg0, arg1, arg2)
}

// GetByParty mocks base method
func (m *MockTradeService) GetByParty(arg0 context.Context, arg1 string, arg2 *filtering.TradeQueryFilters) ([]*proto.Trade, error) {
	ret := m.ctrl.Call(m, "GetByParty", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*proto.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByParty indicates an expected call of GetByParty
func (mr *MockTradeServiceMockRecorder) GetByParty(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParty", reflect.TypeOf((*MockTradeService)(nil).GetByParty), arg0, arg1, arg2)
}

// GetPositionsByParty mocks base method
func (m *MockTradeService) GetPositionsByParty(arg0 context.Context, arg1 string) ([]*proto.MarketPosition, error) {
	ret := m.ctrl.Call(m, "GetPositionsByParty", arg0, arg1)
	ret0, _ := ret[0].([]*proto.MarketPosition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPositionsByParty indicates an expected call of GetPositionsByParty
func (mr *MockTradeServiceMockRecorder) GetPositionsByParty(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPositionsByParty", reflect.TypeOf((*MockTradeService)(nil).GetPositionsByParty), arg0, arg1)
}

// ObservePositions mocks base method
func (m *MockTradeService) ObservePositions(arg0 context.Context, arg1 int, arg2 string) (<-chan *proto.MarketPosition, uint64) {
	ret := m.ctrl.Call(m, "ObservePositions", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan *proto.MarketPosition)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// ObservePositions indicates an expected call of ObservePositions
func (mr *MockTradeServiceMockRecorder) ObservePositions(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObservePositions", reflect.TypeOf((*MockTradeService)(nil).ObservePositions), arg0, arg1, arg2)
}

// ObserveTrades mocks base method
func (m *MockTradeService) ObserveTrades(arg0 context.Context, arg1 int, arg2, arg3 *string) (<-chan []proto.Trade, uint64) {
	ret := m.ctrl.Call(m, "ObserveTrades", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan []proto.Trade)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// ObserveTrades indicates an expected call of ObserveTrades
func (mr *MockTradeServiceMockRecorder) ObserveTrades(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveTrades", reflect.TypeOf((*MockTradeService)(nil).ObserveTrades), arg0, arg1, arg2, arg3)
}
