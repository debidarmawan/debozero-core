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
	group.Post("/auth/refresh", ah.Refresh)
}

// Login godoc
//
//	@Summary	Login to get access token
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		payload	body		dto.Login	true	"Login data"
//	@Success	200		{object}	global.Response[dto.LoginResponse]
//	@Router		/auth/login [post]
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

// Logout godoc
//
//	@Summary	Remove/Invalidate an access token
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	global.Response[dto.Message]
//	@Security	Bearer
//	@Router		/auth/logout [post]
func (ah *AuthHandler) Logout(c *fiber.Ctx) error {
	var request http.Request

	fasthttpadaptor.ConvertRequest(c.Context(), &request, false)

	err := ah.authUseCase.Logout(dto.Logout{Request: &request})
	if err != nil {
		return err.ToResponse(c)
	}

	return global.MessageResponse("Success", fiber.StatusOK, c)
}

// Verify godoc
//
//	@Summary	Verify an access token to get the user id
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		X-Path		header		string	true	"Path"
//	@Param		X-Method	header		string	true	"Method"
//	@Success	200			{object}	global.Response[dto.VerifyResponse]
//	@Security	Bearer
//	@Router		/auth/verify [get]
func (ah *AuthHandler) Verify(c *fiber.Ctx) error {
	// var header dto.VerifyHeader

	// if err := helper.ValidateHeader(c, &header); err != nil {
	// 	return err.ToResponse(c)
	// }

	var httpRequest http.Request
	fasthttpadaptor.ConvertRequest(c.Context(), &httpRequest, false)

	request := dto.Verify{
		Request: &httpRequest,
		// Path:    header.Path,
		// Method:  header.Method,
	}

	result, err := ah.authUseCase.Verify(request)

	if err != nil {
		return err.ToResponse(c)
	}

	return global.CreateResponse(result, fiber.StatusOK, c)
}

// Refresh godoc
//
//	@Summary	Refresh an access token (get a new one)
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		payload	body		dto.RefreshTokenRequest	true	"User data"
//	@Success	200		{object}	global.Response[dto.RefreshTokenResponse]
//	@Router		/auth/refresh [post]
func (ah *AuthHandler) Refresh(c *fiber.Ctx) error {
	var body dto.RefreshTokenRequest
	if err := helper.ValidateBody(c, &body); err != nil {
		return err.ToResponse(c)
	}

	spec := dto.RefreshTokenSpec(body)
	result, err := ah.authUseCase.RefreshToken(spec)

	if err != nil {
		return err.ToResponse(c)
	}

	return global.CreateResponse(result, http.StatusOK, c)
}
