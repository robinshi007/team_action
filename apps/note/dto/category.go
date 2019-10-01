package dto

// NewCategory -
type NewCategory struct {
	Name string `json:"name" binding:"required"`
}

// EditCategory -
type EditCategory struct {
	Name string `json:"name" binding:"required"`
}
