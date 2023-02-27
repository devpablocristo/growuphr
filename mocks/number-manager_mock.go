// Code generated by MockGen. DO NOT EDIT.
// Source: ./number-manager.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/devpablocristo/growuphr/number-manager/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockNumberManager is a mock of NumberManager interface.
type MockNumberManager struct {
	ctrl     *gomock.Controller
	recorder *MockNumberManagerMockRecorder
}

// MockNumberManagerMockRecorder is the mock recorder for MockNumberManager.
type MockNumberManagerMockRecorder struct {
	mock *MockNumberManager
}

// NewMockNumberManager creates a new mock instance.
func NewMockNumberManager(ctrl *gomock.Controller) *MockNumberManager {
	mock := &MockNumberManager{ctrl: ctrl}
	mock.recorder = &MockNumberManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNumberManager) EXPECT() *MockNumberManagerMockRecorder {
	return m.recorder
}

// AddReserveNumber mocks base method.
func (m *MockNumberManager) AddReserveNumber(arg0 context.Context, arg1 *domain.ReservedNumber) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddReserveNumber", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddReserveNumber indicates an expected call of AddReserveNumber.
func (mr *MockNumberManagerMockRecorder) AddReserveNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReserveNumber", reflect.TypeOf((*MockNumberManager)(nil).AddReserveNumber), arg0, arg1)
}

// ReservedNumbers mocks base method.
func (m *MockNumberManager) ReservedNumbers(arg0 context.Context) (map[string]*domain.ReservedNumber, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReservedNumbers", arg0)
	ret0, _ := ret[0].(map[string]*domain.ReservedNumber)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReservedNumbers indicates an expected call of ReservedNumbers.
func (mr *MockNumberManagerMockRecorder) ReservedNumbers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReservedNumbers", reflect.TypeOf((*MockNumberManager)(nil).ReservedNumbers), arg0)
}
