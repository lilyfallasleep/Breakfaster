// Code generated by MockGen. DO NOT EDIT.
// Source: service/core/interface.go

// Package mock_core is a generated GoMock package.
package mock_core

import (
	schema "breakfaster/service/schema"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockEmployeeService is a mock of EmployeeService interface
type MockEmployeeService struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeServiceMockRecorder
}

// MockEmployeeServiceMockRecorder is the mock recorder for MockEmployeeService
type MockEmployeeServiceMockRecorder struct {
	mock *MockEmployeeService
}

// NewMockEmployeeService creates a new mock instance
func NewMockEmployeeService(ctrl *gomock.Controller) *MockEmployeeService {
	mock := &MockEmployeeService{ctrl: ctrl}
	mock.recorder = &MockEmployeeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEmployeeService) EXPECT() *MockEmployeeServiceMockRecorder {
	return m.recorder
}

// GetEmployeeByLineUID mocks base method
func (m *MockEmployeeService) GetEmployeeByLineUID(lineUID string) (*schema.JSONEmployee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeByLineUID", lineUID)
	ret0, _ := ret[0].(*schema.JSONEmployee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeByLineUID indicates an expected call of GetEmployeeByLineUID
func (mr *MockEmployeeServiceMockRecorder) GetEmployeeByLineUID(lineUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeByLineUID", reflect.TypeOf((*MockEmployeeService)(nil).GetEmployeeByLineUID), lineUID)
}

// GetEmployeeByEmpID mocks base method
func (m *MockEmployeeService) GetEmployeeByEmpID(empID string) (*schema.JSONEmployee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeByEmpID", empID)
	ret0, _ := ret[0].(*schema.JSONEmployee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeByEmpID indicates an expected call of GetEmployeeByEmpID
func (mr *MockEmployeeServiceMockRecorder) GetEmployeeByEmpID(empID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeByEmpID", reflect.TypeOf((*MockEmployeeService)(nil).GetEmployeeByEmpID), empID)
}

// UpsertEmployeeByIDs mocks base method
func (m *MockEmployeeService) UpsertEmployeeByIDs(empID, lineUID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertEmployeeByIDs", empID, lineUID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertEmployeeByIDs indicates an expected call of UpsertEmployeeByIDs
func (mr *MockEmployeeServiceMockRecorder) UpsertEmployeeByIDs(empID, lineUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertEmployeeByIDs", reflect.TypeOf((*MockEmployeeService)(nil).UpsertEmployeeByIDs), empID, lineUID)
}

// MockFoodService is a mock of FoodService interface
type MockFoodService struct {
	ctrl     *gomock.Controller
	recorder *MockFoodServiceMockRecorder
}

// MockFoodServiceMockRecorder is the mock recorder for MockFoodService
type MockFoodServiceMockRecorder struct {
	mock *MockFoodService
}

// NewMockFoodService creates a new mock instance
func NewMockFoodService(ctrl *gomock.Controller) *MockFoodService {
	mock := &MockFoodService{ctrl: ctrl}
	mock.recorder = &MockFoodServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFoodService) EXPECT() *MockFoodServiceMockRecorder {
	return m.recorder
}

// GetFoodAll mocks base method
func (m *MockFoodService) GetFoodAll(startDate, endDate string) (*schema.NestedFood, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFoodAll", startDate, endDate)
	ret0, _ := ret[0].(*schema.NestedFood)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFoodAll indicates an expected call of GetFoodAll
func (mr *MockFoodServiceMockRecorder) GetFoodAll(startDate, endDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFoodAll", reflect.TypeOf((*MockFoodService)(nil).GetFoodAll), startDate, endDate)
}

// MockOrderService is a mock of OrderService interface
type MockOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceMockRecorder
}

// MockOrderServiceMockRecorder is the mock recorder for MockOrderService
type MockOrderServiceMockRecorder struct {
	mock *MockOrderService
}

// NewMockOrderService creates a new mock instance
func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
	mock := &MockOrderService{ctrl: ctrl}
	mock.recorder = &MockOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
	return m.recorder
}

// SendOrderConfirmMessage mocks base method
func (m *MockOrderService) SendOrderConfirmMessage(empID string, start, end time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendOrderConfirmMessage", empID, start, end)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendOrderConfirmMessage indicates an expected call of SendOrderConfirmMessage
func (mr *MockOrderServiceMockRecorder) SendOrderConfirmMessage(empID, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendOrderConfirmMessage", reflect.TypeOf((*MockOrderService)(nil).SendOrderConfirmMessage), empID, start, end)
}

// CreateOrders mocks base method
func (m *MockOrderService) CreateOrders(rawOrders *schema.AllOrders) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrders", rawOrders)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrders indicates an expected call of CreateOrders
func (mr *MockOrderServiceMockRecorder) CreateOrders(rawOrders interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrders", reflect.TypeOf((*MockOrderService)(nil).CreateOrders), rawOrders)
}

// GetOrderByEmpID mocks base method
func (m *MockOrderService) GetOrderByEmpID(empID, rawDate string) (*schema.JSONOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByEmpID", empID, rawDate)
	ret0, _ := ret[0].(*schema.JSONOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByEmpID indicates an expected call of GetOrderByEmpID
func (mr *MockOrderServiceMockRecorder) GetOrderByEmpID(empID, rawDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByEmpID", reflect.TypeOf((*MockOrderService)(nil).GetOrderByEmpID), empID, rawDate)
}

// GetOrderByAccessCardNumber mocks base method
func (m *MockOrderService) GetOrderByAccessCardNumber(accessCardNumber, rawDate string) (*schema.JSONOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByAccessCardNumber", accessCardNumber, rawDate)
	ret0, _ := ret[0].(*schema.JSONOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByAccessCardNumber indicates an expected call of GetOrderByAccessCardNumber
func (mr *MockOrderServiceMockRecorder) GetOrderByAccessCardNumber(accessCardNumber, rawDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByAccessCardNumber", reflect.TypeOf((*MockOrderService)(nil).GetOrderByAccessCardNumber), accessCardNumber, rawDate)
}

// SetPick mocks base method
func (m *MockOrderService) SetPick(empID, rawDate string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPick", empID, rawDate)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPick indicates an expected call of SetPick
func (mr *MockOrderServiceMockRecorder) SetPick(empID, rawDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPick", reflect.TypeOf((*MockOrderService)(nil).SetPick), empID, rawDate)
}
