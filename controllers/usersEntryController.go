package controllers

import (
	"Gin-test/initializers"
	"Gin-test/models"
	"Gin-test/structs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateEntry(c *gin.Context) {
	var body models.UserEntry
	err := c.BindJSON(&body)
	if err != nil {
		c.Status(400)
		return
	}

	var existingEntry models.UserEntry

	entryExists := initializers.DB.Where("title = ? AND email = ?", body.Title, body.Email).First(&existingEntry)

	if entryExists.Error == gorm.ErrRecordNotFound {
		userEntry := models.UserEntry{Title: body.Title, Content: body.Content, Email: body.Email}
		createResult := initializers.DB.Create(&userEntry)

		if createResult.Error != nil {
			c.Status(400)
			return
		}

		c.JSON(200, gin.H{
			"message": "Entry created successfully",
		})
		return
	}

	c.JSON(400, gin.H{
		"message": "Entry with same title already exists",
	})
}

func GetEntry(c *gin.Context) {
	entryTitle := c.Query("title")
	entryEmail := c.Query("email")
	var userEntry models.UserEntry

	findResult := initializers.DB.Where("title = ? AND email = ?", entryTitle, entryEmail).Find(&userEntry)

	if findResult.RowsAffected == 0 {
		c.JSON(400, "Entry does not exist")
		return
	}

	if findResult.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user_entry": userEntry,
	})
}

func GetAllEntriesFromUser(c *gin.Context) {
	entryEmail := c.Param("email")
	var userEntries []models.UserEntry

	findResult := initializers.DB.Where("email = ?", entryEmail).Find(&userEntries)

	if findResult.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user_entries": userEntries,
	})
}

func UpdateEntry(c *gin.Context) {
	var updateEntryReq structs.UpdateEntryReq
	c.BindJSON(&updateEntryReq)

	updateResult := initializers.DB.Exec("UPDATE user_entries SET content = ? WHERE email = ? AND title = ?", updateEntryReq.Content, updateEntryReq.Email, updateEntryReq.Title)

	if updateResult.RowsAffected == 0 {
		c.JSON(400, "Entry does not exist")
		return
	}

	if updateResult.Error != nil {
		c.JSON(400, "Error updating the entry")
		return
	}

	var userEntry models.UserEntry
	entryResult := initializers.DB.Where("email = ? AND title = ?", updateEntryReq.Email, updateEntryReq.Title).Find(&userEntry)

	if entryResult.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user_entry": userEntry,
	})
}

func DeleteEntry(c *gin.Context) {
	entryTitle := c.Query("title")
	entryEmail := c.Query("email")

	deleteResult := initializers.DB.Where("title = ? AND email = ?", entryTitle, entryEmail).Delete(&models.UserEntry{})

	if deleteResult.RowsAffected == 0 {
		c.JSON(400, "Entry does not exist")
	}

	if deleteResult.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, "User entry deleted succesfully")
}
