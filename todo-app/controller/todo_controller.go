// Package controller defines the TodoController struct and its methods to handle HTTP requests for Todo operations.
// It includes route registration and handlers for CRUD operations on Todo items.
package controller

import (
	"net/http"
	"strconv"
	"todoapp/controller/request"
	"todoapp/controller/response"
	"todoapp/service"
	"todoapp/service/model"

	"github.com/labstack/echo/v4"
)

// TodoController encapsulates dependencies for handling Todo-related HTTP requests.
type TodoController struct {
	todoService service.ITodoService // Interface to Todo service layer
}

// NewTodoController creates a new instance of TodoController with the provided ITodoService.
func NewTodoController(todoService service.ITodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

// RegisterRoutes registers the HTTP routes for the TodoController with the provided Echo instance.
func (controller *TodoController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/todos", controller.GetAllTodos)
	e.GET("/api/v1/todos/:id", controller.GetTodoById)
	e.GET("/api/v1/todos/done", controller.GetDoneOrUndoneTodos)
	e.GET("/api/v1/todos/undone", controller.GetDoneOrUndoneTodos)
	e.POST("/api/v1/todos", controller.AddNewTodo)
	e.DELETE("/api/v1/todos/:id", controller.DeleteTodoById)
	e.PUT("/api/v1/todos/:id/done", controller.SignTodoAsDone)
	e.PUT("/api/v1/todos/:id/undone", controller.SignTodoAsUndone)
}

// GetAllTodos handles the GET request to retrieve all Todo items.
func (controller *TodoController) GetAllTodos(c echo.Context) error {
	allTodos := controller.todoService.GetAllTodos()
	return c.JSON(http.StatusOK, response.ToResponseList(allTodos))
}

// GetTodoById handles the GET request to retrieve a Todo item by its ID.
func (controller *TodoController) GetTodoById(c echo.Context) error {
	param := c.Param("id")
	todoId, _ := strconv.Atoi(param)

	todo, err := controller.todoService.GetTodoById(int64(todoId))
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ToResponse(todo))
}

// GetDoneOrUndoneTodos handles the GET request to retrieve Todo items filtered by their completion status.
func (controller *TodoController) GetDoneOrUndoneTodos(c echo.Context) error {
	isDone, _ := strconv.ParseBool(c.QueryParam("isDone"))

	todos := controller.todoService.GetDoneOrUndoneTodos(isDone)
	return c.JSON(http.StatusOK, response.ToResponseList(todos))
}

// AddNewTodo handles the POST request to add a new Todo item.
func (controller *TodoController) AddNewTodo(c echo.Context) error {
	var addTodoRequest request.AddTodoRequest
	if bindErr := c.Bind(&addTodoRequest); bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}

	todoToAdd := model.TodoCreate{
		Title:       addTodoRequest.Title,
		Description: addTodoRequest.Description,
	}

	if addErr := controller.todoService.AddNewTodo(todoToAdd); addErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: addErr.Error(),
		})
	}
	return c.JSON(http.StatusCreated, addTodoRequest)
}

// DeleteTodoById handles the DELETE request to remove a Todo item by its ID.
func (controller *TodoController) DeleteTodoById(c echo.Context) error {
	param := c.Param("id")
	todoId, _ := strconv.Atoi(param)

	if deleteErr := controller.todoService.DeleteTodoById(int64(todoId)); deleteErr != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: deleteErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

// SignTodoAsDone handles the PUT request to mark a Todo item as completed.
func (controller *TodoController) SignTodoAsDone(c echo.Context) error {
	param := c.Param("id")
	todoId, _ := strconv.Atoi(param)

	var signTodoAsDoneRequest request.UpdateTodoRequest
	if bindErr := c.Bind(&signTodoAsDoneRequest); bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}

	if signErr := controller.todoService.SignTodoAsDone(int64(todoId), signTodoAsDoneRequest.DueDate); signErr != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: signErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

// SignTodoAsUndone handles the PUT request to mark a Todo item as not completed.
func (controller *TodoController) SignTodoAsUndone(c echo.Context) error {
	param := c.Param("id")
	todoId, _ := strconv.Atoi(param)

	if signErr := controller.todoService.SingTodoAsUndone(int64(todoId)); signErr != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: signErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}
