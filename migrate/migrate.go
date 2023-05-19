package main

import (
	"Gin-gorm/initializers"
	"Gin-gorm/models"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.UserEntry{})
}
