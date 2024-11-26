// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	automation "github.com/goplugin/plugin-common/pkg/types/automation"

	mock "github.com/stretchr/testify/mock"
)

// MockRunnable is an autogenerated mock type for the Runnable type
type MockRunnable struct {
	mock.Mock
}

// CheckUpkeeps provides a mock function with given fields: _a0, _a1
func (_m *MockRunnable) CheckUpkeeps(_a0 context.Context, _a1 ...automation.UpkeepPayload) ([]automation.CheckResult, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []automation.CheckResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...automation.UpkeepPayload) ([]automation.CheckResult, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...automation.UpkeepPayload) []automation.CheckResult); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]automation.CheckResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...automation.UpkeepPayload) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRunnable interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRunnable creates a new instance of MockRunnable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRunnable(t mockConstructorTestingTNewMockRunnable) *MockRunnable {
	mock := &MockRunnable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}