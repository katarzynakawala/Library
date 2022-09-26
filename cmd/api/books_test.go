package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func Test_application_createBookHandlerHappyPath(t *testing.T) {
	app := applicationInstance(4000, "localhost")

	rr := httptest.NewRecorder()

	requestBody := "{\"title\":\"book\",\"author\":\"author\",\"year\":2016,\"pages\":107, \"genres\":[\"animation\",\"adventure\"]}"
	reader := strings.NewReader(requestBody)

	r, err := http.NewRequest(http.MethodPost, "/v1/books", reader)
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

	assert.Equal(t, "{Title:book Author:author Year:2016 Pages:107 Genres:[animation adventure]}\n", string(body))
}


func Test_application_createBookHandlerBadJsonFormat(t *testing.T) {
	app := applicationInstance(4000, "localhost")

	rr := httptest.NewRecorder()

	requestBody := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><note><to>Alex</to></note>"
	reader := strings.NewReader(requestBody)

	r, err := http.NewRequest(http.MethodPost, "/v1/books", reader)
	if err != nil {
		t.Fatal(err)
	}

	app.createBookHandler(rr, r)

	rs := rr.Result()

	assert.Equal(t, http.StatusBadRequest, rs.StatusCode)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "{\n\t\"error\": \"body contains badly-formed JSON (at character 1\"\n}\n", string(body))
}