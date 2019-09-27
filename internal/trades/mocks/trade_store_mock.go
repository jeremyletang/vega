// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/trades (interfaces: TradeStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	storage "code.vegaprotocol.io/vega/internal/storage"
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTradeStore is a mock of TradeStore interface
type MockTradeStore struct {
	ctrl     *gomock.Controller
	recorder *MockTradeStoreMockRecorder
}

// MockTradeStoreMockRecorder is the mock recorder for MockTradeStore
type MockTradeStoreMockRecorder struct {
	mock *MockTradeStore
}

// NewMockTradeStore creates a new mock instance
func NewMockTradeStore(ctrl *gomock.Controller) *MockTradeStore {
	mock := &MockTradeStore{ctrl: ctrl}
	mock.recorder = &MockTradeStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTradeStore) EXPECT() *MockTradeStoreMockRecorder {
	return m.recorder
}

// GetByMarket mocks base method
func (m *MockTradeStore) GetByMarket(arg0 context.Context, arg1 string, arg2, arg3 uint64, arg4 bool) ([]*proto.Trade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMarket", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*proto.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarket indicates an expected call of GetByMarket
func (mr *MockTradeStoreMockRecorder) GetByMarket(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarket", reflect.TypeOf((*MockTradeStore)(nil).GetByMarket), arg0, arg1, arg2, arg3, arg4)
}

// GetByMarketAndID mocks base method
func (m *MockTradeStore) GetByMarketAndID(arg0 context.Context, arg1, arg2 string) (*proto.Trade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMarketAndID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*proto.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarketAndID indicates an expected call of GetByMarketAndID
func (mr *MockTradeStoreMockRecorder) GetByMarketAndID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarketAndID", reflect.TypeOf((*MockTradeStore)(nil).GetByMarketAndID), arg0, arg1, arg2)
}

// GetByOrderID mocks base method
func (m *MockTradeStore) GetByOrderID(arg0 context.Context, arg1 string, arg2, arg3 uint64, arg4 bool, arg5 *string) ([]*proto.Trade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByOrderID", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*proto.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByOrderID indicates an expected call of GetByOrderID
func (mr *MockTradeStoreMockRecorder) GetByOrderID(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByOrderID", reflect.TypeOf((*MockTradeStore)(nil).GetByOrderID), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetByParty mocks base method
func (m *MockTradeStore) GetByParty(arg0 context.Context, arg1 string, arg2, arg3 uint64, arg4 bool, arg5 *string) ([]*proto.Trade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParty", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*proto.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByParty indicates an expected call of GetByParty
func (mr *MockTradeStoreMockRecorder) GetByParty(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParty", reflect.TypeOf((*MockTradeStore)(nil).GetByParty), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetByPartyAndID mocks base method
func (m *MockTradeStore) GetByPartyAndID(arg0 context.Context, arg1, arg2 string) (*proto.Trade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPartyAndID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*proto.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPartyAndID indicates an expected call of GetByPartyAndID
func (mr *MockTradeStoreMockRecorder) GetByPartyAndID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPartyAndID", reflect.TypeOf((*MockTradeStore)(nil).GetByPartyAndID), arg0, arg1, arg2)
}

// GetMarkPrice mocks base method
func (m *MockTradeStore) GetMarkPrice(arg0 context.Context, arg1 string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarkPrice", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarkPrice indicates an expected call of GetMarkPrice
func (mr *MockTradeStoreMockRecorder) GetMarkPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarkPrice", reflect.TypeOf((*MockTradeStore)(nil).GetMarkPrice), arg0, arg1)
}

// GetTradesBySideBuckets mocks base method
func (m *MockTradeStore) GetTradesBySideBuckets(arg0 context.Context, arg1 string) map[string]*storage.MarketBucket {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTradesBySideBuckets", arg0, arg1)
	ret0, _ := ret[0].(map[string]*storage.MarketBucket)
	return ret0
}

// GetTradesBySideBuckets indicates an expected call of GetTradesBySideBuckets
func (mr *MockTradeStoreMockRecorder) GetTradesBySideBuckets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTradesBySideBuckets", reflect.TypeOf((*MockTradeStore)(nil).GetTradesBySideBuckets), arg0, arg1)
}

// Subscribe mocks base method
func (m *MockTradeStore) Subscribe(arg0 chan<- []proto.Trade) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockTradeStoreMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockTradeStore)(nil).Subscribe), arg0)
}

// Unsubscribe mocks base method
func (m *MockTradeStore) Unsubscribe(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockTradeStoreMockRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockTradeStore)(nil).Unsubscribe), arg0)
}
