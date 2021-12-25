// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/validators (interfaces: ValidatorTopology)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockValidatorTopology is a mock of ValidatorTopology interface
type MockValidatorTopology struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorTopologyMockRecorder
}

// MockValidatorTopologyMockRecorder is the mock recorder for MockValidatorTopology
type MockValidatorTopologyMockRecorder struct {
	mock *MockValidatorTopology
}

// NewMockValidatorTopology creates a new mock instance
func NewMockValidatorTopology(ctrl *gomock.Controller) *MockValidatorTopology {
	mock := &MockValidatorTopology{ctrl: ctrl}
	mock.recorder = &MockValidatorTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidatorTopology) EXPECT() *MockValidatorTopologyMockRecorder {
	return m.recorder
}

// AllNodeIDs mocks base method
func (m *MockValidatorTopology) AllNodeIDs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllNodeIDs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// AllNodeIDs indicates an expected call of AllNodeIDs
func (mr *MockValidatorTopologyMockRecorder) AllNodeIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllNodeIDs", reflect.TypeOf((*MockValidatorTopology)(nil).AllNodeIDs))
}

// IsValidator mocks base method
func (m *MockValidatorTopology) IsValidator() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidator")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidator indicates an expected call of IsValidator
func (mr *MockValidatorTopologyMockRecorder) IsValidator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidator", reflect.TypeOf((*MockValidatorTopology)(nil).IsValidator))
}

// IsValidatorNodeID mocks base method
func (m *MockValidatorTopology) IsValidatorNodeID(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidatorNodeID", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidatorNodeID indicates an expected call of IsValidatorNodeID
func (mr *MockValidatorTopologyMockRecorder) IsValidatorNodeID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidatorNodeID", reflect.TypeOf((*MockValidatorTopology)(nil).IsValidatorNodeID), arg0)
}

// Len mocks base method
func (m *MockValidatorTopology) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len
func (mr *MockValidatorTopologyMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockValidatorTopology)(nil).Len))
}

// SelfNodeID mocks base method
func (m *MockValidatorTopology) SelfNodeID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelfNodeID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SelfNodeID indicates an expected call of SelfNodeID
func (mr *MockValidatorTopologyMockRecorder) SelfNodeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelfNodeID", reflect.TypeOf((*MockValidatorTopology)(nil).SelfNodeID))
}
