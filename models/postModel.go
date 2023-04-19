package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Gender   string
	Email    string
	Password string
}
