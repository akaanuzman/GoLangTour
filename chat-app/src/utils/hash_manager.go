package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

type IHashManager interface {
	// Hash hashes the provided password.
	HashPassword(password string) (string, error)
	// Compare compares the provided password with the provided hash.
	ComparePassword(password, hash string) bool
	// Hash hashes the provided messages.
	HashMessage(message string) string
}

// HashManager is a struct that implements the IHashManager interface.
type HashManager struct{}

// NewHashManager creates a new HashManager.
func NewHashManager() IHashManager {
	return &HashManager{}
}

// Hash hashes the provided password.
func (hm *HashManager) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	password = string(hashedPassword)
	return password, nil
}

// Compare compares the provided password with the provided hash.
func (hm *HashManager) ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Hash hashes the provided messages.
func (hm *HashManager) HashMessage(content string) string {
	hash := sha256.New()
	hash.Write([]byte(content))
	return hex.EncodeToString(hash.Sum(nil))
}
