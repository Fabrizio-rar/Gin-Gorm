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
	r.POST("/create_user", controllers.CreateUser)
	r.GET("/get_all_users", controllers.GetAllUsers)
	r.GET("/get_user/:id", controllers.GetUser)
	r.POST("/delete_user/:id", controllers.DeleteUser)
	r.POST("/create_entry/:id", controllers.CreateEntry)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
