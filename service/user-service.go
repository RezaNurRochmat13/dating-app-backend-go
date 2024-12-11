package service

import (
	"dating-app-backend/entities"
	"dating-app-backend/repository"
)


func GetAllUsers() ([]entities.User, error) {
	return repository.FindAllUsers()
}

func GetUserByEmail(email string) (*entities.User, error) {
	return repository.FindUserByEmail(email)
}
