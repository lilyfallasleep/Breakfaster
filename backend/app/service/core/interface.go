package core

import (
	ss "breakfaster/service/schema"
	"time"
)

// EmployeeService is the interface for employee service
type EmployeeService interface {
	GetEmployeeByLineUID(lineUID string) (*ss.JSONEmployee, error)
	GetEmployeeByEmpID(empID string) (*ss.JSONEmployee, error)
	UpsertEmployeeByIDs(empID, lineUID string) error
}

// FoodService is the interface for food service
type FoodService interface {
	GetFoodAll(startDate, endDate string) (*ss.NestedFood, error)
}

// OrderService is the interface for order service
type OrderService interface {
	SendOrderConfirmMessage(empID string, start, end time.Time) error
	CreateOrders(rawOrders *ss.AllOrders) error
	GetOrderByEmpID(empID, rawDate string) (*ss.JSONOrder, error)
	GetOrderByAccessCardNumber(accessCardNumber, rawDate string) (*ss.JSONOrder, error)
	SetPick(empID string, rawDate string) error
}
