package main

import (
	"blog/src/db"
	"blog/src/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	db.ConnectDB()

	e := echo.New()

	// Initialize routes
	routes.InitRoutes(e)

	log.Fatal(e.Start("localhost:8080"))
}
