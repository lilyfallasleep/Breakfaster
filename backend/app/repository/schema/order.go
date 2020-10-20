package schema

import "time"

// SelectOrder is the order schema for querying orders
type SelectOrder struct {
	FoodName string
	Date     time.Time
}

// SelectOrderWithEmployeeEmpID is the order schema for querying a daily order
type SelectOrderWithEmployeeEmpID struct {
	FoodName      string
	EmployeeEmpID string
	Date          time.Time
	Pick          bool
}
