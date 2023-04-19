package main

import (
	"Gin-test/initializers"
	"Gin-test/models"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
