// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/nodewallets (interfaces: Chain)

// Package mocks is a generated GoMock package.
package mocks

import (
	v1 "code.vegaprotocol.io/protos/vega/api/v1"
	v10 "code.vegaprotocol.io/protos/vega/commands/v1"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockChain is a mock of Chain interface
type MockChain struct {
	ctrl     *gomock.Controller
	recorder *MockChainMockRecorder
}

// MockChainMockRecorder is the mock recorder for MockChain
type MockChainMockRecorder struct {
	mock *MockChain
}

// NewMockChain creates a new mock instance
func NewMockChain(ctrl *gomock.Controller) *MockChain {
	mock := &MockChain{ctrl: ctrl}
	mock.recorder = &MockChainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChain) EXPECT() *MockChainMockRecorder {
	return m.recorder
}

// SubmitTransactionV2 mocks base method
func (m *MockChain) SubmitTransactionV2(arg0 context.Context, arg1 *v10.Transaction, arg2 v1.SubmitTransactionRequest_Type) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTransactionV2", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitTransactionV2 indicates an expected call of SubmitTransactionV2
func (mr *MockChainMockRecorder) SubmitTransactionV2(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransactionV2", reflect.TypeOf((*MockChain)(nil).SubmitTransactionV2), arg0, arg1, arg2)
}