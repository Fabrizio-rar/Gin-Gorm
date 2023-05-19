package main

import (
	"Gin-gorm/initializers"
	"Gin-gorm/routes"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func main() {
	routes.InitRoutes()
}
