package service

import (
	"dating-app-backend/entities"
	"dating-app-backend/repository"
	"dating-app-backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserAuthLogin(payload *entities.User, c echo.Context) (token string, errors error) {
	// Find the user by email
	user, err := repository.FindUserByEmail(payload.Email)
	if err != nil {
		return "", c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	// Verify the password
	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		return "", c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	// Generate a JWT token
	token, err = utils.GenerateToken(user.ID)
	if err != nil {
		return "", c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return token, nil
}

func UserAuthRegister(payload *entities.User, c echo.Context) (errors error) {
	// Hash the password
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	// Update the user object with the hashed password
	payload.Password = hashedPassword
	payload.Quota = 10
	payload.IsPremium = false

	// Save the user to the database
	if err := repository.SaveUser(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return nil
}
