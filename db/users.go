package db

import (
	"Gin-gorm/initializers"
	"Gin-gorm/models"
	"fmt"

	"gorm.io/gorm"
)

func UserExists(email string) (exists bool, err error) {
	result := initializers.DB.Where("email = ?", email).First(&models.User{})
	err = result.Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			exists = false
			return
		}
	} else {
		exists = true
	}
	return
}

func CreateUser(newUser models.User) (err error) {
	result := initializers.DB.Create(&newUser)
	err = result.Error
	return
}

func GetUser(email string) (user models.User, err error) {
	result := initializers.DB.Where("email = ?", email).First(&user)
	err = result.Error

	if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("user does not exist")
		return
	}
	return
}

func GetAllUsers() (users []models.User, err error) {
	result := initializers.DB.Find(&users)
	err = result.Error
	return
}

func DeleteUser(email string) (err error) {
	result := initializers.DB.Where("email = ?", email).Delete(&models.User{})
	err = result.Error
	return
}
