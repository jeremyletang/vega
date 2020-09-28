// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/validators (interfaces: Wallet)

// Package mocks is a generated GoMock package.
package mocks

import (
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

// PubKeyOrAddress mocks base method
func (m *MockWallet) PubKeyOrAddress() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubKeyOrAddress")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// PubKeyOrAddress indicates an expected call of PubKeyOrAddress
func (mr *MockWalletMockRecorder) PubKeyOrAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubKeyOrAddress", reflect.TypeOf((*MockWallet)(nil).PubKeyOrAddress))
}
