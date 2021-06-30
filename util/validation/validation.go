package validation

import (
	"firstapp/util/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Util interface {
	Validate(c *fiber.Ctx, body interface{}) ([]*ErrorResponse, error)
	ValidateAndBodyParser(c *fiber.Ctx, body interface{}) ([]*ErrorResponse, error)
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type util struct {
	response response.Util
}

func Init(response response.Util) Util {
	return util{
		response: response,
	}
}

func (u util) process(s interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (u util) Validate(c *fiber.Ctx, body interface{}) ([]*ErrorResponse, error) {
	if processValidation := u.process(body); processValidation != nil {
		resConfig := response.Config{
			Code:    422,
			Message: "validation failure",
		}

		return processValidation, u.response.Send(c, processValidation, resConfig)
	}

	return nil, nil
}

func (u util) ValidateAndBodyParser(c *fiber.Ctx, body interface{}) ([]*ErrorResponse, error) {
	var listErr []*ErrorResponse

	if err := c.BodyParser(body); err != nil {
		listErr = append(listErr, &ErrorResponse{})
		return listErr, err
	}

	return u.Validate(c, body)
}
