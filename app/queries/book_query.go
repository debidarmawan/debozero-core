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

// CreateBook method for creating book by given Book object
func (q *BookQueries) CreateBook(b *models.Book) error {
	// Define query string
	query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Send query to database
	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserID, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		// Return error
		return err
	}

	return nil
}
