package config

import (
	type_conversion "chat-app/src/core/helpers/type"
	"os"

	"github.com/joho/godotenv"
)

// Config is a struct that holds the configuration of the application.
type Config struct {
	MongoURI       string
	Mail           string
	MailPort       int
	MailPassword   string
	MailHost       string
	TypeConversion *type_conversion.TypeConversion
}

var instance *Config

// GetConfig returns the instance of the Config struct.
func GetConfig() *Config {
	if instance == nil {
		instance = &Config{}
		instance.LoadConfig()
	}
	return instance
}

// LoadConfig loads the configuration from the .env file.
func (config *Config) LoadConfig() {
	typeConversion := type_conversion.NewTypeConversion()
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config.MongoURI = os.Getenv("MONGO_URI")                                  // Load the mongo uri from the .env file
	config.Mail = os.Getenv("MAIL")                                           // Load the mail from the .env file
	config.MailPort = typeConversion.ParseStringToInt(os.Getenv("MAIL_PORT")) // Load the mail port from the .env file
	config.MailPassword = os.Getenv("MAIL_PASSWORD")                          // Load the mail password from the .env file
	config.MailHost = os.Getenv("MAIL_HOST")                                  // Load the mail host from the .env file
}
