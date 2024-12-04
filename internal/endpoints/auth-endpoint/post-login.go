package authendpoint

import (
	authcontract "github.com/TampelliniOtavio/my-blog-back/internal/contract/auth-contract"
	internalerrors "github.com/TampelliniOtavio/my-blog-back/internal/internal-errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) postLogin (ctx *fiber.Ctx) error {
    var loginBody authcontract.PostLoginBody
    err := ctx.BodyParser(&loginBody)

    if err != nil {
        return err
    }

    err = internalerrors.ValidateStruct(loginBody)

    if err != nil {
        return err
    }

    hash, err := h.Service.LoginUser(&loginBody)

    if err != nil {
        return err
    }

    return ctx.Status(200).JSON(fiber.Map{"token": hash})
}
