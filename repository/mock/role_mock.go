// Code generated by MockGen. DO NOT EDIT.
// Source: repository/role.go
//
// Generated by this command:
//
//	mockgen -source=repository/role.go -destination=repository/mock/role_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	rolev1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockRoleRepository is a mock of RoleRepository interface.
type MockRoleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoleRepositoryMockRecorder
}

// MockRoleRepositoryMockRecorder is the mock recorder for MockRoleRepository.
type MockRoleRepositoryMockRecorder struct {
	mock *MockRoleRepository
}

// NewMockRoleRepository creates a new mock instance.
func NewMockRoleRepository(ctrl *gomock.Controller) *MockRoleRepository {
	mock := &MockRoleRepository{ctrl: ctrl}
	mock.recorder = &MockRoleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleRepository) EXPECT() *MockRoleRepositoryMockRecorder {
	return m.recorder
}

// AddRole mocks base method.
func (m *MockRoleRepository) AddRole(arg0 context.Context, arg1 *rolev1.Role) (*rolev1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRole", arg0, arg1)
	ret0, _ := ret[0].(*rolev1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRole indicates an expected call of AddRole.
func (mr *MockRoleRepositoryMockRecorder) AddRole(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRole", reflect.TypeOf((*MockRoleRepository)(nil).AddRole), arg0, arg1)
}

// EditRole mocks base method.
func (m *MockRoleRepository) EditRole(arg0 context.Context, arg1 *rolev1.Role) (*rolev1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditRole", arg0, arg1)
	ret0, _ := ret[0].(*rolev1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditRole indicates an expected call of EditRole.
func (mr *MockRoleRepositoryMockRecorder) EditRole(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditRole", reflect.TypeOf((*MockRoleRepository)(nil).EditRole), arg0, arg1)
}

// GetRole mocks base method.
func (m *MockRoleRepository) GetRole(arg0 context.Context, arg1 *rolev1.Role) (*rolev1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0, arg1)
	ret0, _ := ret[0].(*rolev1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockRoleRepositoryMockRecorder) GetRole(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockRoleRepository)(nil).GetRole), arg0, arg1)
}

// GetRoleList mocks base method.
func (m *MockRoleRepository) GetRoleList(arg0 context.Context, arg1 *rolev1.Role) ([]*rolev1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleList", arg0, arg1)
	ret0, _ := ret[0].([]*rolev1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleList indicates an expected call of GetRoleList.
func (mr *MockRoleRepositoryMockRecorder) GetRoleList(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleList", reflect.TypeOf((*MockRoleRepository)(nil).GetRoleList), arg0, arg1)
}

// RemoveRole mocks base method.
func (m *MockRoleRepository) RemoveRole(arg0 context.Context, arg1 *rolev1.Role) (*rolev1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRole", arg0, arg1)
	ret0, _ := ret[0].(*rolev1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveRole indicates an expected call of RemoveRole.
func (mr *MockRoleRepositoryMockRecorder) RemoveRole(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRole", reflect.TypeOf((*MockRoleRepository)(nil).RemoveRole), arg0, arg1)
}
