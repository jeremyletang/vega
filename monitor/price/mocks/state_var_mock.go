// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/monitor/price (interfaces: StateVarEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	statevar "code.vegaprotocol.io/vega/types/statevar"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStateVarEngine is a mock of StateVarEngine interface
type MockStateVarEngine struct {
	ctrl     *gomock.Controller
	recorder *MockStateVarEngineMockRecorder
}

// MockStateVarEngineMockRecorder is the mock recorder for MockStateVarEngine
type MockStateVarEngineMockRecorder struct {
	mock *MockStateVarEngine
}

// NewMockStateVarEngine creates a new mock instance
func NewMockStateVarEngine(ctrl *gomock.Controller) *MockStateVarEngine {
	mock := &MockStateVarEngine{ctrl: ctrl}
	mock.recorder = &MockStateVarEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateVarEngine) EXPECT() *MockStateVarEngineMockRecorder {
	return m.recorder
}

// AddStateVariable mocks base method
func (m *MockStateVarEngine) AddStateVariable(arg0, arg1 string, arg2 statevar.Converter, arg3 func(string, statevar.FinaliseCalculation), arg4 []statevar.StateVarEventType, arg5 func(context.Context, statevar.StateVariableResult) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStateVariable", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddStateVariable indicates an expected call of AddStateVariable
func (mr *MockStateVarEngineMockRecorder) AddStateVariable(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStateVariable", reflect.TypeOf((*MockStateVarEngine)(nil).AddStateVariable), arg0, arg1, arg2, arg3, arg4, arg5)
}
