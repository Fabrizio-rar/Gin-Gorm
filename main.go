package main

import (
	"Gin-test/controllers"
	"Gin-test/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetAllUsers)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
