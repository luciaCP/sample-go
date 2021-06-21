package tests

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sample-go/app"
	"sample-go/app/config"
	"testing"
)

func TestMain(m *testing.M)  {
	app.CurrentApp.InitServer()
	config.Connections.InitDb("postgresql://postgres@0.0.0.0:5432", "db_test")
	m.Run()
}

func setUp() {
	err := config.Connections.RestoreDb()
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

func TestIncrementOne(t *testing.T) {
	defer tearDown()
	setUp()

	writer := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/increment", nil)
	app.CurrentApp.Engine.ServeHTTP(writer, req)

	assert.Nil(t, err)
	assert.Equal(t, 201, writer.Code)

	var identifier, amount int
	config.Connections.GetConnection().QueryRow("SELECT id, incremental FROM go_test").Scan(&identifier, &amount)
	assert.Equal(t, 1, amount)

	var response map[string]int
	json.Unmarshal([]byte(writer.Body.String()), &response)
	assert.Equal(t, response["id"], identifier)
}

func TestIncrementTwo(t *testing.T) {
	defer tearDown()
	setUp()

	writer := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/increment", nil)
	app.CurrentApp.Engine.ServeHTTP(writer, req)

	assert.Nil(t, err)
	assert.Equal(t, 201, writer.Code)

	var firstResponse map[string]int
	json.Unmarshal([]byte(writer.Body.String()), &firstResponse)

	writer2 := httptest.NewRecorder()
	req, err = http.NewRequest("POST", "/increment", nil)
	app.CurrentApp.Engine.ServeHTTP(writer2, req)

	assert.Nil(t, err)
	assert.Equal(t, 201, writer2.Code)

	var secondResponse map[string]int
	json.Unmarshal([]byte(writer2.Body.String()), &secondResponse)


	selectionDd, _ := config.Connections.GetConnection().Query("SELECT id, incremental FROM go_test")
	for i:=0; selectionDd.Next() ; i++ {
		var identifier, amount int
		selectionDd.Scan(&identifier, &amount)

		assert.Equal(t, 1, amount)
		if i == 0 {
			assert.Equal(t, firstResponse["id"], identifier)
		} else {
			assert.Equal(t, secondResponse["id"], identifier)
		}
	}
}

func TestGetAll(t *testing.T) {
	defer tearDown()
	setUp()

	db := config.Connections.GetConnection()
	sqlStatement := `INSERT INTO go_test (incremental) VALUES (1) RETURNING id`
	var firstId, secondId int
	db.QueryRow(sqlStatement).Scan(&firstId)
	db.QueryRow(sqlStatement).Scan(&secondId)

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/increment", nil)
	app.CurrentApp.Engine.ServeHTTP(writer, req)

	assert.Equal(t, 200, writer.Code)

	var response [](map[string]int)
	json.Unmarshal([]byte(writer.Body.String()), &response)

	assert.Equal(t, firstId, response[0]["id"])
	assert.Equal(t, 1, response[0]["amount"])
	assert.Equal(t, secondId, response[1]["id"])
	assert.Equal(t, 1, response[1]["amount"])
}

