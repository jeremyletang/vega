// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/collat (interfaces: AccountBuffer)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAccountBuffer is a mock of AccountBuffer interface
type MockAccountBuffer struct {
	ctrl     *gomock.Controller
	recorder *MockAccountBufferMockRecorder
}

// MockAccountBufferMockRecorder is the mock recorder for MockAccountBuffer
type MockAccountBufferMockRecorder struct {
	mock *MockAccountBuffer
}

// NewMockAccountBuffer creates a new mock instance
func NewMockAccountBuffer(ctrl *gomock.Controller) *MockAccountBuffer {
	mock := &MockAccountBuffer{ctrl: ctrl}
	mock.recorder = &MockAccountBufferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountBuffer) EXPECT() *MockAccountBufferMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockAccountBuffer) Add(arg0 proto.Account) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add
func (mr *MockAccountBufferMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockAccountBuffer)(nil).Add), arg0)
}
