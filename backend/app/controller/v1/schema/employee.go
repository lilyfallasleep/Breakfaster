package schema

// FormEmployeeEmpID is the query string schema for querying line UID by Employee ID
type FormEmployeeEmpID struct {
	EmpID string `form:"emp-id" binding:"required,len=7"`
}

// FormEmployeeLineUID is the query string schema for querying Employee ID by line UID
type FormEmployeeLineUID struct {
	LineUID string `form:"line-uid" binding:"required,len=33"`
}

// PostEmployee is the post schema for upserting employee by IDs
type PostEmployee struct {
	EmpID   string `json:"emp_id" binding:"required,len=7" example:"LW99999"`
	LineUID string `json:"line_uid" binding:"required,len=33" example:"U6664ceab1f4466b30827d936cee888e6"`
}
