package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"sample-go/app/services"
	"sample-go/migrate"
)


func pingController(c *gin.Context) {
	c.String(200, "pong")
}

func incrementController(c *gin.Context) {
	createdId := services.Increase(CurrentApp.DbConnection)
	c.JSON(201, gin.H{"id": createdId})
}

type App struct {
	DbConnection *sql.DB
	Engine       *gin.Engine
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

	if err:=migrate.Up(db) ; err != nil {
		fmt.Printf("Ups! Couldn't migrate DB - %s\n\n", err)
	}

	app.DbConnection = db
}

func (app *App) Run(addr ...string) error {
	return app.Engine.Run(addr...)
}

func (app *App) FlushDb() error {
	if err := app.DbConnection.Ping() ; err != nil {
		fmt.Printf("Ups! Couldn't connect to DB on Flush - %s\n\n", err)
		return err
	}

	return migrate.Down(app.DbConnection)
}

func (app *App) RestoreDb() error {
	if err := app.DbConnection.Ping() ; err != nil {
		fmt.Printf("Ups! Couldn't connect to DB on Restore - %s\n\n", err)
		return err
	}

	return migrate.Up(app.DbConnection)
}

func (app *App) Close() error {
	migrate.Down(app.DbConnection)
	return app.DbConnection.Close()
}