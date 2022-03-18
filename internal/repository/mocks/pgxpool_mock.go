// Code generated by MockGen. DO NOT EDIT.
// Source: base_repo_interface.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"reflect"
)

// MockPgxPoolInterface is a mock of PgxPoolInterface interface.
type MockPgxPoolInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPgxPoolInterfaceMockRecorder
}

// MockPgxPoolInterfaceMockRecorder is the mock recorder for MockPgxPoolInterface.
type MockPgxPoolInterfaceMockRecorder struct {
	mock *MockPgxPoolInterface
}

// NewMockPgxPoolInterface creates a new mock instance.
func NewMockPgxPoolInterface(ctrl *gomock.Controller) *MockPgxPoolInterface {
	mock := &MockPgxPoolInterface{ctrl: ctrl}
	mock.recorder = &MockPgxPoolInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPgxPoolInterface) EXPECT() *MockPgxPoolInterfaceMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockPgxPoolInterface) Begin(ctx context.Context) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", ctx)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockPgxPoolInterfaceMockRecorder) Begin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockPgxPoolInterface)(nil).Begin), ctx)
}

// BeginTxFunc mocks base method.
func (m *MockPgxPoolInterface) BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTxFunc", ctx, txOptions, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginTxFunc indicates an expected call of BeginTxFunc.
func (mr *MockPgxPoolInterfaceMockRecorder) BeginTxFunc(ctx, txOptions, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTxFunc", reflect.TypeOf((*MockPgxPoolInterface)(nil).BeginTxFunc), ctx, txOptions, f)
}

// Exec mocks base method.
func (m *MockPgxPoolInterface) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range arguments {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockPgxPoolInterfaceMockRecorder) Exec(ctx, sql interface{}, arguments ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, arguments...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockPgxPoolInterface)(nil).Exec), varargs...)
}

// Query mocks base method.
func (m *MockPgxPoolInterface) Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range optionsAndArgs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockPgxPoolInterfaceMockRecorder) Query(ctx, sql interface{}, optionsAndArgs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, optionsAndArgs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockPgxPoolInterface)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockPgxPoolInterface) QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range optionsAndArgs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(pgx.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockPgxPoolInterfaceMockRecorder) QueryRow(ctx, sql interface{}, optionsAndArgs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, optionsAndArgs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockPgxPoolInterface)(nil).QueryRow), varargs...)
}

// SendBatch mocks base method.
func (m *MockPgxPoolInterface) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendBatch", ctx, b)
	ret0, _ := ret[0].(pgx.BatchResults)
	return ret0
}

// SendBatch indicates an expected call of SendBatch.
func (mr *MockPgxPoolInterfaceMockRecorder) SendBatch(ctx, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBatch", reflect.TypeOf((*MockPgxPoolInterface)(nil).SendBatch), ctx, b)
}