package internalerror

import "github.com/gofiber/fiber/v2"

var (
	BadRequest = func(str string) *fiber.Error { return fiber.NewError(400, str) }
)
