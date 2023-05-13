package controllers

import (
	"Gin-test/initializers"
	"Gin-test/models"
	"Gin-test/structs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var body models.User
	c.Bind(&body)

	var existingUser models.User

	exists := initializers.DB.Where("email = ?", body.Email).First(&existingUser)

	if exists.Error == gorm.ErrRecordNotFound {
		user := models.User{Name: body.Name, Gender: body.Gender, Email: body.Email, Password: body.Password}
		result := initializers.DB.Create(&user)

		if result.Error != nil {
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
	var emailParam structs.EmailReq
	c.BindJSON(&emailParam)
	var user models.User

	result := initializers.DB.Where("email = ?", emailParam.Email).Find(&user)

	if result.RowsAffected == 0 {
		c.JSON(400, "User does not exist")
		return
	}

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
	var emailParam structs.EmailReq
	c.BindJSON(&emailParam)

	userResult := initializers.DB.Where("email = ?", emailParam.Email).Delete(&models.User{})

	if userResult.RowsAffected == 0 {
		c.JSON(400, "User does not exist")
		return
	}

	if userResult.Error != nil {
		c.Status(400)
		return
	}

	entriesResult := initializers.DB.Where("email = ?", emailParam.Email).Delete(&models.UserEntry{})

	if entriesResult.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, "User deleted succesfully")
}
