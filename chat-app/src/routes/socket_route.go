package routes

import (
	"chat-app/src/core/config"
	"chat-app/src/handlers"

	"github.com/labstack/echo/v4"
)

// InitWebSocketRoutes initializes the WebSocket routes
func InitWebSocketRoutes(e *echo.Group, cfg *config.Config) {
	wsh := handlers.NewWebSocketHandler(cfg)

	e.Static("/", "public")

	// WebSocket endpoint
	e.GET("/ws", func(c echo.Context) error {
		return wsh.HandleConnections(c)
	})

	// Start the message handler
	go wsh.HandleMessages()
}
