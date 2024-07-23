// Package response defines structures and methods for formatting outgoing responses related to Todo items.
// It includes structures for error messages and Todo item details, along with methods to convert domain models into these response structures.
package response

import (
	"time"
	"todoapp/domains"
)

// ErrorResponse represents the structure of an error message in the response.
// It includes a description of the error.
type ErrorResponse struct {
	ErrorDescription string `json:"errorDescription"` // Description of the error encountered
}

// TodoResponse represents the structure of a Todo item in the response.
// It includes details such as the ID, title, description, completion status, creation and update timestamps, and due date.
type TodoResponse struct {
	Id          int64      `json:"id"`          // Unique identifier of the Todo item
	Title       string     `json:"title"`       // Title of the Todo item
	Description string     `json:"description"` // Description of the Todo item
	IsDone      bool       `json:"isDone"`      // Completion status of the Todo item
	CreatedAt   time.Time  `json:"createdAt"`   // Timestamp when the Todo item was created
	UpdatedAt   time.Time  `json:"updatedAt"`   // Timestamp when the Todo item was last updated
	DueDate     *time.Time `json:"dueDate"`     // Optional due date of the Todo item
}

// ToResponse converts a Todo domain model into a TodoResponse.
// This method facilitates the transformation of domain data into a format suitable for the response.
//
// Parameters:
// - todo: A domains.Todo struct representing the domain model of a Todo item.
//
// Returns:
// - A TodoResponse struct populated with the details from the Todo domain model.
func ToResponse(todo domains.Todo) TodoResponse {
	return TodoResponse{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		DueDate:     todo.DueDate,
	}
}

// ToResponseList converts a slice of Todo domain models into a slice of TodoResponse structs.
// This method is useful for converting multiple Todo items into a format suitable for the response.
//
// Parameters:
// - todos: A slice of domains.Todo structs representing the domain models of Todo items.
//
// Returns:
// - A slice of TodoResponse structs populated with the details from the Todo domain models.
func ToResponseList(todos []domains.Todo) []TodoResponse {
	var todoResponseList = []TodoResponse{}
	for _, todo := range todos {
		todoResponseList = append(todoResponseList, ToResponse(todo))
	}
	return todoResponseList
}
