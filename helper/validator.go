package helper

import (
	"debozero-core/dto"
	"debozero-core/global"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct(request any) []dto.ValidationError {
	var validationErrors []dto.ValidationError
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element dto.ValidationError
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			validationErrors = append(validationErrors, element)
		}
	}

	if validationErrors != nil {
		return validationErrors
	}

	return nil
}

func ValidateBody(c *fiber.Ctx, request any) global.ErrorResponse {
	if err := c.BodyParser(request); err != nil {
		return global.BadRequestErrorWithData("Validation Error", err)
	}

	if err := ValidateStruct(request); err != nil {
		return global.BadRequestErrorWithData("Validation Error", err)
	}

	return nil
}

func ValidateParam(c *fiber.Ctx, param any) global.ErrorResponse {
	if err := c.ParamsParser(param); err != nil {
		return global.BadRequestErrorWithData("Validation Error", err)
	}

	if err := ValidateStruct(param); err != nil {
		return global.BadRequestErrorWithData("Validation Error", err)
	}

	return nil
}

func ValidateHeader(c *fiber.Ctx, header any) global.ErrorResponse {
	if err := c.ReqHeaderParser(header); err != nil {
		return global.BadRequestErrorWithData("Validation Error", err)
	}

	if err := ValidateStruct(header); err != nil {
		return global.BadRequestErrorWithData("Validation Error", err)
	}

	return nil
}

func GetUserID(c *fiber.Ctx) (string, global.ErrorResponse) {
	userId := c.Get("X-UserId")
	if userId == "" {
		return "", global.ForbiddenError()
	}
	return userId, nil
}

func ValidateQuery(c *fiber.Ctx, query any) global.ErrorResponse {
	if err := c.QueryParser(query); err != nil {
		return global.BadRequestErrorWithData("Validation Error", err)
	}

	if err := ValidateStruct(query); err != nil {
		return global.BadRequestErrorWithData("Validation Error", err)
	}

	return nil
}
