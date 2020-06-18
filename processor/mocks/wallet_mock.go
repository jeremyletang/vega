// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: Wallet)

// Package mocks is a generated GoMock package.
package mocks

import (
	nodewallet "code.vegaprotocol.io/vega/nodewallet"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWallet is a mock of Wallet interface
type MockWallet struct {
	ctrl     *gomock.Controller
	recorder *MockWalletMockRecorder
}

// MockWalletMockRecorder is the mock recorder for MockWallet
type MockWalletMockRecorder struct {
	mock *MockWallet
}

// NewMockWallet creates a new mock instance
func NewMockWallet(ctrl *gomock.Controller) *MockWallet {
	mock := &MockWallet{ctrl: ctrl}
	mock.recorder = &MockWalletMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWallet) EXPECT() *MockWalletMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockWallet) Get(arg0 nodewallet.Blockchain) (nodewallet.Wallet, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(nodewallet.Wallet)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockWalletMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockWallet)(nil).Get), arg0)
}