package routes

import (
	"chat-app/src/controllers"
	"chat-app/src/core/config"
	"chat-app/src/middlewares"
	"chat-app/src/services"

	"github.com/labstack/echo/v4"
)

// InitAuthRoutes initializes the routes for the authentication endpoints.
func InitAuthRoutes(e *echo.Group, cfg *config.Config) {
	auth := e.Group("/auth")
	authService := services.NewUserService(cfg)
	authController := controllers.NewUserController(authService)
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
	auth.POST("/forgot-password", authController.ForgotPassword)

	// Group requiring JWT middleware
	protected := auth.Group("")
	protected.Use(middlewares.JWTMiddleware())
	protected.POST("/change-password", authController.ChangePassword)
}
