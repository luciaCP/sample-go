package callbacks

import (
	"log"
	"sample-go/config"
)

func Interpret(id string) {
	log.Printf(" > Received message: %s\n", id)

	db := config.Connections.GetDbConnection()
	sqlStatement := `
		UPDATE go_test
		SET upgraded = $2
		WHERE id = $1;`
	result, err := db.Exec(sqlStatement, id, true)
	if r, _ := result.RowsAffected(); r == 0 || err != nil {
		log.Printf(" > Couldn't update model with id: %s\n", id)
		return
	}
	log.Printf(" > Updated model with id: %s\n", id)
}