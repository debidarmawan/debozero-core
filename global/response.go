package global

import (
	"github.com/debidarmawan/debozero-core/dto"
	"github.com/gofiber/fiber/v2"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

var ResponseStatus = struct {
	FailedResponse  string
	SuccessResponse string
	RetryResponse   string
}{
	FailedResponse:  "FAILED",
	SuccessResponse: "OK",
	RetryResponse:   "RETRY",
}

type Result[T any] struct {
	Data       *T
	Error      error
	StatusCode int
}

func CreateMessageResponse(message string, statusCode int, c *fiber.Ctx) error {
	response := Response[any]{
		Code:    statusCode,
		Data:    nil,
		Status:  ResponseStatus.SuccessResponse,
		Message: message,
	}

	return c.Status(statusCode).JSON(response)
}

func CreateResponse[T any](data *T, statusCode int, c *fiber.Ctx) error {
	response := Response[T]{
		Code:    statusCode,
		Data:    *data,
		Status:  ResponseStatus.SuccessResponse,
		Message: "",
	}

	return c.Status(statusCode).JSON(response)
}

func MessageResponse(message string, statusCode int, c *fiber.Ctx) error {
	return CreateResponse(&dto.Message{Message: message}, statusCode, c)
}

func (r *Result[T]) ToResponseError() *Response[*T] {
	return &Response[*T]{
		Code:    r.StatusCode,
		Data:    r.Data,
		Status:  ResponseStatus.FailedResponse,
		Message: r.Error.Error(),
	}
}
