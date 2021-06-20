package controllers

import (
	"github.com/gin-gonic/gin"
	"sample-go/app/services"
)

func IncrementController(c *gin.Context) {
	createdId := services.Increase()
	c.JSON(201, gin.H{"id": createdId})
}
