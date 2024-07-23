// Package request defines structures and methods for handling incoming requests related to Todo items.
// It includes structures for adding and updating Todo items, along with methods to convert these request structures into model representations.
package request

import (
	"time"
	"todoapp/service/model"
)

// AddTodoRequest represents the payload for adding a new Todo item.
// It includes the title and description of the Todo item.
type AddTodoRequest struct {
	Title       string `json:"title"`       // Title of the Todo item
	Description string `json:"description"` // Description of the Todo item
}

// UpdateTodoRequest represents the payload for updating an existing Todo item.
// It includes the completion status and the due date of the Todo item.
type UpdateTodoRequest struct {
	IsDone  bool      `json:"isDone"`  // Completion status of the Todo item
	DueDate time.Time `json:"dueDate"` // Due date of the Todo item
}

// ToModel converts an AddTodoRequest into a TodoCreate model.
// This method facilitates the transformation of request data into a format suitable for the service layer.
//
// Returns:
// - A model.TodoCreate struct populated with the title and description from the AddTodoRequest.
func (addTodoRequest AddTodoRequest) ToModel() model.TodoCreate {
	return model.TodoCreate{
		Title:       addTodoRequest.Title,       // Title from the request
		Description: addTodoRequest.Description, // Description from the request
	}
}
