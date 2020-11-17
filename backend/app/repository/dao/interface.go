package dao

import (
	"breakfaster/repository/model"
	"breakfaster/repository/schema"
	"time"
)

// EmployeeRepository is the interface for employee dao
type EmployeeRepository interface {
	GetEmpID(lineUID string) (string, error)
	GetLineUID(EmpID string) (string, error)
	UpsertEmployeeByIDs(employee *model.Employee) error
}

// FoodRepository is the interface for food dao
type FoodRepository interface {
	GetFoodAll(start, end time.Time) (*[]schema.SelectFood, error)
}

// OrderRepository is the interface for order dao
type OrderRepository interface {
	CreateOrders(orders *[]model.Order) error
	DeleteOrdersByLineUID(lineUID string, start, end time.Time) error
	GetOrdersByLineUID(lineUID string, start, end time.Time) (*[]schema.SelectOrder, error)
	GetOrderByEmpID(empID string, date time.Time) (*schema.SelectOrderWithEmployeeEmpID, error)
	GetOrderByAccessCardNumber(accessCardNumber string, date time.Time) (*schema.SelectOrderWithEmployeeEmpID, error)
	UpdateOrderStatus(empID string, date time.Time, pick bool) error
}
