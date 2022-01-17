// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/banking (interfaces: Collateral)

// Package mocks is a generated GoMock package.
package mocks

import (
	vega "code.vegaprotocol.io/protos/vega"
	types "code.vegaprotocol.io/vega/types"
	num "code.vegaprotocol.io/vega/types/num"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCollateral is a mock of Collateral interface
type MockCollateral struct {
	ctrl     *gomock.Controller
	recorder *MockCollateralMockRecorder
}

// MockCollateralMockRecorder is the mock recorder for MockCollateral
type MockCollateralMockRecorder struct {
	mock *MockCollateral
}

// NewMockCollateral creates a new mock instance
func NewMockCollateral(ctrl *gomock.Controller) *MockCollateral {
	mock := &MockCollateral{ctrl: ctrl}
	mock.recorder = &MockCollateralMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollateral) EXPECT() *MockCollateralMockRecorder {
	return m.recorder
}

// CreatePartyGeneralAccount mocks base method
func (m *MockCollateral) CreatePartyGeneralAccount(arg0 context.Context, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePartyGeneralAccount", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePartyGeneralAccount indicates an expected call of CreatePartyGeneralAccount
func (mr *MockCollateralMockRecorder) CreatePartyGeneralAccount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePartyGeneralAccount", reflect.TypeOf((*MockCollateral)(nil).CreatePartyGeneralAccount), arg0, arg1, arg2)
}

// Deposit mocks base method
func (m *MockCollateral) Deposit(arg0 context.Context, arg1, arg2 string, arg3 *num.Uint) (*types.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deposit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deposit indicates an expected call of Deposit
func (mr *MockCollateralMockRecorder) Deposit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deposit", reflect.TypeOf((*MockCollateral)(nil).Deposit), arg0, arg1, arg2, arg3)
}

// EnableAsset mocks base method
func (m *MockCollateral) EnableAsset(arg0 context.Context, arg1 types.Asset) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableAsset", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableAsset indicates an expected call of EnableAsset
func (mr *MockCollateralMockRecorder) EnableAsset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAsset", reflect.TypeOf((*MockCollateral)(nil).EnableAsset), arg0, arg1)
}

// GetPartyGeneralAccount mocks base method
func (m *MockCollateral) GetPartyGeneralAccount(arg0, arg1 string) (*types.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyGeneralAccount", arg0, arg1)
	ret0, _ := ret[0].(*types.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyGeneralAccount indicates an expected call of GetPartyGeneralAccount
func (mr *MockCollateralMockRecorder) GetPartyGeneralAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyGeneralAccount", reflect.TypeOf((*MockCollateral)(nil).GetPartyGeneralAccount), arg0, arg1)
}

// TransferFunds mocks base method
func (m *MockCollateral) TransferFunds(arg0 context.Context, arg1 []*types.Transfer, arg2 []vega.AccountType, arg3 []string, arg4 []*types.Transfer, arg5 []vega.AccountType) ([]*types.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferFunds", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*types.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferFunds indicates an expected call of TransferFunds
func (mr *MockCollateralMockRecorder) TransferFunds(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferFunds", reflect.TypeOf((*MockCollateral)(nil).TransferFunds), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Withdraw mocks base method
func (m *MockCollateral) Withdraw(arg0 context.Context, arg1, arg2 string, arg3 *num.Uint) (*types.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Withdraw indicates an expected call of Withdraw
func (mr *MockCollateralMockRecorder) Withdraw(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockCollateral)(nil).Withdraw), arg0, arg1, arg2, arg3)
}
