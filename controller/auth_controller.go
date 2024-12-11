package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "Login")
}

func Register(c echo.Context) error {
	return c.JSON(http.StatusOK, "Register")
}
