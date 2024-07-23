// Package domains defines the domain model for the Todo application.
package domains

import "time"

// Todo represents the domain model for a Todo item, including its metadata and state.
type Todo struct {
	Id          int64      // Unique identifier for the Todo item
	Title       string     // Title of the Todo item
	Description string     // Detailed description of the Todo item
	IsDone      bool       // Flag indicating whether the Todo item is completed
	CreatedAt   time.Time  // Timestamp of when the Todo item was created
	UpdatedAt   time.Time  // Timestamp of the last update to the Todo item
	DueDate     *time.Time // Optional due date for the Todo item
}
