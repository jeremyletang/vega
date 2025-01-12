// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/statevar (interfaces: Topology)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	validators "code.vegaprotocol.io/vega/core/validators"
	gomock "github.com/golang/mock/gomock"
)

// MockTopology is a mock of Topology interface.
type MockTopology struct {
	ctrl     *gomock.Controller
	recorder *MockTopologyMockRecorder
}

// MockTopologyMockRecorder is the mock recorder for MockTopology.
type MockTopologyMockRecorder struct {
	mock *MockTopology
}

// NewMockTopology creates a new mock instance.
func NewMockTopology(ctrl *gomock.Controller) *MockTopology {
	mock := &MockTopology{ctrl: ctrl}
	mock.recorder = &MockTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTopology) EXPECT() *MockTopologyMockRecorder {
	return m.recorder
}

// AllNodeIDs mocks base method.
func (m *MockTopology) AllNodeIDs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllNodeIDs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// AllNodeIDs indicates an expected call of AllNodeIDs.
func (mr *MockTopologyMockRecorder) AllNodeIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllNodeIDs", reflect.TypeOf((*MockTopology)(nil).AllNodeIDs))
}

// Get mocks base method.
func (m *MockTopology) Get(arg0 string) *validators.ValidatorData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*validators.ValidatorData)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockTopologyMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTopology)(nil).Get), arg0)
}

// IsValidator mocks base method.
func (m *MockTopology) IsValidator() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidator")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidator indicates an expected call of IsValidator.
func (mr *MockTopologyMockRecorder) IsValidator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidator", reflect.TypeOf((*MockTopology)(nil).IsValidator))
}

// IsValidatorVegaPubKey mocks base method.
func (m *MockTopology) IsValidatorVegaPubKey(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidatorVegaPubKey", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidatorVegaPubKey indicates an expected call of IsValidatorVegaPubKey.
func (mr *MockTopologyMockRecorder) IsValidatorVegaPubKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidatorVegaPubKey", reflect.TypeOf((*MockTopology)(nil).IsValidatorVegaPubKey), arg0)
}

// SelfNodeID mocks base method.
func (m *MockTopology) SelfNodeID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelfNodeID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SelfNodeID indicates an expected call of SelfNodeID.
func (mr *MockTopologyMockRecorder) SelfNodeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelfNodeID", reflect.TypeOf((*MockTopology)(nil).SelfNodeID))
}
