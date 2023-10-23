package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

// StartServer func for starting server with a graceful shutdown
func StartServer(a *fiber.App) {
	// Create channel for idle connections
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals
		<-sigint

		// Received an interrupt signal, shutdown
		if err := a.Shutdown(); err != nil {
			log.Printf("Ooops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Ooops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
