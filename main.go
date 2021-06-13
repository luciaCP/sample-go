package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sample-go/app"
	"sample-go/migrate"
)


func initDb() *sql.DB {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Chan!! Had to recover from panic on init DB - ", r)
		}
	}()
	db, err := sql.Open(
		"postgres",
		"postgresql://postgres@0.0.0.0:5432/go_test?sslmode=disable",
	)
	if err != nil || db.Ping() != nil {
		fmt.Printf("Ups! Couldn't connect to DB - %s\n\n", err)
	}

	err = migrate.Do(db)
	if err != nil {
		fmt.Printf("Ups! Couldn't migrate DB - %s\n\n", err)
	}
	return db
}


func main() {
	engine := app.InitServer()

	db := initDb()
	defer db.Close()

	err := engine.Run(":8080")
	if err != nil {
		fmt.Printf("Ups! Something went wrong %s\n", err)
	}
}
