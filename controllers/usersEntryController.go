package controllers

import (
	"Gin-test/initializers"
	"Gin-test/models"
	"Gin-test/structs"

	"github.com/gin-gonic/gin"
)

func CreateEntry(c *gin.Context) {
	var body models.UserEntry
	err := c.BindJSON(&body)
	if err != nil {
		c.Status(400)
		return
	}

	userEntry := models.UserEntry{Title: body.Title, Content: body.Content, Email: body.Email}

	result := initializers.DB.FirstOrCreate(&userEntry)

	if result.Error != nil {
		c.Status(400)
		return
	}

	if result.RowsAffected == 1 {
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
	type UserEntryReq struct {
		Title string `json:"title" binding:"required"`
	}

	var dataEntry UserEntryReq
	c.BindJSON(&dataEntry)

	var userEntry models.UserEntry
	result := initializers.DB.Find(&userEntry, "title = ?", dataEntry.Title)

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

func DeleteEntry(c *gin.Context) {
	type UserEntryDelReq struct {
		Title string `json:"title" binding:"required"`
	}

	var dataEntry UserEntryDelReq
	c.BindJSON(&dataEntry)

	result := initializers.DB.Where("title = ?", dataEntry.Title).Delete(&models.UserEntry{})

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, "User entry deleted succesfully")
}
