package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"sample-go/migrate"
)


func pingController(c *gin.Context) {
	c.String(200, "pong")
}

func incrementController(c *gin.Context) {

}

type App struct {
	dbConnection *sql.DB
	Engine *gin.Engine
}

func (app *App) InitServer() {
	engine := gin.Default()

	engine.GET("/ping", pingController)
	engine.POST("/increment", incrementController)

	app.Engine = engine
}

func (app *App) InitDb(uriDB string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Chan!! Had to recover from panic on init DB - ", r)
		}
	}()
	db, err := sql.Open(
		"postgres",
		uriDB,
	)

	if err != nil || db.Ping() != nil {
		fmt.Printf("Ups! Couldn't connect to DB - %s\n\n", err)
	}
	defer db.Close()

	if err:=migrate.Up(db) ; err != nil {
		fmt.Printf("Ups! Couldn't migrate DB - %s\n\n", err)
	}

	app.dbConnection = db
}

func (app *App) Run(addr ...string) error {
	return app.Engine.Run(addr...)
}

func (app *App) FlushDb() error {
	return migrate.Down(app.dbConnection)
}

func (app *App) Close() error {
	migrate.Down(app.dbConnection)
	return app.dbConnection.Close()
}