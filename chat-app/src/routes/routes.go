package routes

import (
	"chat-app/src/core/config"

	"github.com/labstack/echo/v4"
)

// Initialize the routes for the application
func InitRoutes(e *echo.Echo, cfg *config.Config) {
	base := e.Group("/api/v1")

	InitAuthRoutes(base, cfg)
}
