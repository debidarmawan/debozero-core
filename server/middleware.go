package server

import (
	"log"
	"runtime/debug"

	"debozero-core/global"

	"github.com/gofiber/fiber/v2"
)

func PanicHandler(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			stack := string(debug.Stack())

			log.Println("[PANIC]")

			if e, ok := r.(error); ok {
				log.Println(e.Error())
			} else {
				log.Println(r)
			}

			log.Println(stack)

			err = c.Status(fiber.StatusInternalServerError).JSON(global.Response[interface{}]{
				Code:    fiber.StatusInternalServerError,
				Status:  global.ResponseStatus.FailedResponse,
				Data:    nil,
				Message: "Internal Server Error",
			})
		}
	}()

	return c.Next()
}
