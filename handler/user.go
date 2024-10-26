package handler

import (
	"debozero-core/dto"
	"debozero-core/global"
	"debozero-core/helper"
	"debozero-core/usecase"

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

// Register godoc
//
//	@Summary	Create an account
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		payload	body		dto.UserRegisterRequest	true	"User data"
//	@Success	200		{object}	global.Response[dto.Message]
//	@Router		/users/register [post]
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
