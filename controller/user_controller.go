package controller

import (
	"dating-app-backend/config"
	"dating-app-backend/entities"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


func GetAllUsers(c echo.Context) error {
	var DB = config.DB
	var users []entities.User
	DB.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{"data": users})
}

func LikeUser(c echo.Context) error {
	var DB = config.DB
	likerIdParam := c.FormValue("liker_id")
	likedIdParam := c.Param("id")

	likerId, _ := strconv.Atoi(likerIdParam)
	likedId, _ := strconv.Atoi(likedIdParam)

	like := entities.Like{
		LikerId: uint(likerId),
		LikedId: uint(likedId),
	}

	if err := DB.Create(&like).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": like})
}
