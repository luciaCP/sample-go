package services

import (
	"fmt"
	"sample-go/app/config"
	"sample-go/app/models"
)

func CreateIncrease() int {
	db := config.Connections.GetDbConnection()
	sqlStatement := `
		INSERT INTO go_test (incremental)
		VALUES ($1)
		RETURNING id`
	var id int
	err := db.QueryRow(sqlStatement, 1).Scan(&id)
	if err != nil {
		panic(err)
	}

	err = config.Connections.Amqp.Publish(config.NotifyQueue, fmt.Sprintf("%d", id))
	if err != nil {
		fmt.Println("Error when send to queue " + err.Error())
	}
	return id
}

func GetAllIncrements() []models.Incremental {
	db := config.Connections.GetDbConnection()
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
	db := config.Connections.GetDbConnection()
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