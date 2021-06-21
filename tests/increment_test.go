package tests

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sample-go/app"
	"sample-go/app/config"
	"strconv"
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
	var upgraded bool
	config.Connections.GetConnection().QueryRow("SELECT * FROM go_test").Scan(&identifier, &amount, &upgraded)
	assert.Equal(t, 1, amount)
	assert.Equal(t, false, upgraded)

	var response map[string]int
	json.Unmarshal([]byte(writer.Body.String()), &response)
	assert.Equal(t, identifier, response["id"])
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


	selectionDd, _ := config.Connections.GetConnection().Query("SELECT * FROM go_test")
	for i:=0; selectionDd.Next() ; i++ {
		var identifier, amount int
		var upgraded bool
		selectionDd.Scan(&identifier, &amount, &upgraded)

		assert.Equal(t, 1, amount)
		assert.Equal(t, false, upgraded)
		if i == 0 {
			assert.Equal(t, identifier, firstResponse["id"])
		} else {
			assert.Equal(t, identifier, secondResponse["id"])
		}
	}
}

func TestGetAll(t *testing.T) {
	defer tearDown()
	setUp()

	db := config.Connections.GetConnection()
	sqlStatement := `INSERT INTO go_test (incremental) VALUES (1) RETURNING id`
	var firstId, secondId float64
	db.QueryRow(sqlStatement).Scan(&firstId)
	db.QueryRow(sqlStatement).Scan(&secondId)

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/increment", nil)
	app.CurrentApp.Engine.ServeHTTP(writer, req)

	assert.Equal(t, 200, writer.Code)

	var response [](map[string]interface{})
	json.Unmarshal([]byte(writer.Body.String()), &response)

	assert.Equal(t, firstId, response[0]["id"])
	assert.Equal(t, float64(1), response[0]["amount"].(float64))
	assert.Equal(t, false, response[0]["upgraded"].(bool))

	assert.Equal(t, secondId, response[1]["id"])
	assert.Equal(t, float64(1), response[1]["amount"].(float64))
	assert.Equal(t, false, response[1]["upgraded"].(bool))
}

func TestGetIncrementByIdReturnsIncrement(t *testing.T) {
	defer tearDown()
	setUp()

	db := config.Connections.GetConnection()
	sqlStatement := `INSERT INTO go_test (incremental) VALUES (1) RETURNING id`
	var firstId float64
	db.QueryRow(sqlStatement).Scan(&firstId)

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/increment/" +  strconv.Itoa(int(firstId)), nil)
	app.CurrentApp.Engine.ServeHTTP(writer, req)

	assert.Equal(t, 200, writer.Code)

	var response map[string]interface{}
	json.Unmarshal([]byte(writer.Body.String()), &response)

	assert.Equal(t, firstId, response["id"])
	assert.Equal(t, float64(1), response["amount"].(float64))
	assert.Equal(t, false, response["upgraded"].(bool))
}

func TestGetIncrementWithInvalidIdReturnsBadRequest(t *testing.T) {
	defer tearDown()
	setUp()

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/increment/missing", nil)
	app.CurrentApp.Engine.ServeHTTP(writer, req)

	assert.Equal(t, 400, writer.Code)

	var response map[string]string
	expected := map[string]string{"message":"Invalid identifier"}
	json.Unmarshal([]byte(writer.Body.String()), &response)

	assert.Equal(t, expected, response)
}

func TestGetIncrementByMissingIdReturnsEmpty(t *testing.T) {
	defer tearDown()
	setUp()

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/increment/1", nil)
	app.CurrentApp.Engine.ServeHTTP(writer, req)

	assert.Equal(t, 200, writer.Code)

	var response map[string]int
	expected := map[string]int{}
	json.Unmarshal([]byte(writer.Body.String()), &response)

	assert.Equal(t, expected, response)
}