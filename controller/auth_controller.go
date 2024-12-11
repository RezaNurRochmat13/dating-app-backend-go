package controller

import (
	"dating-app-backend/entities"
	"dating-app-backend/service"
	"net/http"

	"github.com/labstack/echo/v4"
)



func Login(c echo.Context) error {
	var user entities.User

	// Bind input JSON to the input struct
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	token, err := service.UserAuthLogin(&user, c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}

func RegisterUser(c echo.Context) error {
	var user entities.User

	// Bind user data from the request
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := service.UserAuthRegister(&user, c); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

