package global

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse interface {
	Error() string
	GetMessage() string
	GetCode() int
	GetData() interface{}
	ToResponse(c *fiber.Ctx) error
	GetStack() string
}

type errorResponse struct {
	StatusCode int
	Message    string
	Detail     error
	Data       interface{}
	Stack      string
}

func (e *errorResponse) Error() string {
	if e.Detail != nil {
		return e.Detail.Error()
	}

	return e.Message
}

func (e *errorResponse) GetMessage() string {
	return e.Message
}

func (e *errorResponse) GetCode() int {
	return e.StatusCode
}

func (e *errorResponse) GetData() interface{} {
	return e.Data
}

func (e *errorResponse) GetStack() string {
	return e.Stack
}

func (e *errorResponse) ToResponse(c *fiber.Ctx) error {
	log.Println("[INTERNAL SERVER ERROR]")
	log.Println(e.Detail)
	log.Println(e.Stack)

	response := &Response[interface{}]{
		Code:    e.StatusCode,
		Data:    e.Data,
		Status:  ResponseStatus.FailedResponse,
		Message: e.Message,
	}
	return c.Status(e.StatusCode).JSON(response)
}

func BadRequestError(message string) ErrorResponse {
	return &errorResponse{
		StatusCode: http.StatusBadRequest,
		Detail:     nil,
		Message:    message,
	}
}

func BadRequestErrorWithData(message string, data interface{}) ErrorResponse {
	return &errorResponse{
		StatusCode: http.StatusBadRequest,
		Detail:     nil,
		Message:    message,
		Data:       data,
	}
}

func InternalServerError(err error) ErrorResponse {
	return &errorResponse{
		StatusCode: http.StatusInternalServerError,
		Detail:     err,
		Message:    "Internal Server Error",
		Stack:      string(debug.Stack()),
	}
}

func NotFoundError(messages ...string) ErrorResponse {
	message := "Not Found"
	if len(messages) > 0 {
		message = messages[0]
	}

	return &errorResponse{
		StatusCode: http.StatusNotFound,
		Detail:     nil,
		Message:    message,
	}
}

func ForbiddenError() ErrorResponse {
	return &errorResponse{
		StatusCode: http.StatusForbidden,
		Detail:     nil,
		Message:    "Forbidden",
	}
}

func UnauthorizedError() ErrorResponse {
	return &errorResponse{
		StatusCode: http.StatusUnauthorized,
		Detail:     nil,
		Message:    "Unauthorized",
	}
}
