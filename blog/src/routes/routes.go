package routes

import (
	"blog/src/controllers"
	"blog/src/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	base := e.Group("/api/v1")

	auth := base.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	r := base.Group("/posts")
	r.Use(middlewares.JWTMiddleware())
	r.POST("", controllers.CreatePost)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)
}
