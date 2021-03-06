// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/documentdb (interfaces: DatabaseAccountsClient)

// Package mock_documentdb is a generated GoMock package.
package mock_documentdb

import (
	context "context"
	reflect "reflect"

	documentdb "github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2019-08-01/documentdb"
	gomock "github.com/golang/mock/gomock"
)

// MockDatabaseAccountsClient is a mock of DatabaseAccountsClient interface
type MockDatabaseAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseAccountsClientMockRecorder
}

// MockDatabaseAccountsClientMockRecorder is the mock recorder for MockDatabaseAccountsClient
type MockDatabaseAccountsClientMockRecorder struct {
	mock *MockDatabaseAccountsClient
}

// NewMockDatabaseAccountsClient creates a new mock instance
func NewMockDatabaseAccountsClient(ctrl *gomock.Controller) *MockDatabaseAccountsClient {
	mock := &MockDatabaseAccountsClient{ctrl: ctrl}
	mock.recorder = &MockDatabaseAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseAccountsClient) EXPECT() *MockDatabaseAccountsClientMockRecorder {
	return m.recorder
}

// ListByResourceGroup mocks base method
func (m *MockDatabaseAccountsClient) ListByResourceGroup(arg0 context.Context, arg1 string) (documentdb.DatabaseAccountsListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByResourceGroup", arg0, arg1)
	ret0, _ := ret[0].(documentdb.DatabaseAccountsListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByResourceGroup indicates an expected call of ListByResourceGroup
func (mr *MockDatabaseAccountsClientMockRecorder) ListByResourceGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByResourceGroup", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).ListByResourceGroup), arg0, arg1)
}

// ListKeys mocks base method
func (m *MockDatabaseAccountsClient) ListKeys(arg0 context.Context, arg1, arg2 string) (documentdb.DatabaseAccountListKeysResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].(documentdb.DatabaseAccountListKeysResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys
func (mr *MockDatabaseAccountsClientMockRecorder) ListKeys(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).ListKeys), arg0, arg1, arg2)
}
