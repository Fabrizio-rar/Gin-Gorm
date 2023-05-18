package routes

import (
	"Gin-test/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	router.POST("/create_user", controllers.CreateUser)
	router.GET("/get_all_users", controllers.GetAllUsers)
	router.GET("/get_user/:email", controllers.GetUser)
	router.POST("/delete_user/:email", controllers.DeleteUser)

	router.POST("/create_entry", controllers.CreateEntry)
	router.GET("/get_entry", controllers.GetEntry)
	router.GET("/get_user_entries/:email", controllers.GetAllEntriesFromUser)
	router.POST("/delete_entry", controllers.DeleteEntry)
	router.POST("/update_entry", controllers.UpdateEntry)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
