package post

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) getPost(ctx *fiber.Ctx) error {
	userId := int64(0)

	if user, _ := h.Helper.GetUserFromContext(ctx); user != nil {
		userId = user.Id
	}

	post, err := h.Service.GetPost(ctx.Params("xid"), userId)

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(post)
}
