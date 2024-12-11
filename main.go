package main

import (
	"dating-app-backend/config"
	"dating-app-backend/entities"
	"dating-app-backend/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	println("Hello, World!")

	// Initialize database connection
	config.InitializeDB()

	// Migrate database models
	entities.MigrateDatabase()

	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")

	// Register routes
	routes.RegisterAuthRoutes(api.Group("/auth"))
	routes.RegisterUserRoutes(api.Group("/users"))

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ping sucessfully")
	})


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}