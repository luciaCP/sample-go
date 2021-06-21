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

type asd struct {
	amount int `bson:"amount"`
}

func CreateIncrement(c *gin.Context) {
	var body asd
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	createdId := services.CreateIncrease(body.amount)
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
