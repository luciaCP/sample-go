package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func pingController(c *gin.Context) {
	c.String(200, "pong")
}


func main() {
	engine := gin.Default()

	engine.GET("/ping", pingController)

	err := engine.Run(":8080")
	if err == nil {
		fmt.Printf("Ups! Something go wrong %s\n", err)
	}
}
