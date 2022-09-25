package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/katarzynakawala/Library/internal/data"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title  string   `json:"title"`
		Author string   `json:"author"`
		Year   int32    `json:"year"`
		Pages  int32    `json:"pages"`
		Genres []string `json:"genres"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	book := data.Book{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Empuzjon",
		Author:    "Olga Tokarczuk",
		Pages:     400,
		Genres:    []string{"fiction", "horror"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
