// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/datanode/service (interfaces: OrderStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "code.vegaprotocol.io/vega/datanode/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockOrderStore is a mock of OrderStore interface.
type MockOrderStore struct {
	ctrl     *gomock.Controller
	recorder *MockOrderStoreMockRecorder
}

// MockOrderStoreMockRecorder is the mock recorder for MockOrderStore.
type MockOrderStoreMockRecorder struct {
	mock *MockOrderStore
}

// NewMockOrderStore creates a new mock instance.
func NewMockOrderStore(ctrl *gomock.Controller) *MockOrderStore {
	mock := &MockOrderStore{ctrl: ctrl}
	mock.recorder = &MockOrderStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderStore) EXPECT() *MockOrderStoreMockRecorder {
	return m.recorder
}

// GetLiveOrders mocks base method.
func (m *MockOrderStore) GetLiveOrders(arg0 context.Context) ([]entities.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLiveOrders", arg0)
	ret0, _ := ret[0].([]entities.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLiveOrders indicates an expected call of GetLiveOrders.
func (mr *MockOrderStoreMockRecorder) GetLiveOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLiveOrders", reflect.TypeOf((*MockOrderStore)(nil).GetLiveOrders), arg0)
}
