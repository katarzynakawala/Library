package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_application_showBookHandler(t *testing.T) {
	app := applicationInstance(4000, "localhost")

	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/v1/books/", nil)
	if err != nil {
		t.Fatal(err)
	}

	app.showBookHandler(rr, r)

	rs := rr.Result()

	assert.Equal(t, http.StatusNotFound, rs.StatusCode)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, string(body), "the requested resource could not be found")
}

func Test_application_createBookHandler(t *testing.T) {
	app := applicationInstance(4000, "localhost")

	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodPost, "/v1/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	app.createBookHandler(rr, r)

	rs := rr.Result()

	assert.Equal(t, http.StatusOK, rs.StatusCode)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, string(body), "create a new book\n")
}
