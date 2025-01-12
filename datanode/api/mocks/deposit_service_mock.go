// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/datanode/api (interfaces: DepositService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	vega "code.vegaprotocol.io/vega/protos/vega"
	gomock "github.com/golang/mock/gomock"
)

// MockDepositService is a mock of DepositService interface.
type MockDepositService struct {
	ctrl     *gomock.Controller
	recorder *MockDepositServiceMockRecorder
}

// MockDepositServiceMockRecorder is the mock recorder for MockDepositService.
type MockDepositServiceMockRecorder struct {
	mock *MockDepositService
}

// NewMockDepositService creates a new mock instance.
func NewMockDepositService(ctrl *gomock.Controller) *MockDepositService {
	mock := &MockDepositService{ctrl: ctrl}
	mock.recorder = &MockDepositServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDepositService) EXPECT() *MockDepositServiceMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockDepositService) GetByID(arg0 string) (vega.Deposit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(vega.Deposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockDepositServiceMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockDepositService)(nil).GetByID), arg0)
}

// GetByParty mocks base method.
func (m *MockDepositService) GetByParty(arg0 string, arg1 bool) []vega.Deposit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParty", arg0, arg1)
	ret0, _ := ret[0].([]vega.Deposit)
	return ret0
}

// GetByParty indicates an expected call of GetByParty.
func (mr *MockDepositServiceMockRecorder) GetByParty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParty", reflect.TypeOf((*MockDepositService)(nil).GetByParty), arg0, arg1)
}
