package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"sample-go/app"
)


func main() {
	app.CurrentApp.InitServer()

	app.CurrentApp.InitDb("postgresql://postgres@0.0.0.0:5432/go_test?sslmode=disable")

	err := app.CurrentApp.Run(":8080")
	if err != nil {
		fmt.Printf("Ups! Something went wrong %s\n", err)
	}

}
