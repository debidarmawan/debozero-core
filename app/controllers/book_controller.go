package controllers

import (
	"time"

	"github.com/debidarmawan/debozero-core/app/models"
	"github.com/debidarmawan/debozero-core/pkg/utils"
	"github.com/debidarmawan/debozero-core/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetBooks func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]
func GetBooks(c *fiber.Ctx) error {
	// Create database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all books
	books, err := db.GetBooks()
	if err != nil {
		// Return, if books not found
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "books were not found",
			"count": 0,
			"books": nil,
		})
	}

	// Return status 200 OK
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(books),
		"books": books,
	})
}

// CreateBook func for creates a new book.
// @Description Create a new book.
// @Summary create a new book
// @Tags Book
// @Accept json
// @Produce json
// @Param body body models.Book true "Body"
// @Success 200 {object} models.Book
// @Security ApiKeyAuth
// @Router /v1/book [post]
func CreateBook(c *fiber.Ctx) error {
	// Get Current Time
	now := time.Now().Unix()

	// Get claims from JWT
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": nil,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book
	expires := claims.Expires

	// Checking, if now time greater than expiration time from JWT
	if now > expires {
		// Return status 401 and unauthorized error message
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, expired token",
		})
	}

	// Create new book struct
	book := &models.Book{}

	// Check if received JSON data is valid
	if err := c.BodyParser(book); err != nil {
		// Return status 400 and error message
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a book model
	validate := utils.NewValidator()

	// Set initialized default data for book
	book.ID = uuid.New()
	book.CreatedAt = time.Now()
	book.BookStatus = 1 // 0 == draft, 1 == active

	// Validate book fields
	if err := validate.Struct(book); err != nil {
		// Return, if some fields are not valid
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create book by given ID
	if err := db.CreateBook(book); err != nil {
		// Return status 500 and error message
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"book":  book,
	})
}
