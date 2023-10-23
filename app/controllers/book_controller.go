package controllers

import (
	"github.com/debidarmawan/debozero-core/platform/database"
	"github.com/gofiber/fiber/v2"
)

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
