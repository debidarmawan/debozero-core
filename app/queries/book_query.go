package queries

import (
	"github.com/debidarmawan/debozero-core/app/models"
	"github.com/jmoiron/sqlx"
)

// BookQueries struct for queries from book model
type BookQueries struct {
	*sqlx.DB
}

// GetBooks method for getting all books
func (q *BookQueries) GetBooks() ([]models.Book, error) {
	// Define books variable
	books := []models.Book{}

	// Define query string
	query := `SELECT * FROM books`

	// Send query to database
	err := q.Select(&books, query)
	if err != nil {
		// reutn empty object and error
		return books, err
	}

	// Return query result
	return books, err
}
