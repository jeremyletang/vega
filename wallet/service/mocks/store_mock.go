// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/wallet/service (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	service "code.vegaprotocol.io/vega/wallet/service"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// RSAKeysExists mocks base method.
func (m *MockStore) RSAKeysExists() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RSAKeysExists")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RSAKeysExists indicates an expected call of RSAKeysExists.
func (mr *MockStoreMockRecorder) RSAKeysExists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RSAKeysExists", reflect.TypeOf((*MockStore)(nil).RSAKeysExists))
}

// SaveRSAKeys mocks base method.
func (m *MockStore) SaveRSAKeys(arg0 *service.RSAKeys) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveRSAKeys", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveRSAKeys indicates an expected call of SaveRSAKeys.
func (mr *MockStoreMockRecorder) SaveRSAKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveRSAKeys", reflect.TypeOf((*MockStore)(nil).SaveRSAKeys), arg0)
}
