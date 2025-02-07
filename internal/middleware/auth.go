package middleware

import (
	"errors"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(
		jwtware.Config{
			ErrorHandler: jwtError,
			SigningKey: jwtware.SigningKey{
				Key: []byte(os.Getenv("JWT_SECRET")),
			},
		},
	)
}

func jwtError(c *fiber.Ctx, err error) error {
	code := 401

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	errMessage := err.Error()
	if errMessage == "missing or malformed JWT" {
		errMessage = "Not Authorized"
	}

	return fiber.NewError(code, errMessage)
}
