package validation

import (
	"firstapp/util/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Util interface {
	Validate(c *fiber.Ctx, body interface{}) error
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type util struct {
	res response.Util
}

func Init(res response.Util) Util {
	return util{
		res: res,
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

func (u util) Validate(c *fiber.Ctx, body interface{}) error {
	if processValidation := u.process(body); processValidation != nil {
		resConfig := response.Config{
			Code:    422,
			Message: "Validation Failure",
		}

		return u.res.Send(c, processValidation, resConfig)
	}

	return nil
}
