package service

import (
	"os"
	"testing"
	"time"
	"todoapp/domains"
	"todoapp/service"
	"todoapp/service/model"

	"github.com/stretchr/testify/assert"
)

var todoService service.ITodoService

func TestMain(m *testing.M) {
	initialTodos := []domains.Todo{
		{
			Id:          1,
			Title:       "Learn Golang",
			Description: "Learn Golang basics",
			IsDone:      false,
		},
		{
			Id:          2,
			Title:       "Learn .NET",
			Description: "Learn .NET basics",
			IsDone:      false,
		},
		{
			Id:          3,
			Title:       "Learn Java",
			Description: "Learn Java basics",
			IsDone:      false,
		},
	}

	fakeTodoRepository := NewFakeTodoRepository(initialTodos)

	todoService = service.NewTodoService(fakeTodoRepository)
	os.Exit(m.Run())
}

func Test_ShouldGetAllTodos(t *testing.T) {
	t.Run("Should return all todos", func(t *testing.T) {
		actualTodos := todoService.GetAllTodos()
		assert.Equal(t, 3, len(actualTodos))
	})
}

func Test_ShouldGetTodoById(t *testing.T) {
	t.Run("Should return todo by id", func(t *testing.T) {
		actualTodo, _ := todoService.GetTodoById(1)
		assert.Equal(t, "Learn Golang", actualTodo.Title)
	})
}

func Test_ShouldGetDoneTodos(t *testing.T) {
	t.Run("Should return done todos", func(t *testing.T) {
		actualTodos := todoService.GetDoneOrUndoneTodos(true)
		assert.Equal(t, 0, len(actualTodos))
	})
}

func Test_ShouldGetUndoneTodos(t *testing.T) {
	t.Run("Should return undone todos", func(t *testing.T) {
		actualTodos := todoService.GetDoneOrUndoneTodos(false)
		assert.Equal(t, 3, len(actualTodos))
	})
}

func Test_ShouldAddNewTodo(t *testing.T) {
	t.Run("Should add new todo", func(t *testing.T) {
		todoService.AddNewTodo(model.TodoCreate{
			Title:       "Learn Python",
			Description: "Learn Python basics",
			IsDone:      false,
		})
		actualTodos := todoService.GetAllTodos()
		assert.Equal(t, 4, len(actualTodos))
	})
}

func Test_ShouldAddNewTodoFailValidation(t *testing.T) {
	t.Run("Should fail validation", func(t *testing.T) {
		err := todoService.AddNewTodo(model.TodoCreate{
			Title:       "",
			Description: "Learn Python basics",
			IsDone:      false,
		})
		assert.NotNil(t, err)
	})
}

func Test_ShouldDeleteTodoById(t *testing.T) {
	t.Run("Should delete todo by id", func(t *testing.T) {
		todoService.DeleteTodoById(1)
		actualTodos := todoService.GetAllTodos()
		assert.Equal(t, 2, len(actualTodos))
	})
}

func Test_ShouldSignTodoAsDone(t *testing.T) {
	t.Run("Should sign todo as done", func(t *testing.T) {
		todoService.SignTodoAsDone(2, time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC))
		actualTodo, _ := todoService.GetTodoById(2)
		assert.Equal(t, true, actualTodo.IsDone)
	})
}

func Test_ShouldSignTodoAsUndone(t *testing.T) {
	t.Run("Should sign todo as undone", func(t *testing.T) {
		todoService.SingTodoAsUndone(2)
		actualTodo, _ := todoService.GetTodoById(2)
		assert.Equal(t, false, actualTodo.IsDone)
	})
}
