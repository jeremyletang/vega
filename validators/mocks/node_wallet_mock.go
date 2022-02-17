// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/validators (interfaces: NodeWallets)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	validators "code.vegaprotocol.io/vega/validators"
	gomock "github.com/golang/mock/gomock"
)

// MockNodeWallets is a mock of NodeWallets interface.
type MockNodeWallets struct {
	ctrl     *gomock.Controller
	recorder *MockNodeWalletsMockRecorder
}

// MockNodeWalletsMockRecorder is the mock recorder for MockNodeWallets.
type MockNodeWalletsMockRecorder struct {
	mock *MockNodeWallets
}

// NewMockNodeWallets creates a new mock instance.
func NewMockNodeWallets(ctrl *gomock.Controller) *MockNodeWallets {
	mock := &MockNodeWallets{ctrl: ctrl}
	mock.recorder = &MockNodeWalletsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeWallets) EXPECT() *MockNodeWalletsMockRecorder {
	return m.recorder
}

// GetEthereum mocks base method.
func (m *MockNodeWallets) GetEthereum() validators.Signer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEthereum")
	ret0, _ := ret[0].(validators.Signer)
	return ret0
}

// GetEthereum indicates an expected call of GetEthereum.
func (mr *MockNodeWalletsMockRecorder) GetEthereum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEthereum", reflect.TypeOf((*MockNodeWallets)(nil).GetEthereum))
}

// GetEthereumAddress mocks base method.
func (m *MockNodeWallets) GetEthereumAddress() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEthereumAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEthereumAddress indicates an expected call of GetEthereumAddress.
func (mr *MockNodeWalletsMockRecorder) GetEthereumAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEthereumAddress", reflect.TypeOf((*MockNodeWallets)(nil).GetEthereumAddress))
}

// GetTendermintPubkey mocks base method.
func (m *MockNodeWallets) GetTendermintPubkey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTendermintPubkey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTendermintPubkey indicates an expected call of GetTendermintPubkey.
func (mr *MockNodeWalletsMockRecorder) GetTendermintPubkey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTendermintPubkey", reflect.TypeOf((*MockNodeWallets)(nil).GetTendermintPubkey))
}

// GetVega mocks base method.
func (m *MockNodeWallets) GetVega() validators.Wallet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVega")
	ret0, _ := ret[0].(validators.Wallet)
	return ret0
}

// GetVega indicates an expected call of GetVega.
func (mr *MockNodeWalletsMockRecorder) GetVega() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVega", reflect.TypeOf((*MockNodeWallets)(nil).GetVega))
}
