package domain

import (
	"time"
)

// Represents a story entity
type Story struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Represents the input for creating a new story
type NewStoryInput struct {
	Title   string `json:"title" validate:"required,min=3,max=255"`
	Author  string `json:"author" validate:"required,min=3,max=255"`
	Content string `json:"content" validate:"required,min=10"`
}

// Represents the input for updating a story
type UpdateStoryInput struct {
	Title   *string `json:"title" validate:"omitempty,min=3,max=255"`
	Author  *string `json:"author" validate:"omitempty,min=3,max=255"`
	Content *string `json:"content" validate:"omitempty,min=10"`
}
