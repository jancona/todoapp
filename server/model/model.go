package model

import "github.com/google/uuid"

// ToDoInput is a structure that models todo information
type ToDoInput struct {
	Title     string `json:"title" binding:"required"`
	Order     int    `json:"order"`
	Completed bool   `json:"completed"`
}

// ToDo is a structure that models todo information
type ToDo struct {
	ToDoInput
	ID  *uuid.UUID `json:"id,omitempty"`
	URL string     `json:"url" format:"uuid"`
}

// Error represents an HTTP error status
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
