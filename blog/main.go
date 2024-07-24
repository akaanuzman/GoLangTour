package main

import (
	"blog/common/db"
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

	log.Fatal(e.Start(":8080"))
}
