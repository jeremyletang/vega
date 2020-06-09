// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/collateral (interfaces: Broker)

// Package mocks is a generated GoMock package.
package mocks

import (
	events "code.vegaprotocol.io/vega/events"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBroker is a mock of Broker interface
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockBroker) Send(arg0 events.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0)
}

// Send indicates an expected call of Send
func (mr *MockBrokerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBroker)(nil).Send), arg0)
}
