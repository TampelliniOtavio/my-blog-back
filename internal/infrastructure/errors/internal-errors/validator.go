package internalerrors

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateStruct(object interface{}) error {
	validate := validator.New()

	err := validate.Struct(object)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	field := strings.ToLower(validationError.StructField())
	switch validationError.Tag() {
	case "required":
		return fiber.NewError(400, field+" is required")
	case "max":
		return fiber.NewError(400, field+" is required with max "+validationError.Param())
	case "min":
		return fiber.NewError(400, field+" is required with min "+validationError.Param())
	case "email":
		return fiber.NewError(400, field+" is invalid")
	}

	return nil
}
