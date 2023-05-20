package services

import (
	"Gin-gorm/db"
	"Gin-gorm/models"
	"Gin-gorm/utils"
	"errors"
	"fmt"
)

func CreateEntry(email, title, content string) (err error) {
	entryAlreadyExists, err := db.EntryExists(email, title)

	if err != nil {
		fmt.Println("Error in EntryExists:", err.Error())
		return
	}

	if entryAlreadyExists {
		err = errors.New("entry already exists")
		return
	}

	newEntry := models.UserEntry{Email: email, Title: title, Content: content}
	err = db.CreateEntry(newEntry)
	if err != nil {
		fmt.Println("Error in CreateEntry:", err.Error())
		return
	}
	return
}

func GetEntry(email, title string) (entry models.UserEntry, err error) {
	entry, err = db.GetEntry(email, title)
	if err != nil {
		fmt.Println("Error in GetEntry:", err.Error())
		return
	}
	return
}

func GetAllEntriesFromUser(email string) (entries []models.UserEntry, err error) {
	entries, err = db.GetAllEntriesFromUser(email)
	if err != nil {
		fmt.Println("Error in GetAllUsers:", err.Error())
		return
	}
	return
}

func UpdateEntry(email, password, title, content string) (err error) {
	user, err := GetUser(email)
	if err != nil {
		fmt.Println("Error in GetUser:", err.Error())
		return
	}

	credentialsValid := utils.PasswordValid(user.Password, password)

	if credentialsValid {
		err = db.UpdateEntry(email, title, content)
		if err != nil {
			fmt.Println("Error in UpdateEntry:", err.Error())
			return
		}
	} else {
		err = errors.New("invalid email password combination")
		return
	}

	return
}

func DeleteEntry(email, password, title string) (err error) {
	user, err := GetUser(email)
	if err != nil {
		fmt.Println("Error in GetUser:", err.Error())
		return
	}

	credentialsValid := utils.PasswordValid(user.Password, password)

	if credentialsValid {
		err = db.DeleteEntry(email, title)
		if err != nil {
			fmt.Println("Error in DeleteEntry:", err.Error())
			return
		}
	} else {
		err = errors.New("invalid email password combination")
		return
	}

	return
}
