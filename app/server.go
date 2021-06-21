package app

import (
	"github.com/gin-gonic/gin"
	"sample-go/app/controllers"
)


type App struct {
	Engine       *gin.Engine
}

func (app *App) InitServer() {
	engine := gin.Default()
	engine.GET("/ping", controllers.PingController)

	engine.POST("/increment", controllers.CreateIncrement)
	engine.GET("/increment", controllers.GetAllIncrements)

	engine.GET("/increment/:id", controllers.GetIncrement)

	app.Engine = engine
}

func (app *App) Run(addr ...string) error {
	return app.Engine.Run(addr...)
}
