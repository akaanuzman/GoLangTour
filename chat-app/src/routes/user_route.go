package routes

import (
	"chat-app/src/controllers"
	"chat-app/src/core/config"
	"chat-app/src/middlewares"
	"chat-app/src/services"

	"github.com/labstack/echo/v4"
)

// InitUserRoutes initializes the routes for the user endpoints.
func InitUserRoutes(e *echo.Group, cfg *config.Config) {
	user := e.Group("/user")
	userService := services.NewUserService(cfg)
	userController := controllers.NewUserController(userService)

	user.Use(middlewares.JWTMiddleware())
	user.GET("/:id", userController.GetUserByID)
	user.GET("", userController.GetUserByEmail)
	user.DELETE("/:id", userController.DeleteUser)
}
