package internalerror

import "github.com/gofiber/fiber/v2"

var (
	NotAuthorizedError = fiber.NewError(401, "Not Authorized")
)
