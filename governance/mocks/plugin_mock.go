// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/governance (interfaces: Plugin)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
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

// GetAllGovernanceData mocks base method
func (m *MockPlugin) GetAllGovernanceData(arg0 *proto.Proposal_State) []*proto.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllGovernanceData", arg0)
	ret0, _ := ret[0].([]*proto.GovernanceData)
	return ret0
}

// GetAllGovernanceData indicates an expected call of GetAllGovernanceData
func (mr *MockPluginMockRecorder) GetAllGovernanceData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGovernanceData", reflect.TypeOf((*MockPlugin)(nil).GetAllGovernanceData), arg0)
}

// GetNetworkParametersProposals mocks base method
func (m *MockPlugin) GetNetworkParametersProposals(arg0 *proto.Proposal_State) []*proto.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkParametersProposals", arg0)
	ret0, _ := ret[0].([]*proto.GovernanceData)
	return ret0
}

// GetNetworkParametersProposals indicates an expected call of GetNetworkParametersProposals
func (mr *MockPluginMockRecorder) GetNetworkParametersProposals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkParametersProposals", reflect.TypeOf((*MockPlugin)(nil).GetNetworkParametersProposals), arg0)
}

// GetNewAssetProposals mocks base method
func (m *MockPlugin) GetNewAssetProposals(arg0 *proto.Proposal_State) []*proto.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewAssetProposals", arg0)
	ret0, _ := ret[0].([]*proto.GovernanceData)
	return ret0
}

// GetNewAssetProposals indicates an expected call of GetNewAssetProposals
func (mr *MockPluginMockRecorder) GetNewAssetProposals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewAssetProposals", reflect.TypeOf((*MockPlugin)(nil).GetNewAssetProposals), arg0)
}

// GetNewMarketProposals mocks base method
func (m *MockPlugin) GetNewMarketProposals(arg0 *proto.Proposal_State) []*proto.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewMarketProposals", arg0)
	ret0, _ := ret[0].([]*proto.GovernanceData)
	return ret0
}

// GetNewMarketProposals indicates an expected call of GetNewMarketProposals
func (mr *MockPluginMockRecorder) GetNewMarketProposals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewMarketProposals", reflect.TypeOf((*MockPlugin)(nil).GetNewMarketProposals), arg0)
}

// GetProposalByID mocks base method
func (m *MockPlugin) GetProposalByID(arg0 string) (*proto.GovernanceData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposalByID", arg0)
	ret0, _ := ret[0].(*proto.GovernanceData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProposalByID indicates an expected call of GetProposalByID
func (mr *MockPluginMockRecorder) GetProposalByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposalByID", reflect.TypeOf((*MockPlugin)(nil).GetProposalByID), arg0)
}

// GetProposalByReference mocks base method
func (m *MockPlugin) GetProposalByReference(arg0 string) (*proto.GovernanceData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposalByReference", arg0)
	ret0, _ := ret[0].(*proto.GovernanceData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProposalByReference indicates an expected call of GetProposalByReference
func (mr *MockPluginMockRecorder) GetProposalByReference(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposalByReference", reflect.TypeOf((*MockPlugin)(nil).GetProposalByReference), arg0)
}

// GetProposalsByParty mocks base method
func (m *MockPlugin) GetProposalsByParty(arg0 string, arg1 *proto.Proposal_State) []*proto.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposalsByParty", arg0, arg1)
	ret0, _ := ret[0].([]*proto.GovernanceData)
	return ret0
}

// GetProposalsByParty indicates an expected call of GetProposalsByParty
func (mr *MockPluginMockRecorder) GetProposalsByParty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposalsByParty", reflect.TypeOf((*MockPlugin)(nil).GetProposalsByParty), arg0, arg1)
}

// GetUpdateMarketProposals mocks base method
func (m *MockPlugin) GetUpdateMarketProposals(arg0 string, arg1 *proto.Proposal_State) []*proto.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateMarketProposals", arg0, arg1)
	ret0, _ := ret[0].([]*proto.GovernanceData)
	return ret0
}

// GetUpdateMarketProposals indicates an expected call of GetUpdateMarketProposals
func (mr *MockPluginMockRecorder) GetUpdateMarketProposals(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateMarketProposals", reflect.TypeOf((*MockPlugin)(nil).GetUpdateMarketProposals), arg0, arg1)
}

