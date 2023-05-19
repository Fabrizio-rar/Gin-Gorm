package handlers

import (
	"Gin-gorm/initializers"
	"Gin-gorm/models"
	"Gin-gorm/services"
	"Gin-gorm/structs"

	"github.com/gin-gonic/gin"
)

func CreateEntryHandler(c *gin.Context) {
	var body models.UserEntry
	err := c.BindJSON(&body)
	if err != nil {
		c.Status(400)
		return
	}

	err = services.CreateEntry(body.Email, body.Title, body.Content)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, "Entry created successfully")
}

func GetEntryHandler(c *gin.Context) {
	entryTitle := c.Query("title")
	entryEmail := c.Query("email")

	entry, err := services.GetEntry(entryEmail, entryTitle)
	if err != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"entry": entry,
	})
}

func GetAllEntriesFromUserHandler(c *gin.Context) {
	entryEmail := c.Param("email")

	userEntries, err := services.GetAllEntriesFromUser(entryEmail)
	if err != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"user_entries": userEntries,
	})
}

func UpdateEntryHandler(c *gin.Context) {
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

func DeleteEntryHandler(c *gin.Context) {
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
