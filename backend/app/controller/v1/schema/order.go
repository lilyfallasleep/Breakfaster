package schema

// FormGetOrder is the query string schema for getting an order by card number or employee ID
type FormGetOrder struct {
	Type    string `form:"type" binding:"required,oneof=eid card"`
	Payload string `form:"payload" binding:"required,max=10"`
	Date    string `form:"date" binding:"required,len=10"`
}

// PutPickOrder is the put schema for updating order picking status
type PutPickOrder struct {
	EmpID string `json:"emp_id" binding:"required,len=7" example:"LW99999"`
	Date  string `json:"date" binding:"required,len=10" example:"2020-09-01"`
}
