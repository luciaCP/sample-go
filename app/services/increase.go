package services

import (
	"sample-go/app/config"
)

type incrementalDto struct {
	Id int
	Incremental int
}

func Increase() int {
	db := config.Connections.GetConnection()
	sqlStatement := `
		INSERT INTO go_test (incremental)
		VALUES ($1)
		RETURNING id`
	var id int
	err := db.QueryRow(sqlStatement, 1).Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}

