package schema

// JSONTimeInterval is the api schema for returned time interval
type JSONTimeInterval struct {
	StartDate string `json:"start" example:"2020-09-01"`
	EndDate   string `json:"end" example:"2020-09-05"`
}
