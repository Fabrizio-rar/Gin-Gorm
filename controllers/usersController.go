package controllers

import (
	"Gin-test/initializers"
	"Gin-test/models"
	"Gin-test/structs"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body structs.User

	c.Bind(&body)

	user := models.User{Name: body.Name, Gender: body.Name, Email: body.Email, Password: body.Password}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}
