package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sample-go/consumer/callbacks"
	"sample-go/consumer/config"
	"testing"
)

var mockAmqp = &MockAmqp{}

func TestMain(m *testing.M)  {
	config.Connections.InitDb("postgresql://postgres@0.0.0.0:5432", "db_test")

	m.Run()
}

func setUp() {
	err := config.Connections.RestoreDb()
	mockAmqp.Clean()

	if err != nil {
		fmt.Println(err)
	}
}

func tearDown() {
	err := config.Connections.FlushDb()
	if err != nil {
		fmt.Println(err)
	}
}

func TestInterpretMessageUpdatesDb(t *testing.T) {
	defer tearDown()
	setUp()

	db := config.Connections.GetDbConnection()
	sqlStatement := `INSERT INTO go_test (incremental) VALUES (1) RETURNING id, upgraded`
	var identifier float64
	var initialUpgraded bool
	db.QueryRow(sqlStatement).Scan(&identifier, &initialUpgraded)

	assert.False(t, initialUpgraded)

	callbacks.Interpret(fmt.Sprintf("%d", int(identifier)))

	var id, amount int
	var upgraded bool
	config.Connections.GetDbConnection().QueryRow("SELECT * FROM go_test WHERE id=$1", identifier).Scan(&id, &amount, &upgraded)
	assert.Equal(t, 1, amount)
	assert.Equal(t, true, upgraded)
}

func TestInterpretMessageWithNotFoundIdDoesNotUpdateDb(t *testing.T) {
	defer tearDown()
	setUp()

	defer tearDown()
	setUp()

	dummyId := "1"
	callbacks.Interpret(dummyId)

	var count int
	config.Connections.GetDbConnection().QueryRow("SELECT COUNT(*) FROM go_test").Scan(&count)
	assert.Equal(t, 0, count)
}