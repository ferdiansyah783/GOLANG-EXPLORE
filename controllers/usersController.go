package controllers

import (
	"net/http"
	"strconv"
	"web-api/initializers"
	"web-api/models"

	"github.com/gin-gonic/gin"
)

func UsersCreate(c *gin.Context) {
	var body struct {
		Name string
		Role string
		Hobi string
	}

	c.Bind(&body)

	var user models.User

	role := models.Role{Role: body.Role}

	newUser := models.User{Name: body.Name, Role: role, Hobies: []models.Hobi{
		{Hobi: "belajar"},
		{Hobi: "Membaca"},
		{Hobi: "Catur"},
	}}

	if err := initializers.DB.Where(&models.User{Name: body.Name}).First(&user).Error; err != nil {

		result := initializers.DB.Create(&newUser)

		if result.Error != nil {
			c.JSON(400, gin.H{
				"message": "bad request",
			})
			return
		}

		c.JSON(201, gin.H{
			"data":       newUser,
			"statusCode": http.StatusCreated,
			"message":    "success",
		})
	}

	c.JSON(400, gin.H{
		"message": "Name sudah digunakan",
	})

}

func UsersFind(c *gin.Context) {
	var users []models.User
	initializers.DB.Model(users).Preload("Role").Preload("Hobies").Find(&users)

	c.JSON(200, gin.H{
		"data":       users,
		"statusCode": http.StatusOK,
		"message":    "success",
	})
}

func UsersFindOne(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	c.JSON(200, gin.H{
		"data":       user,
		"statusCode": http.StatusOK,
		"message":    "success",
	})
}

func UsersUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name string
	}

	c.Bind(&body)

	var user models.User
	initializers.DB.First(&user, id)

	initializers.DB.Model(&user).Updates(models.User{
		Name: body.Name,
	})

	c.JSON(200, gin.H{
		"data":       user,
		"statusCode": http.StatusOK,
		"message":    "success",
	})
}

func UsersDelete(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	initializers.DB.Delete(&models.User{}, id)

	c.Status(200)
}

func UserFilterByName(c *gin.Context) {
	name := c.Query("name")

	var user models.User
	initializers.DB.Where("Name = ?", name).First(&user)

	c.JSON(200, gin.H{
		"data":       user,
		"statusCode": http.StatusOK,
		"message":    "success",
	})
}

func UsersPagination(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	var users []models.User
	initializers.DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&users)

	c.JSON(200, gin.H{
		"data":       users,
		"statusCode": http.StatusOK,
		"message":    "success",
	})
}
