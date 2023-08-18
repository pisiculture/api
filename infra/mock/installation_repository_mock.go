// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/installation_repository_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/pisiculture/internal/core/domain"
)

// MockInstallationRepositoryInterface is a mock of InstallationRepositoryInterface interface.
type MockInstallationRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInstallationRepositoryInterfaceMockRecorder
}

// MockInstallationRepositoryInterfaceMockRecorder is the mock recorder for MockInstallationRepositoryInterface.
type MockInstallationRepositoryInterfaceMockRecorder struct {
	mock *MockInstallationRepositoryInterface
}

// NewMockInstallationRepositoryInterface creates a new mock instance.
func NewMockInstallationRepositoryInterface(ctrl *gomock.Controller) *MockInstallationRepositoryInterface {
	mock := &MockInstallationRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockInstallationRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstallationRepositoryInterface) EXPECT() *MockInstallationRepositoryInterfaceMockRecorder {
	return m.recorder
}

// DeleteById mocks base method.
func (m *MockInstallationRepositoryInterface) DeleteById(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockInstallationRepositoryInterfaceMockRecorder) DeleteById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockInstallationRepositoryInterface)(nil).DeleteById), id)
}

// FindByID mocks base method.
func (m *MockInstallationRepositoryInterface) FindByID(id int) (*domain.Installation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*domain.Installation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockInstallationRepositoryInterfaceMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockInstallationRepositoryInterface)(nil).FindByID), id)
}

// Save mocks base method.
func (m *MockInstallationRepositoryInterface) Save(installation domain.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", installation)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockInstallationRepositoryInterfaceMockRecorder) Save(installation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockInstallationRepositoryInterface)(nil).Save), installation)
}
