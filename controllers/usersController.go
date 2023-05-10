package controllers

import (
	"Gin-test/initializers"
	"Gin-test/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body models.User
	c.Bind(&body)

	user := models.User{Name: body.Name, Gender: body.Gender, Email: body.Email, Password: body.Password}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUser(c *gin.Context) {
	userID := c.Param("id")
	var user models.User

	result := initializers.DB.Find(&user, userID)

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

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	result := initializers.DB.Delete(&models.User{}, userID)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, "User deleted succesfully")
}
