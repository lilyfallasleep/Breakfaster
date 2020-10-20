package schema

import "time"

// SelectFood is the food schema with selected fields
type SelectFood struct {
	ID             int
	FoodName       string
	FoodSupplier   string
	PicURL         string
	SupplyDatetime time.Time
}
