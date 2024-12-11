package service

import (
	"dating-app-backend/entities"
	"dating-app-backend/repository"
	"strconv"
)


func GetAllUsers() ([]entities.User, error) {
	return repository.FindAllUsers()
}

func GetUserByEmail(email string) (*entities.User, error) {
	return repository.FindUserByEmail(email)
}

func LikeUser(payload *entities.Like, likedIdParam string) error {
	likedId, _ := strconv.Atoi(likedIdParam)

	payload.LikedId = uint(likedId)

	return repository.LikeUser(payload)
}
