package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("your-secret-key")

func GenerateToken(userID int) (string, error) {
	// Define claims for the token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	return token.SignedString(jwtSecret)
}