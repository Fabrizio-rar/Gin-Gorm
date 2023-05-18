package main

import (
	"Gin-test/initializers"
	"Gin-test/routes"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func main() {
	routes.InitRoutes()
}
