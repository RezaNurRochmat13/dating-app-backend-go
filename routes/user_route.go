package routes

import (
	"dating-app-backend/controller"

	"github.com/labstack/echo/v4"
)


func RegisterUserRoutes(e *echo.Group) {
	e.GET("/", controller.GetAllUsers)
	e.POST("/like", controller.LikeUser)
}
