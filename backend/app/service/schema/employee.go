package schema

// JSONEmployee is the service schema for returned json employee
type JSONEmployee struct {
	EmpID   string `json:"emp_id" example:"LW99999"`
	LineUID string `json:"line_uid" example:"U6664ceab1f4466b30827d936cee888e6"`
}
