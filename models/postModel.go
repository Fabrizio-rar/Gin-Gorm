package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Gender   string
	Email    string
	Password string
}

type UserEntry struct {
	gorm.Model
	UserID  int
	Title   string
	Content string
	User    User `gorm:"foreignKey:UserID"`
}
