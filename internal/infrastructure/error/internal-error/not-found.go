package internalerror

import "github.com/gofiber/fiber/v2"

var (
	NotFound = func(name string) *fiber.Error { return fiber.NewError(404, name+" Not Found") }
)
