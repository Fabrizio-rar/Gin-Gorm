package controllers

import (
	"Gin-test/initializers"
	"Gin-test/models"
	"Gin-test/structs"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body models.User
	c.Bind(&body)

	user := models.User{Name: body.Name, Gender: body.Gender, Email: body.Email, Password: body.Password}

	result := initializers.DB.FirstOrCreate(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	if result.RowsAffected == 1 {
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

	result := initializers.DB.Where("email = ?", emailParam.Email).Delete(&models.User{})

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, "User deleted succesfully")
}
