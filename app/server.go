package app

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"sample-go/app/controllers"
)


type App struct {
	Engine       *gin.Engine
}

func (app *App) InitServer() {
	engine := gin.Default()
	engine.GET("/ping", controllers.PingController)
	engine.POST("/increment", controllers.IncrementController)

	app.Engine = engine
}

func (app *App) Run(addr ...string) error {
	return app.Engine.Run(addr...)
}
