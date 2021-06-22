package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"sample-go/app/config"
	"sample-go/app/server"
)


func main() {
	server.CurrentApp.InitServer()

	config.Connections.InitDb("postgresql://postgres@0.0.0.0:5432", "go_test")
	config.Connections.Amqp.InitAmqpChannel("amqp://localhost:5672/")

	err := server.CurrentApp.Run(":8080")
	if err != nil {
		fmt.Printf("Ups! Something went wrong %s\n", err)
	}

}
