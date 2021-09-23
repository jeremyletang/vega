// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: GovernanceEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	vega "code.vegaprotocol.io/protos/vega"
	governance "code.vegaprotocol.io/vega/governance"
	types "code.vegaprotocol.io/vega/types"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockGovernanceEngine is a mock of GovernanceEngine interface
type MockGovernanceEngine struct {
	ctrl     *gomock.Controller
	recorder *MockGovernanceEngineMockRecorder
}

// MockGovernanceEngineMockRecorder is the mock recorder for MockGovernanceEngine
type MockGovernanceEngineMockRecorder struct {
	mock *MockGovernanceEngine
}

// NewMockGovernanceEngine creates a new mock instance
func NewMockGovernanceEngine(ctrl *gomock.Controller) *MockGovernanceEngine {
	mock := &MockGovernanceEngine{ctrl: ctrl}
	mock.recorder = &MockGovernanceEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGovernanceEngine) EXPECT() *MockGovernanceEngineMockRecorder {
	return m.recorder
}

// AddVote mocks base method
func (m *MockGovernanceEngine) AddVote(arg0 context.Context, arg1 types.VoteSubmission, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVote", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVote indicates an expected call of AddVote
func (mr *MockGovernanceEngineMockRecorder) AddVote(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVote", reflect.TypeOf((*MockGovernanceEngine)(nil).AddVote), arg0, arg1, arg2)
}

// Hash mocks base method
func (m *MockGovernanceEngine) Hash() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Hash indicates an expected call of Hash
func (mr *MockGovernanceEngineMockRecorder) Hash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockGovernanceEngine)(nil).Hash))
}

// OnChainTimeUpdate mocks base method
func (m *MockGovernanceEngine) OnChainTimeUpdate(arg0 context.Context, arg1 time.Time) ([]*governance.ToEnact, []*governance.VoteClosed) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnChainTimeUpdate", arg0, arg1)
	ret0, _ := ret[0].([]*governance.ToEnact)
	ret1, _ := ret[1].([]*governance.VoteClosed)
	return ret0, ret1
}

// OnChainTimeUpdate indicates an expected call of OnChainTimeUpdate
func (mr *MockGovernanceEngineMockRecorder) OnChainTimeUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChainTimeUpdate", reflect.TypeOf((*MockGovernanceEngine)(nil).OnChainTimeUpdate), arg0, arg1)
}

// RejectProposal mocks base method
func (m *MockGovernanceEngine) RejectProposal(arg0 context.Context, arg1 *types.Proposal, arg2 vega.ProposalError, arg3 error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RejectProposal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RejectProposal indicates an expected call of RejectProposal
func (mr *MockGovernanceEngineMockRecorder) RejectProposal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejectProposal", reflect.TypeOf((*MockGovernanceEngine)(nil).RejectProposal), arg0, arg1, arg2, arg3)
}

// SubmitProposal mocks base method
func (m *MockGovernanceEngine) SubmitProposal(arg0 context.Context, arg1 types.ProposalSubmission, arg2, arg3 string) (*governance.ToSubmit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitProposal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*governance.ToSubmit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitProposal indicates an expected call of SubmitProposal
func (mr *MockGovernanceEngineMockRecorder) SubmitProposal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitProposal", reflect.TypeOf((*MockGovernanceEngine)(nil).SubmitProposal), arg0, arg1, arg2, arg3)
}
