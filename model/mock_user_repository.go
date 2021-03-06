// Code generated by MockGen. DO NOT EDIT.
// Source: go-community/model (interfaces: UserRepository)

// Package model is a generated GoMock package.
package model

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	bun "github.com/uptrace/bun"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CountEmail mocks base method.
func (m *MockUserRepository) CountEmail(arg0 context.Context, arg1 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountEmail", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// CountEmail indicates an expected call of CountEmail.
func (mr *MockUserRepositoryMockRecorder) CountEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountEmail", reflect.TypeOf((*MockUserRepository)(nil).CountEmail), arg0, arg1)
}

// Insert mocks base method.
func (m *MockUserRepository) Insert(arg0 bun.Tx, arg1 context.Context, arg2, arg3, arg4 string) (*int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockUserRepositoryMockRecorder) Insert(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserRepository)(nil).Insert), arg0, arg1, arg2, arg3, arg4)
}
