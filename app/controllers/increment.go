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
	id, _ := strconv.Atoi(c.Param("id"))
	value := services.GetIncrement(id)
	c.JSON(200, value)
}
