package post

import "github.com/gofiber/fiber/v2"

func (h *Handler) postAddLikeToPost(ctx *fiber.Ctx) error {
	user, err := h.Helper.GetUserFromContext(ctx)

	if err != nil {
		return err
	}

	post, err := h.Service.AddLikeToPost(ctx.Params("xid"), user.Id)

	if err != nil {
		return err
	}

	return ctx.Status(201).JSON(post)
}
