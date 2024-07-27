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
	initNoteRoutes(base)
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

// initNoteRoutes initializes the routes for the note endpoints.
func initNoteRoutes(e *echo.Group) {
	note := e.Group("/notes")
	noteService := services.NewNoteService()
	noteController := controllers.NewNoteController(noteService)

	note.Use(middlewares.JWTMiddleware())
	note.GET("", noteController.GetAllNotes)
	note.GET("/:id", noteController.GetNoteById)
	note.GET("/user/:userId", noteController.GetNoteByUserId)
	note.GET("/tags", noteController.GetNotesByTags)
	note.POST("", noteController.CreateNote)
	note.PUT("/:id", noteController.UpdateNote)
	note.DELETE("/:id", noteController.DeleteNote)
}
