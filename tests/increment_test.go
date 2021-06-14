package tests

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sample-go/app"
	"testing"
)

func TestMain(m *testing.M)  {
	app.CurrentApp.InitServer()
	app.CurrentApp.InitDb("postgresql://postgres@0.0.0.0:5432/go_test?sslmode=disable")
	m.Run()
}

func setUp() {
	err := app.CurrentApp.RestoreDb()
	if err != nil {
		fmt.Println(err)
	}
}

func tearDown() {
	err := app.CurrentApp.FlushDb()
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
	app.CurrentApp.DbConnection.QueryRow("SELECT id, incremental FROM go_test").Scan(&identifier, &amount)
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


	selectionDd, _ := app.CurrentApp.DbConnection.Query("SELECT id, incremental FROM go_test")
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