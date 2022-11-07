package main

import (
	"web-api/controllers"
	"web-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.POST("/users", controllers.UsersCreate)
	router.GET("/users", controllers.UsersFind)
	router.GET("/users/:id", controllers.UsersFindOne)
	router.PUT("/users/:id", controllers.UsersUpdate)
	router.DELETE("/users/:id", controllers.UsersDelete)
	router.GET("/users/user", controllers.UserFilterByName)
	router.GET("/users/pagination", controllers.UsersPagination)

	router.Run(":8888")
}
