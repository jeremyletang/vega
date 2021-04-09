// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/risk (interfaces: Orderbook)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrderbook is a mock of Orderbook interface
type MockOrderbook struct {
	ctrl     *gomock.Controller
	recorder *MockOrderbookMockRecorder
}

// MockOrderbookMockRecorder is the mock recorder for MockOrderbook
type MockOrderbookMockRecorder struct {
	mock *MockOrderbook
}

// NewMockOrderbook creates a new mock instance
func NewMockOrderbook(ctrl *gomock.Controller) *MockOrderbook {
	mock := &MockOrderbook{ctrl: ctrl}
	mock.recorder = &MockOrderbookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderbook) EXPECT() *MockOrderbookMockRecorder {
	return m.recorder
}

// GetCloseoutPrice mocks base method
func (m *MockOrderbook) GetCloseoutPrice(arg0 uint64, arg1 proto.Side) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCloseoutPrice", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloseoutPrice indicates an expected call of GetCloseoutPrice
func (mr *MockOrderbookMockRecorder) GetCloseoutPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloseoutPrice", reflect.TypeOf((*MockOrderbook)(nil).GetCloseoutPrice), arg0, arg1)
}

// GetIndicativePrice mocks base method
func (m *MockOrderbook) GetIndicativePrice() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndicativePrice")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetIndicativePrice indicates an expected call of GetIndicativePrice
func (mr *MockOrderbookMockRecorder) GetIndicativePrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndicativePrice", reflect.TypeOf((*MockOrderbook)(nil).GetIndicativePrice))
}
