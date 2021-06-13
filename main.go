package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"sample-go/app"
)

func migrateDb(db *sql.DB) error {
	defer func() {
		fmt.Println("On defer")
		if r := recover(); r != nil {
			fmt.Println("Chan!! Had to recover from panic in migration", r)
		}
	}()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrate/versions", "postgres", driver)
	if err != nil {return err} else {fmt.Println("get migration instance - ", m)}

	err = m.Up()
	if err != nil {return err}
	return nil
}

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
	if err != nil {
		fmt.Printf("Ups! Couldn't connect to DB - %s\n\n", err)
	} else {
		err = db.Ping()
		if err != nil {
			panic("DB not connected!!!")
		}
	}

	err = migrateDb(db)
	if err != nil {
		fmt.Printf("Ups! Couldn't migrate DB - %s\n\n", err)
	} else {
		fmt.Println("Run migration finished!")
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
