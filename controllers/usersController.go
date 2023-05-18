package controllers

import (
	"Gin-test/initializers"
	"Gin-test/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var requestBody models.User
	c.Bind(&requestBody)
	var existingUser models.User

	userExists := initializers.DB.Where("email = ?", requestBody.Email).First(&existingUser)

	if userExists.Error == gorm.ErrRecordNotFound {
		user := models.User{Name: requestBody.Name, Gender: requestBody.Gender, Email: requestBody.Email, Password: requestBody.Password}
		createResult := initializers.DB.Create(&user)

		if createResult.Error != nil {
			c.Status(400)
			return
		}

		c.JSON(200, gin.H{
			"message": "User created successfully",
		})
		return
	}

	c.JSON(400, gin.H{
		"message": "User already exists",
	})
}

func GetUser(c *gin.Context) {
	userEmail := c.Param("email")
	var user models.User

	findResult := initializers.DB.Where("email = ?", userEmail).Find(&user)

	if findResult.RowsAffected == 0 {
		c.JSON(400, "User does not exist")
		return
	}

	if findResult.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	findResult := initializers.DB.Find(&users)

	if findResult.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func DeleteUser(c *gin.Context) {
	userEmail := c.Param("email")

	userResult := initializers.DB.Where("email = ?", userEmail).Delete(&models.User{})

	if userResult.RowsAffected == 0 {
		c.JSON(400, "User does not exist")
		return
	}

	if userResult.Error != nil {
		c.Status(400)
		return
	}

	entriesResult := initializers.DB.Where("email = ?", userEmail).Delete(&models.UserEntry{})

	if entriesResult.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, "User deleted succesfully")
}
