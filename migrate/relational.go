package migrate

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"strings"
)

func getFilePath() string {
	workingDirectoryPath, _ := os.Getwd()
	absPath := strings.Split(workingDirectoryPath, "sample-go")
	return fmt.Sprintf("file://%s/sample-go/migrate/versions", absPath[0])
}

func getMigrateInstance(db *sql.DB) (*migrate.Migrate, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Chan!! Had to recover from panic in migration", r)
		}
	}()

	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	return migrate.NewWithDatabaseInstance(getFilePath(), "postgres", driver)
}

func Up(db *sql.DB) error {
	m, _ := getMigrateInstance(db)
	err := m.Up()
	if err != nil && err != migrate.ErrNoChange {return err}
	return nil
}

func Down(db *sql.DB) error {
	m, _ := getMigrateInstance(db)
	err := m.Down()
	if err != nil && err != migrate.ErrNoChange {return err}
	return nil
}