package main

import (
	"dating-app-backend/config"
	"dating-app-backend/entities"
)

func main() {
	println("Hello, World!")

	// Initialize database connection
	config.InitializeDB()

	// Migrate database models
	entities.MigrateDatabase()
}