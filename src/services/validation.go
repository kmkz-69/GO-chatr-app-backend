package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}


var validate = validator.New()

func ValidateStruct(ctx *fiber.Ctx, v interface{}) (bool, []ErrorResponse) {
	err := validate.Struct(v)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errors []ErrorResponse
		for _, e := range errs {
			errors = append(errors, ErrorResponse{
				FailedField: e.Field(),
				Tag:         e.Tag(),
				Value:       e.Value().(string),
			})
		}
		return false, errors
	}
	return true, nil
}
