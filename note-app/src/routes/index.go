package routes

import (
	"note-app/src/controllers"
	"note-app/src/middlewares"
	"note-app/src/services"

	"github.com/labstack/echo/v4"
)

// Initialize the routes for the application

func InitRoutes(e *echo.Echo) {
	base := e.Group("/api/v1")

	initAuthRoutes(base)
}

// initAuthRoutes initializes the routes for the authentication endpoints.
func initAuthRoutes(e *echo.Group) {
	auth := e.Group("/auth")
	authService := services.NewUserService()
	authController := controllers.NewUserController(authService)
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
	auth.POST("/forgot-password", authController.ForgotPassword)

	// Group requiring JWT middleware
	protected := auth.Group("")
	protected.Use(middlewares.JWTMiddleware())
	protected.POST("/change-password", authController.ChangePassword)
}
