package routes

import (
	"dating-app-backend/controller"

	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Group) {
	e.POST("/register", controller.RegisterUser)
	e.POST("/login", controller.Login)
}
