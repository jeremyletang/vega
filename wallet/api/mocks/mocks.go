// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/wallet/api (interfaces: WalletStore,NetworkStore,Interactor,ConnectionsManager,SpamHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	v1 "code.vegaprotocol.io/vega/protos/vega/commands/v1"
	v10 "code.vegaprotocol.io/vega/protos/vega/wallet/v1"
	api "code.vegaprotocol.io/vega/wallet/api"
	types "code.vegaprotocol.io/vega/wallet/api/node/types"
	network "code.vegaprotocol.io/vega/wallet/network"
	wallet "code.vegaprotocol.io/vega/wallet/wallet"
	gomock "github.com/golang/mock/gomock"
)

// MockWalletStore is a mock of WalletStore interface.
type MockWalletStore struct {
	ctrl     *gomock.Controller
	recorder *MockWalletStoreMockRecorder
}

// MockWalletStoreMockRecorder is the mock recorder for MockWalletStore.
type MockWalletStoreMockRecorder struct {
	mock *MockWalletStore
}

// NewMockWalletStore creates a new mock instance.
func NewMockWalletStore(ctrl *gomock.Controller) *MockWalletStore {
	mock := &MockWalletStore{ctrl: ctrl}
	mock.recorder = &MockWalletStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletStore) EXPECT() *MockWalletStoreMockRecorder {
	return m.recorder
}

// CreateWallet mocks base method.
func (m *MockWalletStore) CreateWallet(arg0 context.Context, arg1 wallet.Wallet, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockWalletStoreMockRecorder) CreateWallet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockWalletStore)(nil).CreateWallet), arg0, arg1, arg2)
}

// DeleteWallet mocks base method.
func (m *MockWalletStore) DeleteWallet(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWallet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWallet indicates an expected call of DeleteWallet.
func (mr *MockWalletStoreMockRecorder) DeleteWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWallet", reflect.TypeOf((*MockWalletStore)(nil).DeleteWallet), arg0, arg1)
}

