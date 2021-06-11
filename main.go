package main

import (
	"fmt"
	"sample-go/app"
)

func main() {
	engine := app.InitServer()
	err := engine.Run(":8080")
	if err == nil {
		fmt.Printf("Ups! Something go wrong %s\n", err)
	}
}
