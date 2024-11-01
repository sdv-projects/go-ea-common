// Code generated by MockGen. DO NOT EDIT.
// Source: command.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCommandHandler is a mock of CommandHandler interface.
type MockCommandHandler[C any] struct {
	ctrl     *gomock.Controller
	recorder *MockCommandHandlerMockRecorder[C]
}

// MockCommandHandlerMockRecorder is the mock recorder for MockCommandHandler.
type MockCommandHandlerMockRecorder[C any] struct {
	mock *MockCommandHandler[C]
}

// NewMockCommandHandler creates a new mock instance.
func NewMockCommandHandler[C any](ctrl *gomock.Controller) *MockCommandHandler[C] {
	mock := &MockCommandHandler[C]{ctrl: ctrl}
	mock.recorder = &MockCommandHandlerMockRecorder[C]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandHandler[C]) EXPECT() *MockCommandHandlerMockRecorder[C] {
	return m.recorder
}

// Execute mocks base method.
func (m *MockCommandHandler[C]) Execute(ctx context.Context, command C) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, command)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockCommandHandlerMockRecorder[C]) Execute(ctx, command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCommandHandler[C])(nil).Execute), ctx, command)
}

// MockCommandHandlerWithResult is a mock of CommandHandlerWithResult interface.
type MockCommandHandlerWithResult[C any, R any] struct {
	ctrl     *gomock.Controller
	recorder *MockCommandHandlerWithResultMockRecorder[C, R]
}

// MockCommandHandlerWithResultMockRecorder is the mock recorder for MockCommandHandlerWithResult.
type MockCommandHandlerWithResultMockRecorder[C any, R any] struct {
	mock *MockCommandHandlerWithResult[C, R]
}

// NewMockCommandHandlerWithResult creates a new mock instance.
func NewMockCommandHandlerWithResult[C any, R any](ctrl *gomock.Controller) *MockCommandHandlerWithResult[C, R] {
	mock := &MockCommandHandlerWithResult[C, R]{ctrl: ctrl}
	mock.recorder = &MockCommandHandlerWithResultMockRecorder[C, R]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandHandlerWithResult[C, R]) EXPECT() *MockCommandHandlerWithResultMockRecorder[C, R] {
	return m.recorder
}

// Execute mocks base method.
func (m *MockCommandHandlerWithResult[C, R]) Execute(ctx context.Context, command C) (R, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, command)
	ret0, _ := ret[0].(R)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockCommandHandlerWithResultMockRecorder[C, R]) Execute(ctx, command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCommandHandlerWithResult[C, R])(nil).Execute), ctx, command)
}
