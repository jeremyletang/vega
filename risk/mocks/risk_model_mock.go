// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/risk (interfaces: Model)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "code.vegaprotocol.io/vega/types"
	num "code.vegaprotocol.io/vega/types/num"
	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
)

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// CalculateRiskFactors mocks base method.
func (m *MockModel) CalculateRiskFactors() *types.RiskFactor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateRiskFactors")
	ret0, _ := ret[0].(*types.RiskFactor)
	return ret0
}

// CalculateRiskFactors indicates an expected call of CalculateRiskFactors.
func (mr *MockModelMockRecorder) CalculateRiskFactors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateRiskFactors", reflect.TypeOf((*MockModel)(nil).CalculateRiskFactors))
}

// DefaultRiskFactors mocks base method.
func (m *MockModel) DefaultRiskFactors() *types.RiskFactor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultRiskFactors")
	ret0, _ := ret[0].(*types.RiskFactor)
	return ret0
}

// DefaultRiskFactors indicates an expected call of DefaultRiskFactors.
func (mr *MockModelMockRecorder) DefaultRiskFactors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultRiskFactors", reflect.TypeOf((*MockModel)(nil).DefaultRiskFactors))
}

// GetProjectionHorizon mocks base method.
func (m *MockModel) GetProjectionHorizon() decimal.Decimal {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectionHorizon")
	ret0, _ := ret[0].(decimal.Decimal)
	return ret0
}

// GetProjectionHorizon indicates an expected call of GetProjectionHorizon.
func (mr *MockModelMockRecorder) GetProjectionHorizon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectionHorizon", reflect.TypeOf((*MockModel)(nil).GetProjectionHorizon))
}

// PriceRange mocks base method.
func (m *MockModel) PriceRange(arg0, arg1, arg2 decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PriceRange", arg0, arg1, arg2)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(decimal.Decimal)
	return ret0, ret1
}

// PriceRange indicates an expected call of PriceRange.
func (mr *MockModelMockRecorder) PriceRange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PriceRange", reflect.TypeOf((*MockModel)(nil).PriceRange), arg0, arg1, arg2)
}

// ProbabilityOfTrading mocks base method.
func (m *MockModel) ProbabilityOfTrading(arg0, arg1 *num.Uint, arg2, arg3, arg4 decimal.Decimal, arg5, arg6 bool) decimal.Decimal {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProbabilityOfTrading", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(decimal.Decimal)
	return ret0
}

// ProbabilityOfTrading indicates an expected call of ProbabilityOfTrading.
func (mr *MockModelMockRecorder) ProbabilityOfTrading(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProbabilityOfTrading", reflect.TypeOf((*MockModel)(nil).ProbabilityOfTrading), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
