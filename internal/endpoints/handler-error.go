package endpoints

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func ErrorHandler (ctx *fiber.Ctx, err error) error {
    code := fiber.StatusInternalServerError

    var e *fiber.Error
    if errors.As(err, &e) {
        code = e.Code
    }

    errMessage := "An error Ocourred"

    if code != fiber.StatusInternalServerError {
        errMessage = err.Error()
    }

    log.Error(err.Error())
    err = ctx.Status(code).JSON(fiber.Map{"message": errMessage})
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
    }

    return nil
}
