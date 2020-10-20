package schema

// FormFoodDate is the query string schema for food time interval
type FormFoodDate struct {
	StartDate string `form:"start" binding:"required,len=10"`
	EndDate   string `form:"end" binding:"required,len=10"`
}
