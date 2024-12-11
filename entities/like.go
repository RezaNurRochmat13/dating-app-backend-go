package entities

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	LikerId uint
	LikedId uint
}
