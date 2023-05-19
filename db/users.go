package db

import (
	"Gin-gorm/initializers"
	"Gin-gorm/models"
	"Gin-gorm/structs"
	"fmt"

	"gorm.io/gorm"
)

func UserExists(email string) (exists bool, err error) {
	result := initializers.DB.Where("email = ?", email).First(&models.User{})
	err = result.Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
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
	result := initializers.DB.Where("email = ?", email).Find(&user)

	if result.RowsAffected == 0 {
		err = fmt.Errorf("user does not exist")
		return
	}
	return
}

func GetAllUsers() (users []structs.GetUserResp, err error) {
	result := initializers.DB.Find(&users)
	err = result.Error
	return
}

func DeleteUser(email string) (err error) {
	result := initializers.DB.Where("email = ?", email).Delete(&models.User{})
	err = result.Error
	return
}
