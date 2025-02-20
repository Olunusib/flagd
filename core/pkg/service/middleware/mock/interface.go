// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/middleware/interface.go

// Package middlewaremock is a generated GoMock package.
package middlewaremock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIMiddleware is a mock of IMiddleware interface.
type MockIMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockIMiddlewareMockRecorder
}

// MockIMiddlewareMockRecorder is the mock recorder for MockIMiddleware.
type MockIMiddlewareMockRecorder struct {
	mock *MockIMiddleware
}

// NewMockIMiddleware creates a new mock instance.
func NewMockIMiddleware(ctrl *gomock.Controller) *MockIMiddleware {
	mock := &MockIMiddleware{ctrl: ctrl}
	mock.recorder = &MockIMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMiddleware) EXPECT() *MockIMiddlewareMockRecorder {
	return m.recorder
}

// Handler mocks base method.
func (m *MockIMiddleware) Handler(handler http.Handler) http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handler", handler)
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// Handler indicates an expected call of Handler.
func (mr *MockIMiddlewareMockRecorder) Handler(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockIMiddleware)(nil).Handler), handler)
}
