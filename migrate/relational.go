package migrate

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Do(db *sql.DB) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Chan!! Had to recover from panic in migration", r)
		}
	}()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrate/versions", "postgres", driver)
	if err != nil {return err}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {return err}
	return nil
}

