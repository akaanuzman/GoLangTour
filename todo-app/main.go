package main

import (
	"context"
	"todoapp/common/app"
	"todoapp/common/postresql"
	"todoapp/controller"
	"todoapp/persistence"
	"todoapp/service"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	configurationManager := app.NewConfigurationManager()

	dbPool := postresql.GetConnectionPool(ctx, configurationManager.PostreSqlConfig)

	todoRepository := persistence.NewTodoRepository(dbPool)

	todoService := service.NewTodoService(todoRepository)

	todoController := controller.NewTodoController(todoService)

	todoController.RegisterRoutes(e)

	e.Start("localhost:8080")
}
