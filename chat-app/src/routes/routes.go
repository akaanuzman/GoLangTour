package routes

import (
	"chat-app/src/core/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Initialize the routes for the application
func InitRoutes(e *echo.Echo, cfg *config.Config) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	base := e.Group("/api/v1")

	// Initialize the WebSocket routes
	InitWebSocketRoutes(base)

	// Initialize the authentication routes
	InitAuthRoutes(base, cfg)
}
