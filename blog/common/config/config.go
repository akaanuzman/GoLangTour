package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI            string
	JwtSecretKey        string
	JwtExpire           string
	ResetPasswordExpire int
}

func (config *Config) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config.MongoURI = os.Getenv("MONGO_URI")
	config.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	config.JwtExpire = os.Getenv("JWT_EXPIRE")
	resetPasswordExpString := os.Getenv("RESET_PASSWORD_EXPIRE")
	config.ResetPasswordExpire = parseResetPasswordExpire(resetPasswordExpString)
}

func parseResetPasswordExpire(resetPasswordExpire string) int {
	resetPasswordExp, err := strconv.Atoi(resetPasswordExpire)
	if err != nil {
		panic("Error parsing RESET_PASSWORD_EXPIRE")
	}
	return resetPasswordExp
}
