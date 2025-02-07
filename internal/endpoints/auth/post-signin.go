package auth

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
	internalerrors "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/errors/internal-errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) postSignin(ctx *fiber.Ctx) error {
	var signinBody auth.PostSigninBody
	err := ctx.BodyParser(&signinBody)

	if err != nil {
		return err
	}

	err = internalerrors.ValidateStruct(signinBody)

	if err != nil {
		return err
	}

	user, err := h.Service.CreateUser(&signinBody)

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "User created successfully", "user": user})
}
