package routes

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/debidarmawan/debozero-core/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes
func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook)
}
