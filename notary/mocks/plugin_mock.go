// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/notary (interfaces: Plugin)

// Package mocks is a generated GoMock package.
package mocks

import (
	v1 "code.vegaprotocol.io/vega/proto/commands/v1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPlugin is a mock of Plugin interface
type MockPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockPluginMockRecorder
}

// MockPluginMockRecorder is the mock recorder for MockPlugin
type MockPluginMockRecorder struct {
	mock *MockPlugin
}

// NewMockPlugin creates a new mock instance
func NewMockPlugin(ctrl *gomock.Controller) *MockPlugin {
	mock := &MockPlugin{ctrl: ctrl}
	mock.recorder = &MockPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlugin) EXPECT() *MockPluginMockRecorder {
	return m.recorder
}

// GetByID mocks base method
func (m *MockPlugin) GetByID(arg0 string) ([]v1.NodeSignature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].([]v1.NodeSignature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockPluginMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockPlugin)(nil).GetByID), arg0)
}
