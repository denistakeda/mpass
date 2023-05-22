// Code generated by MockGen. DO NOT EDIT.
// Source: internal/auth_service/auth_service.go

// Package auth_service_mock is a generated GoMock package.
package auth_service_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockuserStore is a mock of userStore interface.
type MockuserStore struct {
	ctrl     *gomock.Controller
	recorder *MockuserStoreMockRecorder
}

// MockuserStoreMockRecorder is the mock recorder for MockuserStore.
type MockuserStoreMockRecorder struct {
	mock *MockuserStore
}

// NewMockuserStore creates a new mock instance.
func NewMockuserStore(ctrl *gomock.Controller) *MockuserStore {
	mock := &MockuserStore{ctrl: ctrl}
	mock.recorder = &MockuserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserStore) EXPECT() *MockuserStoreMockRecorder {
	return m.recorder
}

// AddNewUser mocks base method.
func (m *MockuserStore) AddNewUser(ctx context.Context, login, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewUser", ctx, login, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewUser indicates an expected call of AddNewUser.
func (mr *MockuserStoreMockRecorder) AddNewUser(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewUser", reflect.TypeOf((*MockuserStore)(nil).AddNewUser), ctx, login, password)
}
