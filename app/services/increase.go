package services

import (
	"database/sql"
)

type incrementalDto struct {
	Id int
	Incremental int
}

func Increase(db *sql.DB) int {
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

