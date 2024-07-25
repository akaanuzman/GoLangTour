package routes

import (
	"blog/src/controllers"
	"blog/src/middlewares"
	"blog/src/services"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	base := e.Group("/api/v1")

	initAuthRoutes(base)

	r := base.Group("/posts")
	r.Use(middlewares.JWTMiddleware())
	r.POST("", controllers.CreatePost)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)
}

// initAuthRoutes initializes the routes for the authentication endpoints.
func initAuthRoutes(e *echo.Group) {
	auth := e.Group("/auth")
	authService := services.NewUserService()
	authController := controllers.NewUserController(authService)
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
}
