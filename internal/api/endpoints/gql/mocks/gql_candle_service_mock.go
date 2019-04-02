// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/api/endpoints/gql (interfaces: CandleService)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockCandleService is a mock of CandleService interface
type MockCandleService struct {
	ctrl     *gomock.Controller
	recorder *MockCandleServiceMockRecorder
}

// MockCandleServiceMockRecorder is the mock recorder for MockCandleService
type MockCandleServiceMockRecorder struct {
	mock *MockCandleService
}

// NewMockCandleService creates a new mock instance
func NewMockCandleService(ctrl *gomock.Controller) *MockCandleService {
	mock := &MockCandleService{ctrl: ctrl}
	mock.recorder = &MockCandleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCandleService) EXPECT() *MockCandleServiceMockRecorder {
	return m.recorder
}

// GetCandles mocks base method
func (m *MockCandleService) GetCandles(arg0 context.Context, arg1 string, arg2 time.Time, arg3 proto.Interval) ([]*proto.Candle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCandles", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*proto.Candle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandles indicates an expected call of GetCandles
func (mr *MockCandleServiceMockRecorder) GetCandles(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandles", reflect.TypeOf((*MockCandleService)(nil).GetCandles), arg0, arg1, arg2, arg3)
}

// ObserveCandles mocks base method
func (m *MockCandleService) ObserveCandles(arg0 context.Context, arg1 int, arg2 *string, arg3 *proto.Interval) (<-chan *proto.Candle, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObserveCandles", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan *proto.Candle)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// ObserveCandles indicates an expected call of ObserveCandles
func (mr *MockCandleServiceMockRecorder) ObserveCandles(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveCandles", reflect.TypeOf((*MockCandleService)(nil).ObserveCandles), arg0, arg1, arg2, arg3)
}
