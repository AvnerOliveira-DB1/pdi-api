package controllers

import (
	"strconv"

	"github.com/avner.oliveira/pdi-api/models"
	"github.com/gin-gonic/gin"
)

func FindById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
	}

	if idInt == 3 {
		c.JSON(200, gin.H{
			"id":    3,
			"name":  "Test DB1",
			"email": "test@db1.com.br",
			"age":   30,
		})

	}

}

func Createuser(c *gin.Context) {
	var u models.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	println("ID: ", u.Id)
	println("Name: ", u.Name)
	println("Email: ", u.Email)
	println("Age: ", u.Age)

}
