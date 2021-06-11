package app

import "github.com/gin-gonic/gin"


func pingController(c *gin.Context) {
	c.String(200, "pong")
}


func InitServer() *gin.Engine {
	engine := gin.Default()

	engine.GET("/ping", pingController)
	return engine
}
