package handlers

import (
	"Gin-gorm/models"
	"Gin-gorm/services"
	"Gin-gorm/structs"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	var requestBody models.User

	err := c.Bind(&requestBody)
	if err != nil {
		c.Status(400)
		return
	}

	err = services.CreateUser(requestBody.Email, requestBody.Password, requestBody.Name, requestBody.Gender)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "User created successfully")
}

func GetUserHandler(c *gin.Context) {
	userEmail := c.Param("email")

	user, err := services.GetUser(userEmail)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"user": structs.GetUserResp{
			Name:   user.Name,
			Gender: user.Gender,
			Email:  user.Email,
		},
	})
}

func GetAllUsersHandler(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func DeleteUserHandler(c *gin.Context) {
	var emailAndPassword structs.EmailAndPasswordReq
	c.BindJSON(&emailAndPassword)

	err := services.DeleteUser(emailAndPassword.Email, emailAndPassword.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "User deleted succesfully")
}
