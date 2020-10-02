// Code generated by MockGen. DO NOT EDIT.
// Source: usuario.go

// Package Mocks is a generated GoMock package.
package Mocks

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockIUsuario is a mock of IUsuario interface
type MockIUsuario struct {
	ctrl     *gomock.Controller
	recorder *MockIUsuarioMockRecorder
}

// MockIUsuarioMockRecorder is the mock recorder for MockIUsuario
type MockIUsuarioMockRecorder struct {
	mock *MockIUsuario
}

// NewMockIUsuario creates a new mock instance
func NewMockIUsuario(ctrl *gomock.Controller) *MockIUsuario {
	mock := &MockIUsuario{ctrl: ctrl}
	mock.recorder = &MockIUsuarioMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUsuario) EXPECT() *MockIUsuarioMockRecorder {
	return m.recorder
}

// GetUsers mocks base method
func (m *MockIUsuario) GetUsers(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUsers", c)
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockIUsuarioMockRecorder) GetUsers(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockIUsuario)(nil).GetUsers), c)
}

// PostUsers mocks base method
func (m *MockIUsuario) PostUsers(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PostUsers", c)
}

// PostUsers indicates an expected call of PostUsers
func (mr *MockIUsuarioMockRecorder) PostUsers(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUsers", reflect.TypeOf((*MockIUsuario)(nil).PostUsers), c)
}

// PutUsers mocks base method
func (m *MockIUsuario) PutUsers(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutUsers", c)
}

// PutUsers indicates an expected call of PutUsers
func (mr *MockIUsuarioMockRecorder) PutUsers(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUsers", reflect.TypeOf((*MockIUsuario)(nil).PutUsers), c)
}

// DeleteUsers mocks base method
func (m *MockIUsuario) DeleteUsers(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUsers", c)
}

// DeleteUsers indicates an expected call of DeleteUsers
func (mr *MockIUsuarioMockRecorder) DeleteUsers(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsers", reflect.TypeOf((*MockIUsuario)(nil).DeleteUsers), c)
}
