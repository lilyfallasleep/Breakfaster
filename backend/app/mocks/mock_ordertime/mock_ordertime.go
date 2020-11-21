// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/ordertime/interface.go

// Package mock_ordertime is a generated GoMock package.
package mock_ordertime

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockOrderTimer is a mock of OrderTimer interface
type MockOrderTimer struct {
	ctrl     *gomock.Controller
	recorder *MockOrderTimerMockRecorder
}

// MockOrderTimerMockRecorder is the mock recorder for MockOrderTimer
type MockOrderTimerMockRecorder struct {
	mock *MockOrderTimer
}

// NewMockOrderTimer creates a new mock instance
func NewMockOrderTimer(ctrl *gomock.Controller) *MockOrderTimer {
	mock := &MockOrderTimer{ctrl: ctrl}
	mock.recorder = &MockOrderTimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderTimer) EXPECT() *MockOrderTimerMockRecorder {
	return m.recorder
}

// GetNextWeekInterval mocks base method
func (m *MockOrderTimer) GetNextWeekInterval() (time.Time, time.Time) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextWeekInterval")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(time.Time)
	return ret0, ret1
}

// GetNextWeekInterval indicates an expected call of GetNextWeekInterval
func (mr *MockOrderTimerMockRecorder) GetNextWeekInterval() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextWeekInterval", reflect.TypeOf((*MockOrderTimer)(nil).GetNextWeekInterval))
}