package entities

import "gorm.io/gorm"

type Matches struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Message []Message
	UserId uint
	MatchedAt string `gorm:"type:date"`
}
