// This Go code snippet defines a model for creating a new Todo item. It includes fields for the title, description, completion status, creation and update timestamps, and an optional due date.

// Package model contains data structures for service layer operations.
package model

import "time"

// TodoCreate represents the data required to create a new Todo item.
type TodoCreate struct {
	Title       string     // Title of the Todo item.
	Description string     // Description of the Todo item.
	IsDone      bool       // Completion status of the Todo item.
	CreatedAt   time.Time  // Timestamp when the Todo item was created.
	UpdatedAt   time.Time  // Timestamp when the Todo item was last updated.
	DueDate     *time.Time // Optional due date for the Todo item.
}
