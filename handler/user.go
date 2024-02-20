package handler

import (
	"github.com/debidarmawan/debozero-core/dto"
	"github.com/debidarmawan/debozero-core/global"
	"github.com/debidarmawan/debozero-core/helper"
	"github.com/debidarmawan/debozero-core/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHanler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (uh *UserHandler) Routes(group fiber.Router) {
	group.Post("/users/register", uh.Register)
}

func (uh *UserHandler) Register(c *fiber.Ctx) error {
	var request dto.UserRegisterRequest

	if err := helper.ValidateBody(c, &request); err != nil {
		return err.ToResponse(c)
	}

	_, err := uh.userUseCase.Register(request)
	if err != nil {
		return err.ToResponse(c)
	}

	return global.CreateMessageResponse("Success", fiber.StatusOK, c)
}
