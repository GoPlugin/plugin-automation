// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	commontypes "github.com/goplugin/plugin-libocr/commontypes"
	mock "github.com/stretchr/testify/mock"
)

// MockLogger is an autogenerated mock type for the Logger type
type MockLogger struct {
	mock.Mock
}

// Critical provides a mock function with given fields: msg, fields
func (_m *MockLogger) Critical(msg string, fields commontypes.LogFields) {
	_m.Called(msg, fields)
}

// Debug provides a mock function with given fields: msg, fields
func (_m *MockLogger) Debug(msg string, fields commontypes.LogFields) {
	_m.Called(msg, fields)
}

// Error provides a mock function with given fields: msg, fields
func (_m *MockLogger) Error(msg string, fields commontypes.LogFields) {
	_m.Called(msg, fields)
}

// Info provides a mock function with given fields: msg, fields
func (_m *MockLogger) Info(msg string, fields commontypes.LogFields) {
	_m.Called(msg, fields)
}

// Trace provides a mock function with given fields: msg, fields
func (_m *MockLogger) Trace(msg string, fields commontypes.LogFields) {
	_m.Called(msg, fields)
}

// Warn provides a mock function with given fields: msg, fields
func (_m *MockLogger) Warn(msg string, fields commontypes.LogFields) {
	_m.Called(msg, fields)
}

type mockConstructorTestingTNewMockLogger interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockLogger creates a new instance of MockLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockLogger(t mockConstructorTestingTNewMockLogger) *MockLogger {
	mock := &MockLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
