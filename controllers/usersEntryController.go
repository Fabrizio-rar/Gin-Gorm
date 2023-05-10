package controllers

import (
	"Gin-test/initializers"
	"Gin-test/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEntry(c *gin.Context) {
	var body models.UserEntry
	c.Bind(&body)

	userId := c.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.Status(400)
		return
	}

	userEntry := models.UserEntry{Title: body.Title, Content: body.Content, UserID: userIdInt}

	result := initializers.DB.Create(&userEntry)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user_entry": userEntry,
	})
}
