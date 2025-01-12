// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/datanode/api (interfaces: VegaTime)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockVegaTime is a mock of VegaTime interface.
type MockVegaTime struct {
	ctrl     *gomock.Controller
	recorder *MockVegaTimeMockRecorder
}

// MockVegaTimeMockRecorder is the mock recorder for MockVegaTime.
type MockVegaTimeMockRecorder struct {
	mock *MockVegaTime
}

// NewMockVegaTime creates a new mock instance.
func NewMockVegaTime(ctrl *gomock.Controller) *MockVegaTime {
	mock := &MockVegaTime{ctrl: ctrl}
	mock.recorder = &MockVegaTimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVegaTime) EXPECT() *MockVegaTimeMockRecorder {
	return m.recorder
}

// GetTimeNow mocks base method.
func (m *MockVegaTime) GetTimeNow() (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeNow")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimeNow indicates an expected call of GetTimeNow.
func (mr *MockVegaTimeMockRecorder) GetTimeNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeNow", reflect.TypeOf((*MockVegaTime)(nil).GetTimeNow))
}
