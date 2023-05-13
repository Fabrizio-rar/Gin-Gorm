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

	exists := initializers.DB.Where("title = ? AND email = ?", body.Title, body.Email).First(&existingEntry)

	if exists.Error == gorm.ErrRecordNotFound {
		userEntry := models.UserEntry{Title: body.Title, Content: body.Content, Email: body.Email}
		result := initializers.DB.Create(&userEntry)

		if result.Error != nil {
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
	var titleParam structs.TitleReq
	c.BindJSON(&titleParam)
	var userEntry models.UserEntry

	result := initializers.DB.Where("title = ?", titleParam.Title).Find(&userEntry)

	if result.RowsAffected == 0 {
		c.JSON(400, "Entry does not exist")
		return
	}

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user_entry": userEntry,
	})
}

func GetAllEntriesFromUser(c *gin.Context) {
	var emailParam structs.EmailReq
	c.BindJSON(&emailParam)

	var userEntries []models.UserEntry

	result := initializers.DB.Where("email = ?", emailParam.Email).Find(&userEntries)

	if result.Error != nil {
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

	result := initializers.DB.Exec("UPDATE user_entries SET content = ? WHERE email = ? AND title = ?", updateEntryReq.Content, updateEntryReq.Email, updateEntryReq.Title)

	if result.RowsAffected == 0 {
		c.JSON(400, "Entry does not exist")
		return
	}

	if result.Error != nil {
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
	var titleParam structs.TitleReq
	c.BindJSON(&titleParam)

	result := initializers.DB.Where("title = ?", titleParam.Title).Delete(&models.UserEntry{})

	if result.RowsAffected == 0 {
		c.JSON(400, "Entry does not exist")
	}

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, "User entry deleted succesfully")
}
