// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/validators/erc20multisig (interfaces: Witness)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	validators "code.vegaprotocol.io/vega/validators"
	gomock "github.com/golang/mock/gomock"
)

// MockWitness is a mock of Witness interface.
type MockWitness struct {
	ctrl     *gomock.Controller
	recorder *MockWitnessMockRecorder
}

// MockWitnessMockRecorder is the mock recorder for MockWitness.
type MockWitnessMockRecorder struct {
	mock *MockWitness
}

// NewMockWitness creates a new mock instance.
func NewMockWitness(ctrl *gomock.Controller) *MockWitness {
	mock := &MockWitness{ctrl: ctrl}
	mock.recorder = &MockWitnessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWitness) EXPECT() *MockWitnessMockRecorder {
	return m.recorder
}

// RestoreResource mocks base method.
func (m *MockWitness) RestoreResource(arg0 validators.Resource, arg1 func(interface{}, bool)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreResource", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreResource indicates an expected call of RestoreResource.
func (mr *MockWitnessMockRecorder) RestoreResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreResource", reflect.TypeOf((*MockWitness)(nil).RestoreResource), arg0, arg1)
}

// StartCheck mocks base method.
func (m *MockWitness) StartCheck(arg0 validators.Resource, arg1 func(interface{}, bool), arg2 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCheck", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartCheck indicates an expected call of StartCheck.
func (mr *MockWitnessMockRecorder) StartCheck(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCheck", reflect.TypeOf((*MockWitness)(nil).StartCheck), arg0, arg1, arg2)
}
