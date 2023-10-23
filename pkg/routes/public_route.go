package routes

import (
	"github.com/debidarmawan/debozero-core/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes
func PublicRoutes(a *fiber.App) {
	// Router group
	route := a.Group("/api/v1")

	route.Get("/books", controllers.GetBooks)
}
