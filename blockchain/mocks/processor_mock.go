// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/blockchain (interfaces: Processor)

// Package mocks is a generated GoMock package.
package mocks

import (
	blockchain "code.vegaprotocol.io/vega/blockchain"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProcessor is a mock of Processor interface
type MockProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorMockRecorder
}

// MockProcessorMockRecorder is the mock recorder for MockProcessor
type MockProcessorMockRecorder struct {
	mock *MockProcessor
}

// NewMockProcessor creates a new mock instance
func NewMockProcessor(ctrl *gomock.Controller) *MockProcessor {
	mock := &MockProcessor{ctrl: ctrl}
	mock.recorder = &MockProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcessor) EXPECT() *MockProcessorMockRecorder {
	return m.recorder
}

// Process mocks base method
func (m *MockProcessor) Process(arg0 context.Context, arg1, arg2 []byte, arg3 blockchain.Command) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Process indicates an expected call of Process
func (mr *MockProcessorMockRecorder) Process(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockProcessor)(nil).Process), arg0, arg1, arg2, arg3)
}

// ValidateSigned mocks base method
func (m *MockProcessor) ValidateSigned(arg0, arg1 []byte, arg2 blockchain.Command) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateSigned", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateSigned indicates an expected call of ValidateSigned
func (mr *MockProcessorMockRecorder) ValidateSigned(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateSigned", reflect.TypeOf((*MockProcessor)(nil).ValidateSigned), arg0, arg1, arg2)
}
