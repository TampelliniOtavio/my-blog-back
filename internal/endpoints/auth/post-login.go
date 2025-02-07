package auth

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
	internalerror "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/error/internal-error"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) postLogin(ctx *fiber.Ctx) error {
	var loginBody auth.PostLoginBody
	err := ctx.BodyParser(&loginBody)

	if err != nil {
		return err
	}

	err = internalerror.ValidateStruct(loginBody)

	if err != nil {
		return err
	}

	hash, err := h.Service.LoginUser(&loginBody)

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{"token": hash})
}
