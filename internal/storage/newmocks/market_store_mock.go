// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/storage (interfaces: MarketStore)

// Package newmocks is a generated GoMock package.
package newmocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMarketStore is a mock of MarketStore interface
type MockMarketStore struct {
	ctrl     *gomock.Controller
	recorder *MockMarketStoreMockRecorder
}

// MockMarketStoreMockRecorder is the mock recorder for MockMarketStore
type MockMarketStoreMockRecorder struct {
	mock *MockMarketStore
}

// NewMockMarketStore creates a new mock instance
func NewMockMarketStore(ctrl *gomock.Controller) *MockMarketStore {
	mock := &MockMarketStore{ctrl: ctrl}
	mock.recorder = &MockMarketStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMarketStore) EXPECT() *MockMarketStoreMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockMarketStore) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockMarketStoreMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMarketStore)(nil).Close))
}

// Commit mocks base method
func (m *MockMarketStore) Commit() error {
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockMarketStoreMockRecorder) Commit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockMarketStore)(nil).Commit))
}

// GetAll mocks base method
func (m *MockMarketStore) GetAll() ([]*proto.Market, error) {
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*proto.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockMarketStoreMockRecorder) GetAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockMarketStore)(nil).GetAll))
}

// GetByName mocks base method
func (m *MockMarketStore) GetByName(arg0 string) (*proto.Market, error) {
	ret := m.ctrl.Call(m, "GetByName", arg0)
	ret0, _ := ret[0].(*proto.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName
func (mr *MockMarketStoreMockRecorder) GetByName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockMarketStore)(nil).GetByName), arg0)
}

// Post mocks base method
func (m *MockMarketStore) Post(arg0 *proto.Market) error {
	ret := m.ctrl.Call(m, "Post", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post
func (mr *MockMarketStoreMockRecorder) Post(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockMarketStore)(nil).Post), arg0)
}
