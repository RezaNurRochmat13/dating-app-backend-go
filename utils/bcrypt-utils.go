package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword checks if a provided password matches the hashed password
func CheckPasswordHash(password, hash string) bool {
	// bcrypt.CompareHashAndPassword returns nil if the hash matches the password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}