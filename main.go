package main

import (
	"github.com/debidarmawan/debozero-core/pkg/configs"
	"github.com/debidarmawan/debozero-core/pkg/middleware"
	"github.com/debidarmawan/debozero-core/pkg/routes"
	"github.com/debidarmawan/debozero-core/pkg/utils"
	"github.com/gofiber/fiber/v2"

	_ "github.com/debidarmawan/debozero-core/docs"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title DeboZero Core API
// @version 1.0
// @description API Docs of DeboZero Core API Service
// @contact.name Debi Darmawan
// @contact.email debidarmawan1998@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	// Define fiber config
	config := configs.FiberConfig()

	// Define a new Fiber app with config
	app := fiber.New(config)

	// Middlewares
	middleware.FiberMiddleware(app)

	// Routes
	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	// Start server (with graceful shutdown)
	utils.StartServer(app)
}
