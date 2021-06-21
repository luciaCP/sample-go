package config

import "database/sql"

type ConfigApp struct {
	db     *sql.DB
	dbName string
}

var Connections = ConfigApp{}

