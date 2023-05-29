// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package api is a generated GoMock package.
package api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRemoteCalls is a mock of RemoteCalls interface.
type MockRemoteCalls struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteCallsMockRecorder
}

// MockRemoteCallsMockRecorder is the mock recorder for MockRemoteCalls.
type MockRemoteCallsMockRecorder struct {
	mock *MockRemoteCalls
}

// NewMockRemoteCalls creates a new mock instance.
func NewMockRemoteCalls(ctrl *gomock.Controller) *MockRemoteCalls {
	mock := &MockRemoteCalls{ctrl: ctrl}
	mock.recorder = &MockRemoteCallsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteCalls) EXPECT() *MockRemoteCallsMockRecorder {
	return m.recorder
}

// Auth mocks base method.
func (m *MockRemoteCalls) Auth(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Auth indicates an expected call of Auth.
func (mr *MockRemoteCallsMockRecorder) Auth(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockRemoteCalls)(nil).Auth), key)
}

// RunInSandboxMode mocks base method.
func (m *MockRemoteCalls) RunInSandboxMode() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RunInSandboxMode")
}

// RunInSandboxMode indicates an expected call of RunInSandboxMode.
func (mr *MockRemoteCallsMockRecorder) RunInSandboxMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInSandboxMode", reflect.TypeOf((*MockRemoteCalls)(nil).RunInSandboxMode))
}
