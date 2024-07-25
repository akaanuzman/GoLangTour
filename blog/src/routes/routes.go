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
	initPostRoutes(base)
}

// initAuthRoutes initializes the routes for the authentication endpoints.
func initAuthRoutes(e *echo.Group) {
	auth := e.Group("/auth")
	authService := services.NewUserService()
	authController := controllers.NewUserController(authService)
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
}

// initPostRoutes initializes the routes for the post endpoints.
func initPostRoutes(e *echo.Group) {
	postService := services.NewPostService()
	postController := controllers.NewPostController(postService)

	r := e.Group("/posts")
	r.Use(middlewares.JWTMiddleware())
	r.POST("", postController.CreatePost)
	r.PUT("/:id", postController.UpdatePost)
	r.DELETE("/:id", postController.DeletePost)
	r.GET("", postController.GetAllPosts)
	r.GET("/:id", postController.GetPostById)
}
