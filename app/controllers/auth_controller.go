package controllers

import (
	"github.com/debidarmawan/debozero-core/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// GetNewAccessToken method for create a new access token
// @Description Create a new access token
// @Summary create a new access token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/token/new [get]
func GetNewAccessToken(c *fiber.Ctx) error {
	// Generate new access token
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		// Return status 500 and token generator error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}
