package main

import (
	"web-api/initializers"
	"web-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Role{}, &models.User{}, &models.Hobi{})
}
