package schema

// JSONFood is the service schema for returned json food
type JSONFood struct {
	ID       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"burger"`
	Supplier string `json:"supplier" example:"McDonald"`
	PicURL   string `json:"picurl" example:"www.example.com"`
}

// NestedFood is the service schema for mapping datetime to food
type NestedFood map[string][]JSONFood
