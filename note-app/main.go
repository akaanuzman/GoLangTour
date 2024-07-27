package main

import (
	"log"
	"note-app/src/core/db"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
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

	log.Fatal(e.Start("localhost:8080"))
}
