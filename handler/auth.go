package handler

import (
	"net/http"

	"github.com/debidarmawan/debozero-core/dto"
	"github.com/debidarmawan/debozero-core/global"
	"github.com/debidarmawan/debozero-core/helper"
	"github.com/debidarmawan/debozero-core/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type AuthHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(authUseCase usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (ah *AuthHandler) Routes(group fiber.Router) {
	group.Post("/auth/login", ah.Login)
	group.Post("/auth/logout", ah.Logout)
	group.Get("/auth/verify", ah.Verify)
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

func (ah *AuthHandler) Logout(c *fiber.Ctx) error {
	var request http.Request

	fasthttpadaptor.ConvertRequest(c.Context(), &request, false)

	err := ah.authUseCase.Logout(dto.Logout{Request: &request})
	if err != nil {
		return err.ToResponse(c)
	}

	return global.MessageResponse("Success", fiber.StatusOK, c)
}

func (ah *AuthHandler) Verify(c *fiber.Ctx) error {
	var header dto.VerifyHeader

	if err := helper.ValidateHeader(c, &header); err != nil {
		return err.ToResponse(c)
	}

	var httpRequest http.Request
	fasthttpadaptor.ConvertRequest(c.Context(), &httpRequest, false)

	request := dto.Verify{
		Request: &httpRequest,
		Path:    header.Path,
		Method:  header.Method,
	}

	result, err := ah.authUseCase.Verify(request)

	if err != nil {
		return err.ToResponse(c)
	}

	return global.CreateResponse(result, fiber.StatusOK, c)
}
