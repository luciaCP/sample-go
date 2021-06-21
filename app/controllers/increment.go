package controllers

import (
	"github.com/gin-gonic/gin"
	"sample-go/app/services"
	"strconv"
)

type IncrementsController interface {
	Create(*gin.Context)
	Get(*gin.Context)
}

func CreateIncrement(c *gin.Context) {
	createdId := services.CreateIncrease()
	c.JSON(201, gin.H{"id": createdId})
}

func GetAllIncrements(c *gin.Context) {
	values := services.GetAllIncrements()
	c.JSON(200, values)
}

func GetIncrement(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid identifier"})
		return
	}
	
	value := services.GetIncrement(id)
	if value == nil {
		c.JSON(200, gin.H{})
		return
	}
	c.JSON(200, value)
}
