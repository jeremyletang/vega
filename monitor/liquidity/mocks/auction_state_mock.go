// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/monitor/liquidity (interfaces: AuctionState)

// Package mocks is a generated GoMock package.
package mocks

import (
	types "code.vegaprotocol.io/vega/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockAuctionState is a mock of AuctionState interface
type MockAuctionState struct {
	ctrl     *gomock.Controller
	recorder *MockAuctionStateMockRecorder
}

// MockAuctionStateMockRecorder is the mock recorder for MockAuctionState
type MockAuctionStateMockRecorder struct {
	mock *MockAuctionState
}

// NewMockAuctionState creates a new mock instance
func NewMockAuctionState(ctrl *gomock.Controller) *MockAuctionState {
	mock := &MockAuctionState{ctrl: ctrl}
	mock.recorder = &MockAuctionStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuctionState) EXPECT() *MockAuctionStateMockRecorder {
	return m.recorder
}

// ExpiresAt mocks base method
func (m *MockAuctionState) ExpiresAt() *time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpiresAt")
	ret0, _ := ret[0].(*time.Time)
	return ret0
}

// ExpiresAt indicates an expected call of ExpiresAt
func (mr *MockAuctionStateMockRecorder) ExpiresAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpiresAt", reflect.TypeOf((*MockAuctionState)(nil).ExpiresAt))
}

// ExtendAuctionLiquidity mocks base method
func (m *MockAuctionState) ExtendAuctionLiquidity(arg0 types.AuctionDuration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExtendAuctionLiquidity", arg0)
}

// ExtendAuctionLiquidity indicates an expected call of ExtendAuctionLiquidity
func (mr *MockAuctionStateMockRecorder) ExtendAuctionLiquidity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendAuctionLiquidity", reflect.TypeOf((*MockAuctionState)(nil).ExtendAuctionLiquidity), arg0)
}

// InAuction mocks base method
func (m *MockAuctionState) InAuction() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InAuction")
	ret0, _ := ret[0].(bool)
	return ret0
}

// InAuction indicates an expected call of InAuction
func (mr *MockAuctionStateMockRecorder) InAuction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InAuction", reflect.TypeOf((*MockAuctionState)(nil).InAuction))
}

// IsLiquidityAuction mocks base method
func (m *MockAuctionState) IsLiquidityAuction() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLiquidityAuction")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLiquidityAuction indicates an expected call of IsLiquidityAuction
func (mr *MockAuctionStateMockRecorder) IsLiquidityAuction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLiquidityAuction", reflect.TypeOf((*MockAuctionState)(nil).IsLiquidityAuction))
}

// SetReadyToLeave mocks base method
func (m *MockAuctionState) SetReadyToLeave() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReadyToLeave")
}

// SetReadyToLeave indicates an expected call of SetReadyToLeave
func (mr *MockAuctionStateMockRecorder) SetReadyToLeave() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadyToLeave", reflect.TypeOf((*MockAuctionState)(nil).SetReadyToLeave))
}

// StartLiquidityAuction mocks base method
func (m *MockAuctionState) StartLiquidityAuction(arg0 time.Time, arg1 *types.AuctionDuration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartLiquidityAuction", arg0, arg1)
}

// StartLiquidityAuction indicates an expected call of StartLiquidityAuction
func (mr *MockAuctionStateMockRecorder) StartLiquidityAuction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartLiquidityAuction", reflect.TypeOf((*MockAuctionState)(nil).StartLiquidityAuction), arg0, arg1)
}
