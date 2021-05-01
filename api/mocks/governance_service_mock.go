// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/api (interfaces: GovernanceService)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	v1 "code.vegaprotocol.io/vega/proto/commands/v1"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGovernanceService is a mock of GovernanceService interface
type MockGovernanceService struct {
	ctrl     *gomock.Controller
	recorder *MockGovernanceServiceMockRecorder
}

// MockGovernanceServiceMockRecorder is the mock recorder for MockGovernanceService
type MockGovernanceServiceMockRecorder struct {
	mock *MockGovernanceService
}

// NewMockGovernanceService creates a new mock instance
func NewMockGovernanceService(ctrl *gomock.Controller) *MockGovernanceService {
	mock := &MockGovernanceService{ctrl: ctrl}
	mock.recorder = &MockGovernanceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGovernanceService) EXPECT() *MockGovernanceServiceMockRecorder {
	return m.recorder
}

// PrepareProposal mocks base method
func (m *MockGovernanceService) PrepareProposal(arg0 context.Context, arg1 string, arg2 *proto.ProposalTerms) (*v1.ProposalSubmission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareProposal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ProposalSubmission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareProposal indicates an expected call of PrepareProposal
func (mr *MockGovernanceServiceMockRecorder) PrepareProposal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareProposal", reflect.TypeOf((*MockGovernanceService)(nil).PrepareProposal), arg0, arg1, arg2)
}

// PrepareVote mocks base method
func (m *MockGovernanceService) PrepareVote(arg0 *v1.VoteSubmission) (*v1.VoteSubmission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareVote", arg0)
	ret0, _ := ret[0].(*v1.VoteSubmission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareVote indicates an expected call of PrepareVote
func (mr *MockGovernanceServiceMockRecorder) PrepareVote(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareVote", reflect.TypeOf((*MockGovernanceService)(nil).PrepareVote), arg0)
}
