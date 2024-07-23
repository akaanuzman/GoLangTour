package service

import (
	"errors"
	"time"
	"todoapp/domains"
	"todoapp/persistence"
	"todoapp/service/model"
)

type ITodoService interface {
	GetAllTodos() []domains.Todo
	GetTodoById(id int64) (domains.Todo, error)
	GetDoneOrUndoneTodos(isDone bool) []domains.Todo
	AddNewTodo(todo model.TodoCreate) error
	DeleteTodoById(id int64) error
	SignTodoAsDone(id int64, dueDate time.Time) error
	SingTodoAsUndone(id int64) error
}

type TodoService struct {
	todoRepository persistence.ITodoRepository
}

func NewTodoService(todoRepository persistence.ITodoRepository) ITodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (todoService *TodoService) GetAllTodos() []domains.Todo {
	return todoService.todoRepository.GetAllTodos()
}

func (todoService *TodoService) GetTodoById(id int64) (domains.Todo, error) {
	return todoService.todoRepository.GetTodoById(id)
}

func (todoService *TodoService) GetDoneOrUndoneTodos(isDone bool) []domains.Todo {
	return todoService.todoRepository.GetDoneOrUndoneTodos(isDone)
}

func (todoService *TodoService) AddNewTodo(todo model.TodoCreate) error {
	validateErr := validateTodoCreate(todo)
	if validateErr != nil {
		return validateErr
	}

	todoToAdd := domains.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		DueDate:     todo.DueDate,
	}

	return todoService.todoRepository.AddNewTodo(todoToAdd)
}

func (todoService *TodoService) DeleteTodoById(id int64) error {
	return todoService.todoRepository.DeleteTodoById(id)
}

func (todoService *TodoService) SignTodoAsDone(id int64, dueDate time.Time) error {
	return todoService.todoRepository.SignTodoAsDone(id, dueDate)
}

func (todoService *TodoService) SingTodoAsUndone(id int64) error {
	return todoService.todoRepository.SingTodoAsUndone(id)
}

func validateTodoCreate(todo model.TodoCreate) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}

	if todo.Description == "" {
		return errors.New("description is required")
	}

	return nil
}
