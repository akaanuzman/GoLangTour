package models

import (
	"chat-app/src/utils"
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var hashManager = utils.NewHashManager()

// User represents a user in the system
type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email            string             `bson:"email" json:"email"`
	Password         string             `bson:"password,omitempty" json:"password,omitempty"`
	CreatedAt        primitive.DateTime `bson:"createdAt" json:"createdAt"`
	ResetToken       string             `bson:"reset_token,omitempty" json:"reset_token,omitempty"`
	ResetTokenExpiry *time.Time         `bson:"reset_token_expiry,omitempty" json:"reset_token_expiry,omitempty"`
}

// ValidateEmail checks if the email is in valid format
func ValidateEmail(email string) error {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !regex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// HashPassword hashes the user's password
func (u *User) HashPassword() error {
	password, err := hashManager.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = password
	return nil
}

// ComparePassword compares the provided password with the stored hashed password
func (u *User) ComparePassword(password string) error {
	if !hashManager.ComparePassword(password, u.Password) {
		return errors.New("invalid password")
	}
	return nil
}

// GenerateJWT generates a JWT token for the user
func (u *User) GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["user_id"] = u.ID.Hex()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}
