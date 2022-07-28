// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/datanode/api (interfaces: AccountsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	vega "code.vegaprotocol.io/protos/vega"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountsService is a mock of AccountsService interface.
type MockAccountsService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsServiceMockRecorder
}

// MockAccountsServiceMockRecorder is the mock recorder for MockAccountsService.
type MockAccountsServiceMockRecorder struct {
	mock *MockAccountsService
}

// NewMockAccountsService creates a new mock instance.
func NewMockAccountsService(ctrl *gomock.Controller) *MockAccountsService {
	mock := &MockAccountsService{ctrl: ctrl}
	mock.recorder = &MockAccountsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountsService) EXPECT() *MockAccountsServiceMockRecorder {
	return m.recorder
}

// GetAccountSubscribersCount mocks base method.
func (m *MockAccountsService) GetAccountSubscribersCount() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountSubscribersCount")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetAccountSubscribersCount indicates an expected call of GetAccountSubscribersCount.
func (mr *MockAccountsServiceMockRecorder) GetAccountSubscribersCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountSubscribersCount", reflect.TypeOf((*MockAccountsService)(nil).GetAccountSubscribersCount))
}

// GetFeeInfrastructureAccounts mocks base method.
func (m *MockAccountsService) GetFeeInfrastructureAccounts(arg0 string) ([]*vega.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeeInfrastructureAccounts", arg0)
	ret0, _ := ret[0].([]*vega.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeeInfrastructureAccounts indicates an expected call of GetFeeInfrastructureAccounts.
func (mr *MockAccountsServiceMockRecorder) GetFeeInfrastructureAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeeInfrastructureAccounts", reflect.TypeOf((*MockAccountsService)(nil).GetFeeInfrastructureAccounts), arg0)
}

// GetGlobalRewardPoolAccounts mocks base method.
func (m *MockAccountsService) GetGlobalRewardPoolAccounts(arg0 string) ([]*vega.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGlobalRewardPoolAccounts", arg0)
	ret0, _ := ret[0].([]*vega.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGlobalRewardPoolAccounts indicates an expected call of GetGlobalRewardPoolAccounts.
func (mr *MockAccountsServiceMockRecorder) GetGlobalRewardPoolAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGlobalRewardPoolAccounts", reflect.TypeOf((*MockAccountsService)(nil).GetGlobalRewardPoolAccounts), arg0)
}

// GetMarketAccounts mocks base method.
func (m *MockAccountsService) GetMarketAccounts(arg0, arg1 string) ([]*vega.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketAccounts", arg0, arg1)
	ret0, _ := ret[0].([]*vega.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketAccounts indicates an expected call of GetMarketAccounts.
func (mr *MockAccountsServiceMockRecorder) GetMarketAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketAccounts", reflect.TypeOf((*MockAccountsService)(nil).GetMarketAccounts), arg0, arg1)
}

// GetPartyAccounts mocks base method.
func (m *MockAccountsService) GetPartyAccounts(arg0, arg1, arg2 string, arg3 vega.AccountType) ([]*vega.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyAccounts", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*vega.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyAccounts indicates an expected call of GetPartyAccounts.
func (mr *MockAccountsServiceMockRecorder) GetPartyAccounts(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyAccounts", reflect.TypeOf((*MockAccountsService)(nil).GetPartyAccounts), arg0, arg1, arg2, arg3)
}

// ObserveAccounts mocks base method.
func (m *MockAccountsService) ObserveAccounts(arg0 context.Context, arg1 int, arg2, arg3, arg4 string, arg5 vega.AccountType) (<-chan []*vega.Account, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObserveAccounts", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(<-chan []*vega.Account)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// ObserveAccounts indicates an expected call of ObserveAccounts.
func (mr *MockAccountsServiceMockRecorder) ObserveAccounts(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveAccounts", reflect.TypeOf((*MockAccountsService)(nil).ObserveAccounts), arg0, arg1, arg2, arg3, arg4, arg5)
}
