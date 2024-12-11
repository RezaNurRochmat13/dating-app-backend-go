package entities

import (
	"dating-app-backend/config"
	"log"
)


func MigrateDatabase() {
	db := config.DB
	err := db.AutoMigrate(&User{}, &Profile{}, &Matches{}, &Message{})
	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}
	log.Println("Database migrated successfully!")
}
