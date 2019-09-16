package dto

import (
	uuid "github.com/satori/go.uuid"
)

// NewNote-
type NewNote struct {
	CategoryID uuid.UUID `json:"category_id"`
	Title      string    `json:"title" binding:"required"`
	Body       string    `json:"body" binding:"required"`
}

// EditNote-
type EditNote struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}
