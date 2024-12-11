package controller

import (
	"dating-app-backend/entities"
	"dating-app-backend/service"
	"net/http"

	"github.com/labstack/echo/v4"
)


func GetAllUsers(c echo.Context) error {
	users, err := service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{"data": users})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": users})
}

func LikeUser(c echo.Context) error {
	likedIdParam := c.Param("id")
	var like entities.Like

	if err := c.Bind(&like); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := service.LikeUser(&like, likedIdParam); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": like})
}
