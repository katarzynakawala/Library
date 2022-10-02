package data

import (
	"database/sql"
	"time"

	"github.com/katarzynakawala/Library/internal/validator"
	"github.com/lib/pq"
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int32     `json:"year,omitempty"`
	Pages     Pages     `json:"pages,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateBook(v *validator.Validator, book *Book){
	v.Check(book.Title != "", "title", "must be provided") 
	v.Check(len(book.Title) <= 100, "title", "must not be more than 100 bytes long")

	v.Check(book.Author != "", "author", "must be provided")
	v.Check(len(book.Author) <= 100, "author", "must not be more than 100 bytes long")

	v.Check(book.Year != 0, "year", "must be provided")
	v.Check(book.Year >= 1000, "year", "must be greater than 1000")
	v.Check(book.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(book.Pages != 0, "pages", "must be provided")
	v.Check(book.Pages > 0, "pages", "must be a positive number") 

	v.Check(book.Genres != nil, "genres", "must be provided")
	v.Check(len(book.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(book.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(book.Genres), "genres", "must not contain duplicate values")
}

type BookModel struct {
	DB *sql.DB
}

func (b BookModel) Insert(book *Book) error {
	query := `
		INSERT INTO books (title, author, year, pages, genres)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version`
	
	args := []any{book.Title, book.Author, book.Year, book.Pages, pq.Array(book.Genres)}
	
	return b.DB.QueryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.Version)
}

func (b BookModel) Get(id int64) (*Book, error) {
	return nil, nil
}

func (b BookModel) Update(book *Book) error {
	return nil
}

func (b BookModel) Delete(id int64) error {
	return nil
}