// GetWallet mocks base method.
func (m *MockWalletStore) GetWallet(arg0 context.Context, arg1 string) (wallet.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWallet", arg0, arg1)
	ret0, _ := ret[0].(wallet.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWallet indicates an expected call of GetWallet.
func (mr *MockWalletStoreMockRecorder) GetWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWallet", reflect.TypeOf((*MockWalletStore)(nil).GetWallet), arg0, arg1)
}

// GetWalletPath mocks base method.
func (m *MockWalletStore) GetWalletPath(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletPath", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWalletPath indicates an expected call of GetWalletPath.
func (mr *MockWalletStoreMockRecorder) GetWalletPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletPath", reflect.TypeOf((*MockWalletStore)(nil).GetWalletPath), arg0)
}

// IsWalletAlreadyUnlocked mocks base method.
func (m *MockWalletStore) IsWalletAlreadyUnlocked(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWalletAlreadyUnlocked", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsWalletAlreadyUnlocked indicates an expected call of IsWalletAlreadyUnlocked.
func (mr *MockWalletStoreMockRecorder) IsWalletAlreadyUnlocked(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWalletAlreadyUnlocked", reflect.TypeOf((*MockWalletStore)(nil).IsWalletAlreadyUnlocked), arg0, arg1)
}

// ListWallets mocks base method.
func (m *MockWalletStore) ListWallets(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWallets", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWallets indicates an expected call of ListWallets.
func (mr *MockWalletStoreMockRecorder) ListWallets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWallets", reflect.TypeOf((*MockWalletStore)(nil).ListWallets), arg0)
}

// LockWallet mocks base method.
func (m *MockWalletStore) LockWallet(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockWallet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockWallet indicates an expected call of LockWallet.
func (mr *MockWalletStoreMockRecorder) LockWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockWallet", reflect.TypeOf((*MockWalletStore)(nil).LockWallet), arg0, arg1)
}

// RenameWallet mocks base method.
func (m *MockWalletStore) RenameWallet(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameWallet", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameWallet indicates an expected call of RenameWallet.
func (mr *MockWalletStoreMockRecorder) RenameWallet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameWallet", reflect.TypeOf((*MockWalletStore)(nil).RenameWallet), arg0, arg1, arg2)
}

// UnlockWallet mocks base method.
func (m *MockWalletStore) UnlockWallet(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlockWallet", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnlockWallet indicates an expected call of UnlockWallet.
func (mr *MockWalletStoreMockRecorder) UnlockWallet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockWallet", reflect.TypeOf((*MockWalletStore)(nil).UnlockWallet), arg0, arg1, arg2)
}

// UpdatePassphrase mocks base method.
func (m *MockWalletStore) UpdatePassphrase(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassphrase", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassphrase indicates an expected call of UpdatePassphrase.
func (mr *MockWalletStoreMockRecorder) UpdatePassphrase(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassphrase", reflect.TypeOf((*MockWalletStore)(nil).UpdatePassphrase), arg0, arg1, arg2)
}

// UpdateWallet mocks base method.
func (m *MockWalletStore) UpdateWallet(arg0 context.Context, arg1 wallet.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWallet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWallet indicates an expected call of UpdateWallet.
func (mr *MockWalletStoreMockRecorder) UpdateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWallet", reflect.TypeOf((*MockWalletStore)(nil).UpdateWallet), arg0, arg1)
}

// WalletExists mocks base method.
func (m *MockWalletStore) WalletExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletExists indicates an expected call of WalletExists.
func (mr *MockWalletStoreMockRecorder) WalletExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletExists", reflect.TypeOf((*MockWalletStore)(nil).WalletExists), arg0, arg1)
}

// MockNetworkStore is a mock of NetworkStore interface.
type MockNetworkStore struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkStoreMockRecorder
}

// MockNetworkStoreMockRecorder is the mock recorder for MockNetworkStore.
type MockNetworkStoreMockRecorder struct {
	mock *MockNetworkStore
}

// NewMockNetworkStore creates a new mock instance.
func NewMockNetworkStore(ctrl *gomock.Controller) *MockNetworkStore {
	mock := &MockNetworkStore{ctrl: ctrl}
	mock.recorder = &MockNetworkStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkStore) EXPECT() *MockNetworkStoreMockRecorder {
	return m.recorder
}

// DeleteNetwork mocks base method.
func (m *MockNetworkStore) DeleteNetwork(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetwork", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNetwork indicates an expected call of DeleteNetwork.
func (mr *MockNetworkStoreMockRecorder) DeleteNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetwork", reflect.TypeOf((*MockNetworkStore)(nil).DeleteNetwork), arg0)
}

// GetNetwork mocks base method.
func (m *MockNetworkStore) GetNetwork(arg0 string) (*network.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetwork", arg0)
	ret0, _ := ret[0].(*network.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetwork indicates an expected call of GetNetwork.
func (mr *MockNetworkStoreMockRecorder) GetNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetwork", reflect.TypeOf((*MockNetworkStore)(nil).GetNetwork), arg0)
}

// GetNetworkPath mocks base method.
func (m *MockNetworkStore) GetNetworkPath(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkPath", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNetworkPath indicates an expected call of GetNetworkPath.
func (mr *MockNetworkStoreMockRecorder) GetNetworkPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkPath", reflect.TypeOf((*MockNetworkStore)(nil).GetNetworkPath), arg0)
}

// ListNetworks mocks base method.
func (m *MockNetworkStore) ListNetworks() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNetworks")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNetworks indicates an expected call of ListNetworks.
func (mr *MockNetworkStoreMockRecorder) ListNetworks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetworks", reflect.TypeOf((*MockNetworkStore)(nil).ListNetworks))
}

// NetworkExists mocks base method.
func (m *MockNetworkStore) NetworkExists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkExists indicates an expected call of NetworkExists.
func (mr *MockNetworkStoreMockRecorder) NetworkExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkExists", reflect.TypeOf((*MockNetworkStore)(nil).NetworkExists), arg0)
}

// RenameNetwork mocks base method.
func (m *MockNetworkStore) RenameNetwork(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameNetwork", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameNetwork indicates an expected call of RenameNetwork.
func (mr *MockNetworkStoreMockRecorder) RenameNetwork(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameNetwork", reflect.TypeOf((*MockNetworkStore)(nil).RenameNetwork), arg0, arg1)
}

// SaveNetwork mocks base method.
func (m *MockNetworkStore) SaveNetwork(arg0 *network.Network) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNetwork", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNetwork indicates an expected call of SaveNetwork.
func (mr *MockNetworkStoreMockRecorder) SaveNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNetwork", reflect.TypeOf((*MockNetworkStore)(nil).SaveNetwork), arg0)
}

// MockInteractor is a mock of Interactor interface.
type MockInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockInteractorMockRecorder
}

// MockInteractorMockRecorder is the mock recorder for MockInteractor.
type MockInteractorMockRecorder struct {
	mock *MockInteractor
}

// NewMockInteractor creates a new mock instance.
func NewMockInteractor(ctrl *gomock.Controller) *MockInteractor {
	mock := &MockInteractor{ctrl: ctrl}
	mock.recorder = &MockInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInteractor) EXPECT() *MockInteractorMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockInteractor) Log(arg0 context.Context, arg1 string, arg2 api.LogType, arg3 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Log", arg0, arg1, arg2, arg3)
}

// Log indicates an expected call of Log.
func (mr *MockInteractorMockRecorder) Log(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockInteractor)(nil).Log), arg0, arg1, arg2, arg3)
}

