package post

import "github.com/gofiber/fiber/v2"

func (h *Handler) getPost(ctx *fiber.Ctx) error {
	post, err := h.Service.GetPost(ctx.Params("xid"))

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(post)
}