// GetVotesByParty mocks base method
func (m *MockPlugin) GetVotesByParty(arg0 string) []*proto.Vote {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVotesByParty", arg0)
	ret0, _ := ret[0].([]*proto.Vote)
	return ret0
}

// GetVotesByParty indicates an expected call of GetVotesByParty
func (mr *MockPluginMockRecorder) GetVotesByParty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotesByParty", reflect.TypeOf((*MockPlugin)(nil).GetVotesByParty), arg0)
}

// SubscribeAll mocks base method
func (m *MockPlugin) SubscribeAll() (<-chan []proto.GovernanceData, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeAll")
	ret0, _ := ret[0].(<-chan []proto.GovernanceData)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// SubscribeAll indicates an expected call of SubscribeAll
func (mr *MockPluginMockRecorder) SubscribeAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeAll", reflect.TypeOf((*MockPlugin)(nil).SubscribeAll))
}

// SubscribePartyProposals mocks base method
func (m *MockPlugin) SubscribePartyProposals(arg0 string) (<-chan []proto.GovernanceData, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribePartyProposals", arg0)
	ret0, _ := ret[0].(<-chan []proto.GovernanceData)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// SubscribePartyProposals indicates an expected call of SubscribePartyProposals
func (mr *MockPluginMockRecorder) SubscribePartyProposals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribePartyProposals", reflect.TypeOf((*MockPlugin)(nil).SubscribePartyProposals), arg0)
}

// SubscribePartyVotes mocks base method
func (m *MockPlugin) SubscribePartyVotes(arg0 string) (<-chan []proto.Vote, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribePartyVotes", arg0)
	ret0, _ := ret[0].(<-chan []proto.Vote)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// SubscribePartyVotes indicates an expected call of SubscribePartyVotes
func (mr *MockPluginMockRecorder) SubscribePartyVotes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribePartyVotes", reflect.TypeOf((*MockPlugin)(nil).SubscribePartyVotes), arg0)
}

// SubscribeProposalVotes mocks base method
func (m *MockPlugin) SubscribeProposalVotes(arg0 string) (<-chan []proto.Vote, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeProposalVotes", arg0)
	ret0, _ := ret[0].(<-chan []proto.Vote)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// SubscribeProposalVotes indicates an expected call of SubscribeProposalVotes
func (mr *MockPluginMockRecorder) SubscribeProposalVotes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeProposalVotes", reflect.TypeOf((*MockPlugin)(nil).SubscribeProposalVotes), arg0)
}

// UnsubscribeAll mocks base method
func (m *MockPlugin) UnsubscribeAll(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnsubscribeAll", arg0)
}

// UnsubscribeAll indicates an expected call of UnsubscribeAll
func (mr *MockPluginMockRecorder) UnsubscribeAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeAll", reflect.TypeOf((*MockPlugin)(nil).UnsubscribeAll), arg0)
}

// UnsubscribePartyProposals mocks base method
func (m *MockPlugin) UnsubscribePartyProposals(arg0 string, arg1 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnsubscribePartyProposals", arg0, arg1)
}

// UnsubscribePartyProposals indicates an expected call of UnsubscribePartyProposals
func (mr *MockPluginMockRecorder) UnsubscribePartyProposals(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribePartyProposals", reflect.TypeOf((*MockPlugin)(nil).UnsubscribePartyProposals), arg0, arg1)
}

// UnsubscribePartyVotes mocks base method
func (m *MockPlugin) UnsubscribePartyVotes(arg0 string, arg1 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnsubscribePartyVotes", arg0, arg1)
}

// UnsubscribePartyVotes indicates an expected call of UnsubscribePartyVotes
func (mr *MockPluginMockRecorder) UnsubscribePartyVotes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribePartyVotes", reflect.TypeOf((*MockPlugin)(nil).UnsubscribePartyVotes), arg0, arg1)
}

// UnsubscribeProposalVotes mocks base method
func (m *MockPlugin) UnsubscribeProposalVotes(arg0 string, arg1 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnsubscribeProposalVotes", arg0, arg1)
}

// UnsubscribeProposalVotes indicates an expected call of UnsubscribeProposalVotes
func (mr *MockPluginMockRecorder) UnsubscribeProposalVotes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeProposalVotes", reflect.TypeOf((*MockPlugin)(nil).UnsubscribeProposalVotes), arg0, arg1)
}
