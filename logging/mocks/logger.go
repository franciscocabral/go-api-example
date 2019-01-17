package mocks

import (
	"github.com/franciscocabral/go-api-example/logging"
	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

// Debug - mock for Debug
func (_m *MockLogger) Debug(message string, data ...logging.AdditionalData) {
	_m.Called(message, data)
}

// Info - mock for Info
func (_m *MockLogger) Info(message string, data ...logging.AdditionalData) {
	_m.Called(message, data)
}

// Warning - mock for Warning
func (_m *MockLogger) Warning(message string, data ...logging.AdditionalData) {
	_m.Called(message, data)
}

// Error - mock for Error
func (_m *MockLogger) Error(message string, data ...logging.AdditionalData) {
	_m.Called(message, data)
}

// Critical - mock for Critical
func (_m *MockLogger) Critical(message string, data ...logging.AdditionalData) {
	_m.Called(message, data)
}

// SetLevel - mock for SetLevel
func (_m *MockLogger) SetLevel(level logging.LogLevel) {
	_m.Called(level)
}
