package dto

// NewNote-
type NewNote struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

// EditNote-
type EditNote struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}