// NotifyError mocks base method.
func (m *MockInteractor) NotifyError(arg0 context.Context, arg1 string, arg2 api.ErrorType, arg3 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyError", arg0, arg1, arg2, arg3)
}

// NotifyError indicates an expected call of NotifyError.
func (mr *MockInteractorMockRecorder) NotifyError(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyError", reflect.TypeOf((*MockInteractor)(nil).NotifyError), arg0, arg1, arg2, arg3)
}

// NotifyFailedTransaction mocks base method.
func (m *MockInteractor) NotifyFailedTransaction(arg0 context.Context, arg1 string, arg2 byte, arg3, arg4 string, arg5 error, arg6 time.Time, arg7 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyFailedTransaction", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// NotifyFailedTransaction indicates an expected call of NotifyFailedTransaction.
func (mr *MockInteractorMockRecorder) NotifyFailedTransaction(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyFailedTransaction", reflect.TypeOf((*MockInteractor)(nil).NotifyFailedTransaction), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// NotifyInteractionSessionBegan mocks base method.
func (m *MockInteractor) NotifyInteractionSessionBegan(arg0 context.Context, arg1 string, arg2 api.WorkflowType, arg3 byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyInteractionSessionBegan", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyInteractionSessionBegan indicates an expected call of NotifyInteractionSessionBegan.
func (mr *MockInteractorMockRecorder) NotifyInteractionSessionBegan(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyInteractionSessionBegan", reflect.TypeOf((*MockInteractor)(nil).NotifyInteractionSessionBegan), arg0, arg1, arg2, arg3)
}

// NotifyInteractionSessionEnded mocks base method.
func (m *MockInteractor) NotifyInteractionSessionEnded(arg0 context.Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyInteractionSessionEnded", arg0, arg1)
}

// NotifyInteractionSessionEnded indicates an expected call of NotifyInteractionSessionEnded.
func (mr *MockInteractorMockRecorder) NotifyInteractionSessionEnded(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyInteractionSessionEnded", reflect.TypeOf((*MockInteractor)(nil).NotifyInteractionSessionEnded), arg0, arg1)
}

// NotifySuccessfulRequest mocks base method.
func (m *MockInteractor) NotifySuccessfulRequest(arg0 context.Context, arg1 string, arg2 byte, arg3 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifySuccessfulRequest", arg0, arg1, arg2, arg3)
}

// NotifySuccessfulRequest indicates an expected call of NotifySuccessfulRequest.
func (mr *MockInteractorMockRecorder) NotifySuccessfulRequest(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifySuccessfulRequest", reflect.TypeOf((*MockInteractor)(nil).NotifySuccessfulRequest), arg0, arg1, arg2, arg3)
}

// NotifySuccessfulTransaction mocks base method.
func (m *MockInteractor) NotifySuccessfulTransaction(arg0 context.Context, arg1 string, arg2 byte, arg3, arg4, arg5 string, arg6 time.Time, arg7 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifySuccessfulTransaction", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// NotifySuccessfulTransaction indicates an expected call of NotifySuccessfulTransaction.
func (mr *MockInteractorMockRecorder) NotifySuccessfulTransaction(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifySuccessfulTransaction", reflect.TypeOf((*MockInteractor)(nil).NotifySuccessfulTransaction), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// RequestPassphrase mocks base method.
func (m *MockInteractor) RequestPassphrase(arg0 context.Context, arg1 string, arg2 byte, arg3, arg4 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestPassphrase", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestPassphrase indicates an expected call of RequestPassphrase.
func (mr *MockInteractorMockRecorder) RequestPassphrase(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestPassphrase", reflect.TypeOf((*MockInteractor)(nil).RequestPassphrase), arg0, arg1, arg2, arg3, arg4)
}

// RequestPermissionsReview mocks base method.
func (m *MockInteractor) RequestPermissionsReview(arg0 context.Context, arg1 string, arg2 byte, arg3, arg4 string, arg5 map[string]string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestPermissionsReview", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestPermissionsReview indicates an expected call of RequestPermissionsReview.
func (mr *MockInteractorMockRecorder) RequestPermissionsReview(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestPermissionsReview", reflect.TypeOf((*MockInteractor)(nil).RequestPermissionsReview), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RequestTransactionReviewForSending mocks base method.
func (m *MockInteractor) RequestTransactionReviewForSending(arg0 context.Context, arg1 string, arg2 byte, arg3, arg4, arg5, arg6 string, arg7 time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestTransactionReviewForSending", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestTransactionReviewForSending indicates an expected call of RequestTransactionReviewForSending.
func (mr *MockInteractorMockRecorder) RequestTransactionReviewForSending(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestTransactionReviewForSending", reflect.TypeOf((*MockInteractor)(nil).RequestTransactionReviewForSending), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// RequestTransactionReviewForSigning mocks base method.
func (m *MockInteractor) RequestTransactionReviewForSigning(arg0 context.Context, arg1 string, arg2 byte, arg3, arg4, arg5, arg6 string, arg7 time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestTransactionReviewForSigning", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestTransactionReviewForSigning indicates an expected call of RequestTransactionReviewForSigning.
func (mr *MockInteractorMockRecorder) RequestTransactionReviewForSigning(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestTransactionReviewForSigning", reflect.TypeOf((*MockInteractor)(nil).RequestTransactionReviewForSigning), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// RequestWalletConnectionReview mocks base method.
func (m *MockInteractor) RequestWalletConnectionReview(arg0 context.Context, arg1 string, arg2 byte, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestWalletConnectionReview", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestWalletConnectionReview indicates an expected call of RequestWalletConnectionReview.
func (mr *MockInteractorMockRecorder) RequestWalletConnectionReview(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestWalletConnectionReview", reflect.TypeOf((*MockInteractor)(nil).RequestWalletConnectionReview), arg0, arg1, arg2, arg3)
}

// RequestWalletSelection mocks base method.
func (m *MockInteractor) RequestWalletSelection(arg0 context.Context, arg1 string, arg2 byte, arg3 string, arg4 []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestWalletSelection", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestWalletSelection indicates an expected call of RequestWalletSelection.
func (mr *MockInteractorMockRecorder) RequestWalletSelection(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestWalletSelection", reflect.TypeOf((*MockInteractor)(nil).RequestWalletSelection), arg0, arg1, arg2, arg3, arg4)
}

// MockConnectionsManager is a mock of ConnectionsManager interface.
type MockConnectionsManager struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionsManagerMockRecorder
}

// MockConnectionsManagerMockRecorder is the mock recorder for MockConnectionsManager.
type MockConnectionsManagerMockRecorder struct {
	mock *MockConnectionsManager
}

// NewMockConnectionsManager creates a new mock instance.
func NewMockConnectionsManager(ctrl *gomock.Controller) *MockConnectionsManager {
	mock := &MockConnectionsManager{ctrl: ctrl}
	mock.recorder = &MockConnectionsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionsManager) EXPECT() *MockConnectionsManagerMockRecorder {
	return m.recorder
}

// EndSessionConnection mocks base method.
func (m *MockConnectionsManager) EndSessionConnection(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EndSessionConnection", arg0, arg1)
}

// EndSessionConnection indicates an expected call of EndSessionConnection.
func (mr *MockConnectionsManagerMockRecorder) EndSessionConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndSessionConnection", reflect.TypeOf((*MockConnectionsManager)(nil).EndSessionConnection), arg0, arg1)
}

// ListSessionConnections mocks base method.
func (m *MockConnectionsManager) ListSessionConnections() []api.Connection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSessionConnections")
	ret0, _ := ret[0].([]api.Connection)
	return ret0
}

// ListSessionConnections indicates an expected call of ListSessionConnections.
func (mr *MockConnectionsManagerMockRecorder) ListSessionConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSessionConnections", reflect.TypeOf((*MockConnectionsManager)(nil).ListSessionConnections))
}

// MockSpamHandler is a mock of SpamHandler interface.
type MockSpamHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSpamHandlerMockRecorder
}

// MockSpamHandlerMockRecorder is the mock recorder for MockSpamHandler.
type MockSpamHandlerMockRecorder struct {
	mock *MockSpamHandler
}

// NewMockSpamHandler creates a new mock instance.
func NewMockSpamHandler(ctrl *gomock.Controller) *MockSpamHandler {
	mock := &MockSpamHandler{ctrl: ctrl}
	mock.recorder = &MockSpamHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpamHandler) EXPECT() *MockSpamHandlerMockRecorder {
	return m.recorder
}

// CheckSubmission mocks base method.
func (m *MockSpamHandler) CheckSubmission(arg0 *v10.SubmitTransactionRequest, arg1 *types.SpamStatistics) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSubmission", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckSubmission indicates an expected call of CheckSubmission.
func (mr *MockSpamHandlerMockRecorder) CheckSubmission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSubmission", reflect.TypeOf((*MockSpamHandler)(nil).CheckSubmission), arg0, arg1)
}

// GenerateProofOfWork mocks base method.
func (m *MockSpamHandler) GenerateProofOfWork(arg0 string, arg1 *types.SpamStatistics) (*v1.ProofOfWork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateProofOfWork", arg0, arg1)
	ret0, _ := ret[0].(*v1.ProofOfWork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateProofOfWork indicates an expected call of GenerateProofOfWork.
func (mr *MockSpamHandlerMockRecorder) GenerateProofOfWork(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateProofOfWork", reflect.TypeOf((*MockSpamHandler)(nil).GenerateProofOfWork), arg0, arg1)
}
