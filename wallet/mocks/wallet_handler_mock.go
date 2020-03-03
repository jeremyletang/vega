// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/wallet (interfaces: WalletHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	wallet "code.vegaprotocol.io/vega/wallet"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWalletHandler is a mock of WalletHandler interface
type MockWalletHandler struct {
	ctrl     *gomock.Controller
	recorder *MockWalletHandlerMockRecorder
}

// MockWalletHandlerMockRecorder is the mock recorder for MockWalletHandler
type MockWalletHandlerMockRecorder struct {
	mock *MockWalletHandler
}

// NewMockWalletHandler creates a new mock instance
func NewMockWalletHandler(ctrl *gomock.Controller) *MockWalletHandler {
	mock := &MockWalletHandler{ctrl: ctrl}
	mock.recorder = &MockWalletHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWalletHandler) EXPECT() *MockWalletHandlerMockRecorder {
	return m.recorder
}

// CreateWallet mocks base method
func (m *MockWalletHandler) CreateWallet(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet
func (mr *MockWalletHandlerMockRecorder) CreateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockWalletHandler)(nil).CreateWallet), arg0, arg1)
}

// GenerateKeypair mocks base method
func (m *MockWalletHandler) GenerateKeypair(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateKeypair", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateKeypair indicates an expected call of GenerateKeypair
func (mr *MockWalletHandlerMockRecorder) GenerateKeypair(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateKeypair", reflect.TypeOf((*MockWalletHandler)(nil).GenerateKeypair), arg0, arg1)
}

// ListPublicKeys mocks base method
func (m *MockWalletHandler) ListPublicKeys(arg0 string) ([]wallet.Keypair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPublicKeys", arg0)
	ret0, _ := ret[0].([]wallet.Keypair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPublicKeys indicates an expected call of ListPublicKeys
func (mr *MockWalletHandlerMockRecorder) ListPublicKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPublicKeys", reflect.TypeOf((*MockWalletHandler)(nil).ListPublicKeys), arg0)
}

// LoginWallet mocks base method
func (m *MockWalletHandler) LoginWallet(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginWallet", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginWallet indicates an expected call of LoginWallet
func (mr *MockWalletHandlerMockRecorder) LoginWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginWallet", reflect.TypeOf((*MockWalletHandler)(nil).LoginWallet), arg0, arg1)
}

// RevokeToken mocks base method
func (m *MockWalletHandler) RevokeToken(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeToken indicates an expected call of RevokeToken
func (mr *MockWalletHandlerMockRecorder) RevokeToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeToken", reflect.TypeOf((*MockWalletHandler)(nil).RevokeToken), arg0)
}

// SignTx mocks base method
func (m *MockWalletHandler) SignTx(arg0, arg1, arg2 string) (wallet.SignedBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(wallet.SignedBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignTx indicates an expected call of SignTx
func (mr *MockWalletHandlerMockRecorder) SignTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTx", reflect.TypeOf((*MockWalletHandler)(nil).SignTx), arg0, arg1, arg2)
}

// TaintKey mocks base method
func (m *MockWalletHandler) TaintKey(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaintKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// TaintKey indicates an expected call of TaintKey
func (mr *MockWalletHandlerMockRecorder) TaintKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaintKey", reflect.TypeOf((*MockWalletHandler)(nil).TaintKey), arg0, arg1, arg2)
}

// UpdateMeta mocks base method
func (m *MockWalletHandler) UpdateMeta(arg0, arg1, arg2 string, arg3 []wallet.Meta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMeta", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMeta indicates an expected call of UpdateMeta
func (mr *MockWalletHandlerMockRecorder) UpdateMeta(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMeta", reflect.TypeOf((*MockWalletHandler)(nil).UpdateMeta), arg0, arg1, arg2, arg3)
}

// WalletPath mocks base method
func (m *MockWalletHandler) WalletPath(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletPath", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletPath indicates an expected call of WalletPath
func (mr *MockWalletHandlerMockRecorder) WalletPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletPath", reflect.TypeOf((*MockWalletHandler)(nil).WalletPath), arg0)
}
