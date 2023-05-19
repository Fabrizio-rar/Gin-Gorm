package routes

import (
	"Gin-gorm/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	router.POST("/create_user", handlers.CreateUserHandler)
	router.GET("/get_all_users", handlers.GetAllUsersHandler)
	router.GET("/get_user/:email", handlers.GetUserHandler)
	router.POST("/delete_user", handlers.DeleteUserHandler)

	router.POST("/create_entry", handlers.CreateEntryHandler)
	router.GET("/get_entry", handlers.GetEntryHandler)
	router.GET("/get_user_entries/:email", handlers.GetAllEntriesFromUserHandler)
	router.POST("/delete_entry", handlers.DeleteEntryHandler)
	router.POST("/update_entry", handlers.UpdateEntryHandler)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
