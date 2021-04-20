// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/cmd/vegabenchmark/mocks (interfaces: EvtForwarder)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEvtForwarder is a mock of EvtForwarder interface
type MockEvtForwarder struct {
	ctrl     *gomock.Controller
	recorder *MockEvtForwarderMockRecorder
}

// MockEvtForwarderMockRecorder is the mock recorder for MockEvtForwarder
type MockEvtForwarderMockRecorder struct {
	mock *MockEvtForwarder
}

// NewMockEvtForwarder creates a new mock instance
func NewMockEvtForwarder(ctrl *gomock.Controller) *MockEvtForwarder {
	mock := &MockEvtForwarder{ctrl: ctrl}
	mock.recorder = &MockEvtForwarderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEvtForwarder) EXPECT() *MockEvtForwarderMockRecorder {
	return m.recorder
}

// Ack mocks base method
func (m *MockEvtForwarder) Ack(arg0 *proto.ChainEvent) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ack", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Ack indicates an expected call of Ack
func (mr *MockEvtForwarderMockRecorder) Ack(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*MockEvtForwarder)(nil).Ack), arg0)
}
