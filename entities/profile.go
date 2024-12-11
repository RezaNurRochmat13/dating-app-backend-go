package entities

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	UserId uint   `gorm:"unique;not null"` // Foreign key
	Interest string `gorm:"size:255"`
	Location string `gorm:"size:255"`
	Preferences string `gorm:"size:255"`
}
