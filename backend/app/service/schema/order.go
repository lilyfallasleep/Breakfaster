package schema

// Order is the service schema for a single order
type Order struct {
	FoodID int    `json:"food_id" binding:"required" example:"1"`
	Date   string `json:"date" binding:"required,len=10" example:"2020-09-01"`
}

// AllOrders is the service schema for all orders of an employee
type AllOrders struct {
	EmpID string  `json:"emp_id" binding:"required,len=7" example:"LW99999"`
	Foods []Order `json:"foods" binding:"required,max=5"`
}

// JSONOrder is the service schema for returned json order
type JSONOrder struct {
	FoodName string `json:"food_name" example:"burger"`
	EmpID    string `json:"emp_id" example:"LW99999"`
	Date     string `json:"date" example:"2020-09-01"`
	Pick     bool   `json:"pick" example:"false"`
}
