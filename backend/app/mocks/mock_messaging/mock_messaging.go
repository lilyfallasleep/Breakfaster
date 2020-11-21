// Code generated by MockGen. DO NOT EDIT.
// Source: messaging/interface.go

// Package mock_messaging is a generated GoMock package.
package mock_messaging

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMessageController is a mock of MessageController interface
type MockMessageController struct {
	ctrl     *gomock.Controller
	recorder *MockMessageControllerMockRecorder
}

// MockMessageControllerMockRecorder is the mock recorder for MockMessageController
type MockMessageControllerMockRecorder struct {
	mock *MockMessageController
}

// NewMockMessageController creates a new mock instance
func NewMockMessageController(ctrl *gomock.Controller) *MockMessageController {
	mock := &MockMessageController{ctrl: ctrl}
	mock.recorder = &MockMessageControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageController) EXPECT() *MockMessageControllerMockRecorder {
	return m.recorder
}

// BroadCastMenu mocks base method
func (m *MockMessageController) BroadCastMenu() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadCastMenu")
	ret0, _ := ret[0].(error)
	return ret0
}

// BroadCastMenu indicates an expected call of BroadCastMenu
func (mr *MockMessageControllerMockRecorder) BroadCastMenu() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadCastMenu", reflect.TypeOf((*MockMessageController)(nil).BroadCastMenu))
}