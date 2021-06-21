package services

import (
	"sample-go/app/config"
	"sample-go/app/models"
)

func CreateIncrease(amount int) int {
	db := config.Connections.GetConnection()
	sqlStatement := `
		INSERT INTO go_test (incremental)
		VALUES ($1)
		RETURNING id`
	var id int
	err := db.QueryRow(sqlStatement, amount).Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}

func GetAllIncrements() []models.Incremental {
	db := config.Connections.GetConnection()
	sqlStatement := `SELECT * FROM go_test`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	var increments []models.Incremental
	for rows.Next(){
		var oneIncrement models.Incremental
		rows.Scan(&oneIncrement.Id, &oneIncrement.Amount, &oneIncrement.Upgraded)

		increments = append(increments, oneIncrement)
	}

	return increments
}

func GetIncrement(id int) *models.Incremental {
	db := config.Connections.GetConnection()
	sqlStatement := `SELECT * FROM go_test WHERE id=$1`

	increment := new(models.Incremental)
	err := db.QueryRow(sqlStatement, id).Scan(
		&increment.Id, &increment.Amount, &increment.Upgraded,
	)
	if err != nil {
		return nil
	}

	return increment
}