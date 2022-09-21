package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_application_healthcheckHandler(t *testing.T) {
	app := applicationInstance(8080, "development")

	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/v1/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	app.healthcheckHandler(rr, r) 

	rs := rr.Result()

	assert.Equal(t, http.StatusOK, rs.StatusCode)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := "{\"environment\":\"development\",\"status\":\"available\",\"version\":\"1.0.0\"}\n"
	assert.Equal(t, rs.Header.Get("Content-Type"), "application/json")
	assert.Equal(t, expected, string(body))
}
