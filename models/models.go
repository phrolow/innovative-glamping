package models

import (
	"gorm.io/gorm"
)

// User представляет данные пользователя
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);uniqueIndex;not null"`
}

// Post представляет данные поста
type Post struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null"`
	Content string `gorm:"type:text;not null"`
	UserID  uint   `gorm:"not null"`
}
