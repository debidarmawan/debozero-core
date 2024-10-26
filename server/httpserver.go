package server

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"debozero-core/constants"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const DefaultPort = 3598

func ServeHTTP(db *gorm.DB) {
	f := fiber.New()

	Routes(f, db)

	var port uint16
	if portEnv, ok := os.LookupEnv(constants.Port); ok {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			port = DefaultPort
		} else {
			port = uint16(portInt)
		}
	} else {
		port = DefaultPort
	}

	listenerPort := fmt.Sprintf(":%d", port)
	log.Fatal(f.Listen(listenerPort))
}
