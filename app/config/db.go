package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sample-go/migrate"
)

type ConfigApp struct {
	db *sql.DB
}

var Connections = ConfigApp{}

func (config *ConfigApp) InitDb(uriDB string) {
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

	config.db = db
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