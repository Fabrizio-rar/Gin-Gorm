package main

import (
	"Gin-test/initializers"
	"Gin-test/models"
	"fmt"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.User{},
		&models.UserEntry{},
	)
	fmt.Printf("Api started")
}
