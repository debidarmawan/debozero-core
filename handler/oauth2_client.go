package handler

import (
	"debozero-core/dto"
	"debozero-core/global"
	"debozero-core/helper"
	"debozero-core/usecase"

	"github.com/gofiber/fiber/v2"
)

type Oauth2ClientHandler struct {
	oauth2UseCase usecase.Oauth2UseCase
}

func NewOauth2Handler(oauth2UseCase usecase.Oauth2UseCase) *Oauth2ClientHandler {
	return &Oauth2ClientHandler{
		oauth2UseCase: oauth2UseCase,
	}
}

func (oh *Oauth2ClientHandler) Routes(group fiber.Router) {
	group.Post("/oauth2/client", oh.CreateClient)
}

// Oauth2Client godoc
//
//	@Summary		Create Oauth2 Client
//	@Description	Create Oauth2 Client
//	@Tags			Oauth2 Client
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		dto.Oauth2Client	true	"Oauth2 Client data"
//	@Success		200		{object}	global.Response[dto.Oauth2ClientResponse]
//	@Router			/oauth2/client [post]
func (oh *Oauth2ClientHandler) CreateClient(c *fiber.Ctx) error {
	var request dto.Oauth2Client

	if err := helper.ValidateBody(c, &request); err != nil {
		return err.ToResponse(c)
	}

	result, err := oh.oauth2UseCase.AddClient(request)
	if err != nil {
		return err.ToResponse(c)
	}

	return global.CreateResponse(result, fiber.StatusOK, c)
}
