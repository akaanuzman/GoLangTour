package service

import (
	"errors"
	"time"
	"todoapp/domains"
	"todoapp/persistence"
)

type FakeTodoRepository struct {
	todos []domains.Todo
}

func NewFakeTodoRepository(initialTodos []domains.Todo) persistence.ITodoRepository {
	return &FakeTodoRepository{todos: initialTodos}
}

func (todoRepository *FakeTodoRepository) GetAllTodos() []domains.Todo {
	return todoRepository.todos
}

func (todoRepository *FakeTodoRepository) GetTodoById(id int64) (domains.Todo, error) {
	for _, todo := range todoRepository.todos {
		if todo.Id == id {
			return todo, nil
		}
	}
	return domains.Todo{}, errors.New("todo not found")
}

func (todoRepository *FakeTodoRepository) GetDoneOrUndoneTodos(isDone bool) []domains.Todo {
	var filteredTodos []domains.Todo
	for _, todo := range todoRepository.todos {
		if todo.IsDone == isDone {
			filteredTodos = append(filteredTodos, todo)
		}
	}
	return filteredTodos
}

func (todoRepository *FakeTodoRepository) AddNewTodo(todo domains.Todo) error {
	todo.Id = int64(len(todoRepository.todos) + 1)
	todoRepository.todos = append(todoRepository.todos, todo)
	return nil
}

func (todoRepository *FakeTodoRepository) DeleteTodoById(id int64) error {
	for i, todo := range todoRepository.todos {
		if todo.Id == id {
			todoRepository.todos = append(todoRepository.todos[:i], todoRepository.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}

func (todoRepository *FakeTodoRepository) SignTodoAsDone(id int64, dueDate time.Time) error {
	for i, todo := range todoRepository.todos {
		if todo.Id == id {
			todoRepository.todos[i].IsDone = true
			todoRepository.todos[i].DueDate = &dueDate
			return nil
		}
	}
	return errors.New("todo not found")
}

func (todoRepository *FakeTodoRepository) SingTodoAsUndone(id int64) error {
	for i, todo := range todoRepository.todos {
		if todo.Id == id {
			todoRepository.todos[i].IsDone = false
			return nil
		}
	}
	return errors.New("todo not found")
}
