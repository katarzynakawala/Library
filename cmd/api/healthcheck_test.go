package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_application_healthcheckHandler(t *testing.T) {
	var cfg config
	cfg.port = 8080
	cfg.env = "development"

	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/v1/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	app.healthcheckHandler(rr, r) 

	rs := rr.Result()

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, string(body),"status: available")
	assert.Contains(t, string(body),"environment: development")
	assert.Contains(t, string(body),"version: 1.0.0")
}
