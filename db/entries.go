package db

import (
	"Gin-gorm/initializers"
	"Gin-gorm/models"
	"fmt"

	"gorm.io/gorm"
)

func EntryExists(email, title string) (exists bool, err error) {
	result := initializers.DB.Where("email = ? AND title = ?", email).First(&models.UserEntry{})
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

func CreateEntry(newEntry models.UserEntry) (err error) {
	result := initializers.DB.Create(&newEntry)
	err = result.Error
	return
}

func GetEntry(email, title string) (entry models.UserEntry, err error) {
	result := initializers.DB.Where("email = ? AND title = ?", email, title).Find(&entry)

	if result.RowsAffected == 0 {
		err = fmt.Errorf("entry does not exist")
		return
	}
	return
}

func GetAllEntriesFromUser(email string) (entries []models.UserEntry, err error) {
	result := initializers.DB.Where("email = ?", email).Find(&entries)
	err = result.Error
	return
}
