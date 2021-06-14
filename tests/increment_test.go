package tests

import (
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

	var amount int
	app.CurrentApp.DbConnection.QueryRow("SELECT COUNT(*) FROM go_test").Scan(&amount)
	assert.Equal(t, 1, amount)
}

func TestIncrementTwo(t *testing.T) {
	defer tearDown()
	setUp()

	writer := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/increment", nil)
	app.CurrentApp.Engine.ServeHTTP(writer, req)

	assert.Nil(t, err)
	assert.Equal(t, 201, writer.Code)

	var amount int
	app.CurrentApp.DbConnection.QueryRow("SELECT COUNT(*) FROM go_test").Scan(&amount)
	assert.Equal(t, 1, amount)

	req, err = http.NewRequest("POST", "/increment", nil)
	app.CurrentApp.Engine.ServeHTTP(writer, req)

	assert.Nil(t, err)
	assert.Equal(t, 201, writer.Code)

	app.CurrentApp.DbConnection.QueryRow("SELECT COUNT(*) FROM go_test").Scan(&amount)
	assert.Equal(t, 2, amount)
}