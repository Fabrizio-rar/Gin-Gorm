package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Gender   string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type UserEntry struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Email   string `gorm:"not null"`
	Title   string `gorm:"not null"`
	Content string
	User    User `gorm:"foreignKey:Email"`
}
