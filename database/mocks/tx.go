// Code generated by MockGen. DO NOT EDIT.
// Source: tx.go

// Package database_mocks is a generated GoMock package.
package database_mocks

import (
	context "context"
	reflect "reflect"

	database "github.com/sdv-projects/go-ea-common/database"
	gomock "go.uber.org/mock/gomock"
)

// MockTx is a mock of Tx interface.
type MockTx struct {
	ctrl     *gomock.Controller
	recorder *MockTxMockRecorder
}

// MockTxMockRecorder is the mock recorder for MockTx.
type MockTxMockRecorder struct {
	mock *MockTx
}

// NewMockTx creates a new mock instance.
func NewMockTx(ctrl *gomock.Controller) *MockTx {
	mock := &MockTx{ctrl: ctrl}
	mock.recorder = &MockTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTx) EXPECT() *MockTxMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockTx) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTxMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTx)(nil).Commit))
}

// Rollback mocks base method.
func (m *MockTx) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTxMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTx)(nil).Rollback))
}

// MockTxFactory is a mock of TxFactory interface.
type MockTxFactory struct {
	ctrl     *gomock.Controller
	recorder *MockTxFactoryMockRecorder
}

// MockTxFactoryMockRecorder is the mock recorder for MockTxFactory.
type MockTxFactoryMockRecorder struct {
	mock *MockTxFactory
}

// NewMockTxFactory creates a new mock instance.
func NewMockTxFactory(ctrl *gomock.Controller) *MockTxFactory {
	mock := &MockTxFactory{ctrl: ctrl}
	mock.recorder = &MockTxFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxFactory) EXPECT() *MockTxFactoryMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockTxFactory) BeginTx(ctx context.Context) (database.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx)
	ret0, _ := ret[0].(database.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockTxFactoryMockRecorder) BeginTx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockTxFactory)(nil).BeginTx), ctx)
}
