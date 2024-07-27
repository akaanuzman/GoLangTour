package main

import (
	"log"
	"note-app/src/core/db"
	"note-app/src/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database := db.Database{}

	// Connect to the database
	database.ConnectDB()

	e := echo.New()

	// Initialize the routes
	routes.InitRoutes(e)

	log.Fatal(e.Start("localhost:8080"))
}
