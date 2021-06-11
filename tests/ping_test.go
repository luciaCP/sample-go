package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sample-go/app"
	"testing"
)

func TestPing(t *testing.T)  {
	router := app.InitServer()

	writer := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(writer, req)

	assert.Nil(t, err)
	assert.Equal(t, 200, writer.Code)
	assert.Equal(t, "pong", writer.Body.String())
}
