package handler

import (
	"github.com/debidarmawan/debozero-core/dto"
	"github.com/debidarmawan/debozero-core/global"
	"github.com/debidarmawan/debozero-core/helper"
	"github.com/debidarmawan/debozero-core/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(authUseCase usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (ah *AuthHandler) Routes(group fiber.Router) {
	group.Post("/auth/login", ah.Login)
}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	var request dto.Login

	if err := helper.ValidateBody(c, &request); err != nil {
		return err.ToResponse(c)
	}

	result, err := ah.authUseCase.Login(request)
	if err != nil {
		return err.ToResponse(c)
	}

	return global.CreateResponse(result, fiber.StatusOK, c)
}
