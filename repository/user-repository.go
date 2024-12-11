package repository

import (
	"dating-app-backend/config"
	"dating-app-backend/entities"
)


func FindAllUsers() ([]entities.User, error) {
	var DB = config.DB
	var users []entities.User
	if err := DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func SaveUser(user *entities.User) error {
	var DB = config.DB
	return DB.Create(&user).Error
}

func FindUserByEmail(email string) (*entities.User, error) {
	var DB = config.DB
	var user entities.User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}