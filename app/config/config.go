package config

import (
	"database/sql"
)

type ConfigApp struct {
	db   *sql.DB
	Amqp AmqpConnection
}

var Connections = ConfigApp{}

func (config *ConfigApp) CloseConfig() {
	config.Close()
	config.CloseAmqp()
}
