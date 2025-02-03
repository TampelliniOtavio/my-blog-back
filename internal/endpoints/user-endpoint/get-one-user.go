package userendpoint

import "github.com/gofiber/fiber/v2"

func (h *Handler) getOneUser(ctx *fiber.Ctx) error {
	user, err := h.Service.GetByUsername(ctx.Params("username"))

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(user)
}
