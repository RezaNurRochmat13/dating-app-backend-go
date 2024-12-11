package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"unique","size:100"`
	Password string `gorm:"size:255"`
	Gender   string `gorm:"size:255"`
	DateOfBirth string `gorm:"size:255"`
	Bio      string `gorm:"size:255"`
	ProfilePicture string `gorm:"size:255"`
	Quota    int    `gorm:"size:255"`
	IsPremium bool  `gorm:"size:255"`
	Like     []Like `gorm:"foreignKey:LikerId"`
	Profile Profile `gorm:"foreignKey:UserId"`
	Matches []Matches
}
