package routes

import "github.com/gofiber/fiber/v2"

// NotFoundRoute func for describe 404 Error not found route
func NotFoundRoute(a *fiber.App) {
	// Register new special route
	a.Use(
		// Anonymous Function
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "sorry, route is not found",
			})
		},
	)
}
