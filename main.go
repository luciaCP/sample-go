package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"sample-go/app"
)

func main() {
	myApp := app.App{}
	myApp.InitServer()

	myApp.InitDb("postgresql://postgres@0.0.0.0:5432/go_test?sslmode=disable")

	err := myApp.Run(":8080")
	if err != nil {
		fmt.Printf("Ups! Something went wrong %s\n", err)
	}
}
