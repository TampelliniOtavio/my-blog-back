package postendpoint

import "github.com/gofiber/fiber/v2"

func (h *Handler) deletePost(ctx *fiber.Ctx) error {
	auth, err := h.Helper.GetUserFromContext(ctx)

	if err != nil {
		return err
	}

	post, err := h.Service.DeletePost(ctx.Params("xid"), auth.Id)

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(post)
}
