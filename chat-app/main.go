package main

import (
	"chat-app/src/core/config"
	"chat-app/src/core/db"
	"chat-app/src/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the MongoDB URI from the configuration
	cfg := config.GetConfig()

	database := db.Database{}

	// Connect to the database
	database.ConnectDB(cfg)

	e := echo.New()

	// Initialize the routes
	routes.InitRoutes(e, cfg)

	log.Fatal(e.Start("localhost:8080"))
}
