// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/validators (interfaces: Commander)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	txn "code.vegaprotocol.io/vega/txn"
	gomock "github.com/golang/mock/gomock"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
)

// MockCommander is a mock of Commander interface.
type MockCommander struct {
	ctrl     *gomock.Controller
	recorder *MockCommanderMockRecorder
}

// MockCommanderMockRecorder is the mock recorder for MockCommander.
type MockCommanderMockRecorder struct {
	mock *MockCommander
}

// NewMockCommander creates a new mock instance.
func NewMockCommander(ctrl *gomock.Controller) *MockCommander {
	mock := &MockCommander{ctrl: ctrl}
	mock.recorder = &MockCommanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommander) EXPECT() *MockCommanderMockRecorder {
	return m.recorder
}

// Command mocks base method.
func (m *MockCommander) Command(arg0 context.Context, arg1 txn.Command, arg2 protoiface.MessageV1, arg3 func(error)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Command", arg0, arg1, arg2, arg3)
}

// Command indicates an expected call of Command.
func (mr *MockCommanderMockRecorder) Command(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockCommander)(nil).Command), arg0, arg1, arg2, arg3)
}

// CommandSync mocks base method.
func (m *MockCommander) CommandSync(arg0 context.Context, arg1 txn.Command, arg2 protoiface.MessageV1, arg3 func(error)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CommandSync", arg0, arg1, arg2, arg3)
}

// CommandSync indicates an expected call of CommandSync.
func (mr *MockCommanderMockRecorder) CommandSync(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandSync", reflect.TypeOf((*MockCommander)(nil).CommandSync), arg0, arg1, arg2, arg3)
}
