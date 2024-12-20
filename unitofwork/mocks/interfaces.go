// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	database "github.com/sdv-projects/go-ea-common/database"
	unitofwork "github.com/sdv-projects/go-ea-common/unitofwork"
	gomock "go.uber.org/mock/gomock"
)

// MockUnitOfWork is a mock of UnitOfWork interface.
type MockUnitOfWork struct {
	ctrl     *gomock.Controller
	recorder *MockUnitOfWorkMockRecorder
}

// MockUnitOfWorkMockRecorder is the mock recorder for MockUnitOfWork.
type MockUnitOfWorkMockRecorder struct {
	mock *MockUnitOfWork
}

// NewMockUnitOfWork creates a new mock instance.
func NewMockUnitOfWork(ctrl *gomock.Controller) *MockUnitOfWork {
	mock := &MockUnitOfWork{ctrl: ctrl}
	mock.recorder = &MockUnitOfWorkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnitOfWork) EXPECT() *MockUnitOfWorkMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockUnitOfWork) Commit(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockUnitOfWorkMockRecorder) Commit(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockUnitOfWork)(nil).Commit), ctx)
}

// RegisterDeleted mocks base method.
func (m *MockUnitOfWork) RegisterDeleted(entity any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterDeleted", entity)
}

// RegisterDeleted indicates an expected call of RegisterDeleted.
func (mr *MockUnitOfWorkMockRecorder) RegisterDeleted(entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterDeleted", reflect.TypeOf((*MockUnitOfWork)(nil).RegisterDeleted), entity)
}

// RegisterDirty mocks base method.
func (m *MockUnitOfWork) RegisterDirty(entity any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterDirty", entity)
}

// RegisterDirty indicates an expected call of RegisterDirty.
func (mr *MockUnitOfWorkMockRecorder) RegisterDirty(entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterDirty", reflect.TypeOf((*MockUnitOfWork)(nil).RegisterDirty), entity)
}

// RegisterNew mocks base method.
func (m *MockUnitOfWork) RegisterNew(entity any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterNew", entity)
}

// RegisterNew indicates an expected call of RegisterNew.
func (mr *MockUnitOfWorkMockRecorder) RegisterNew(entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNew", reflect.TypeOf((*MockUnitOfWork)(nil).RegisterNew), entity)
}

// MockUoWRepositoryFactory is a mock of UoWRepositoryFactory interface.
type MockUoWRepositoryFactory struct {
	ctrl     *gomock.Controller
	recorder *MockUoWRepositoryFactoryMockRecorder
}

// MockUoWRepositoryFactoryMockRecorder is the mock recorder for MockUoWRepositoryFactory.
type MockUoWRepositoryFactoryMockRecorder struct {
	mock *MockUoWRepositoryFactory
}

// NewMockUoWRepositoryFactory creates a new mock instance.
func NewMockUoWRepositoryFactory(ctrl *gomock.Controller) *MockUoWRepositoryFactory {
	mock := &MockUoWRepositoryFactory{ctrl: ctrl}
	mock.recorder = &MockUoWRepositoryFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUoWRepositoryFactory) EXPECT() *MockUoWRepositoryFactoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUoWRepositoryFactory) Get(ctx context.Context, et reflect.Type) (unitofwork.UoWRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, et)
	ret0, _ := ret[0].(unitofwork.UoWRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUoWRepositoryFactoryMockRecorder) Get(ctx, et interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUoWRepositoryFactory)(nil).Get), ctx, et)
}

// MockUoWRepository is a mock of UoWRepository interface.
type MockUoWRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUoWRepositoryMockRecorder
}

// MockUoWRepositoryMockRecorder is the mock recorder for MockUoWRepository.
type MockUoWRepositoryMockRecorder struct {
	mock *MockUoWRepository
}

// NewMockUoWRepository creates a new mock instance.
func NewMockUoWRepository(ctrl *gomock.Controller) *MockUoWRepository {
	mock := &MockUoWRepository{ctrl: ctrl}
	mock.recorder = &MockUoWRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUoWRepository) EXPECT() *MockUoWRepositoryMockRecorder {
	return m.recorder
}

// TxDelete mocks base method.
func (m *MockUoWRepository) TxDelete(ctx context.Context, tx database.Tx, entity ...any) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx}
	for _, a := range entity {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TxDelete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// TxDelete indicates an expected call of TxDelete.
func (mr *MockUoWRepositoryMockRecorder) TxDelete(ctx, tx interface{}, entity ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx}, entity...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxDelete", reflect.TypeOf((*MockUoWRepository)(nil).TxDelete), varargs...)
}

// TxInsert mocks base method.
func (m *MockUoWRepository) TxInsert(ctx context.Context, tx database.Tx, entity ...any) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx}
	for _, a := range entity {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TxInsert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// TxInsert indicates an expected call of TxInsert.
func (mr *MockUoWRepositoryMockRecorder) TxInsert(ctx, tx interface{}, entity ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx}, entity...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxInsert", reflect.TypeOf((*MockUoWRepository)(nil).TxInsert), varargs...)
}

// TxUpdate mocks base method.
func (m *MockUoWRepository) TxUpdate(ctx context.Context, tx database.Tx, entity ...any) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx}
	for _, a := range entity {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TxUpdate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// TxUpdate indicates an expected call of TxUpdate.
func (mr *MockUoWRepositoryMockRecorder) TxUpdate(ctx, tx interface{}, entity ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx}, entity...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxUpdate", reflect.TypeOf((*MockUoWRepository)(nil).TxUpdate), varargs...)
}
