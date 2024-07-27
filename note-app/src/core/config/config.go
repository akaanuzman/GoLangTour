package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI     string
	Mail         string
	MailPort     int
	MailPassword string
	MailHost     string
}

var instance *Config

func GetConfig() *Config {
	if instance == nil {
		instance = &Config{}
		instance.LoadConfig()
	}
	return instance
}

func (config *Config) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config.MongoURI = os.Getenv("MONGO_URI")
	config.Mail = os.Getenv("MAIL")
	config.MailPort = parseStringToInt(os.Getenv("MAIL_PORT"))
	config.MailPassword = os.Getenv("MAIL_PASSWORD")
	config.MailHost = os.Getenv("MAIL_HOST")
}

func parseStringToInt(str string) int {
	parsedInt, err := strconv.Atoi(str)
	if err != nil {
		panic("Error parsing MAIL_PORT")
	}
	return parsedInt
}
