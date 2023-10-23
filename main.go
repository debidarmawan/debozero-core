package main

import (
	"github.com/debidarmawan/debozero-core/pkg/configs"
	"github.com/debidarmawan/debozero-core/pkg/middleware"
	"github.com/debidarmawan/debozero-core/pkg/routes"
	"github.com/debidarmawan/debozero-core/pkg/utils"
	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {
	// Define fiber config
	config := configs.FiberConfig()

	// Define a new Fiber app with config
	app := fiber.New(config)

	// Middlewares
	middleware.FiberMiddleware(app)

	// Routes
	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)

	// Start server (with graceful shutdown)
	utils.StartServer(app)
}
