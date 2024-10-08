// Code generated by mockery v2.45.0. DO NOT EDIT.

package readfile

import (
	model "github.com/bagusandrian/reconciliation-service/internals/model"
	mock "github.com/stretchr/testify/mock"
)

// MockReadFile is an autogenerated mock type for the ReadFile type
type MockReadFile struct {
	mock.Mock
}

// GetBankReconciliationCSV provides a mock function with given fields: req
func (_m *MockReadFile) GetBankReconciliationCSV(req model.ReconciliationRequest) (model.DataBank, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for GetBankReconciliationCSV")
	}

	var r0 model.DataBank
	var r1 error
	if rf, ok := ret.Get(0).(func(model.ReconciliationRequest) (model.DataBank, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(model.ReconciliationRequest) model.DataBank); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(model.DataBank)
	}

	if rf, ok := ret.Get(1).(func(model.ReconciliationRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSystemReconciliationCSV provides a mock function with given fields: req
func (_m *MockReadFile) GetSystemReconciliationCSV(req model.ReconciliationRequest) (model.DataSystem, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for GetSystemReconciliationCSV")
	}

	var r0 model.DataSystem
	var r1 error
	if rf, ok := ret.Get(0).(func(model.ReconciliationRequest) (model.DataSystem, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(model.ReconciliationRequest) model.DataSystem); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(model.DataSystem)
	}

	if rf, ok := ret.Get(1).(func(model.ReconciliationRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockReadFile creates a new instance of MockReadFile. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReadFile(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReadFile {
	mock := &MockReadFile{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
