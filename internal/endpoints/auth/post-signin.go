package auth

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
	internalerror "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/error/internal-error"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) postSignin(ctx *fiber.Ctx) error {
	var signinBody auth.PostSigninBody
	err := ctx.BodyParser(&signinBody)

	if err != nil {
		return err
	}

	err = internalerror.ValidateStruct(signinBody)

	if err != nil {
		return err
	}

	user, err := h.Service.CreateUser(&signinBody)

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "User created successfully", "user": user})
}
