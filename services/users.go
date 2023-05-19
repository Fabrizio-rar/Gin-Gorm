package services

import (
	"Gin-gorm/db"
	"Gin-gorm/models"
	"Gin-gorm/structs"
	"Gin-gorm/utils"
	"errors"
	"fmt"
)

func CreateUser(email, password, name, gender string) (err error) {
	userAlreadyExists, err := db.UserExists(email)
	if err != nil {
		fmt.Println("Error in UserExists:", err.Error())
		return
	}

	if userAlreadyExists {
		err = errors.New("user already exists")
		return
	}

	newUser := models.User{Name: name, Gender: gender, Email: email, Password: password}
	err = db.CreateUser(newUser)
	if err != nil {
		fmt.Println("Error in CreateUser:", err.Error())
		return
	}
	return
}

func GetUser(email string) (user models.User, err error) {
	user, err = db.GetUser(email)
	if err != nil {
		fmt.Println("Error in GetUser:", err.Error())
		return
	}
	return
}

func GetAllUsers() (users []structs.GetUserResp, err error) {
	users, err = db.GetAllUsers()
	if err != nil {
		fmt.Println("Error in GetAllUsers:", err.Error())
		return
	}
	return
}

func DeleteUser(email, password string) (err error) {
	user, err := GetUser(email)
	if err != nil {
		fmt.Println("Error in GetUser:", err.Error())
		return
	}

	credentialsValid := utils.PasswordValid(user.Password, password)

	if credentialsValid {
		err = db.DeleteUser(email)
		if err != nil {
			fmt.Println("Error in DeleteUser:", err.Error())
			return
		}
	} else {
		err = errors.New("invalid email password combination")
		return
	}

	return
}
