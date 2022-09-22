package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/katarzynakawala/Library/internal/data"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new book")
}

func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
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

	err = app.writeJSON(w, http.StatusOK, book, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
