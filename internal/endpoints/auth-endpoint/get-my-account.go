package authendpoint

import "github.com/gofiber/fiber/v2"

func (h *Handler) getMyAccount (ctx *fiber.Ctx) error {
    user, err := h.Helper.GetUserFromContext(ctx)

    if err != nil {
        return err
    }

    return ctx.Status(200).JSON(user)
}
