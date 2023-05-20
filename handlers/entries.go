package handlers

import (
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
		return
	}

	c.JSON(200, "Entry created successfully")
}

func GetEntryHandler(c *gin.Context) {
	entryTitle := c.Query("title")
	entryEmail := c.Query("email")

	entry, err := services.GetEntry(entryEmail, entryTitle)
	if err != nil {
		c.Status(400)
		return
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
		return
	}

	c.JSON(200, gin.H{
		"user_entries": userEntries,
	})
}

func UpdateEntryHandler(c *gin.Context) {
	var body structs.UpdateEntryReq
	c.BindJSON(&body)

	err := services.UpdateEntry(body.Email, body.Password, body.Title, body.Content)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "Entry updated successfully")
}

func DeleteEntryHandler(c *gin.Context) {
	var body structs.DeleteEntryReq
	c.BindJSON(&body)

	err := services.DeleteEntry(body.Email, body.Password, body.Title)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "User entry deleted succesfully")
}
