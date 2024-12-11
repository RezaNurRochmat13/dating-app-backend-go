package controller

import (
	"dating-app-backend/config"
	"dating-app-backend/entities"
	"dating-app-backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)



func Login(c echo.Context) error {
	var user entities.User
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind input JSON to the input struct
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Find the user by email
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	// Verify the password
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	// Generate a JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}

func RegisterUser(c echo.Context) error {
	var DB = config.DB
	var user entities.User

	// Bind user data from the request
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	// Update the user object with the hashed password
	user.Password = hashedPassword
	user.Quota = 10
	user.IsPremium = false

	// Save the user to the database
	if err := DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

