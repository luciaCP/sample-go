package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sample-go/migrate"
)

type DbConnection interface {
	InitDb(uriDB string, dbName string)
	GetConnection() *sql.DB
	FlushDb() error
	RestoreDb() error
	Close() error
}

func (config *ConfigApp) InitDb(uriDB string, dbName string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Chan!! Had to recover from panic on init DB - ", r)
		}
	}()
	if dbName == "" {
		dbName = "go_test"
	}
	db, err := sql.Open(
		"postgres",
		uriDB + "/" + dbName + "?sslmode=disable",
	)

	if err != nil || db.Ping() != nil {
		fmt.Printf("Ups! Couldn't connect to DB - %s\n\n", err)
	}

	if err:=migrate.Up(db) ; err != nil {
		fmt.Printf("Ups! Couldn't migrate DB - %s\n\n", err)
	}

	config.db = db
	config.dbName = dbName
}

func (config *ConfigApp) GetConnection() *sql.DB {
	return config.db
}

func (config *ConfigApp) FlushDb() error {
	if err := config.db.Ping() ; err != nil {
		fmt.Printf("Ups! Couldn't connect to DB on Flush - %s\n\n", err)
		return err
	}

	return migrate.Down(config.db)
}

func (config *ConfigApp) RestoreDb() error {
	if err := config.db.Ping() ; err != nil {
		fmt.Printf("Ups! Couldn't connect to DB on Restore - %s\n\n", err)
		return err
	}

	return migrate.Up(config.db)
}

func (config *ConfigApp) Close() error {
	migrate.Down(config.db)
	return config.db.Close()
}