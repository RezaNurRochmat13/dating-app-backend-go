package entities

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	MatchesID uint
	SenderId int    `gorm:"size:255"`
	Content string `gorm:"size:255"`
	
}
