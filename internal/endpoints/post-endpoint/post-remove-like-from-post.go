
package postendpoint

import "github.com/gofiber/fiber/v2"

func (h *Handler) postRemoveLikeFromPost(ctx *fiber.Ctx) error {
    user, err := h.Helper.GetUserFromContext(ctx)

    if err != nil {
        return err
    }

	post, err := h.Service.RemoveLikeFromPost(ctx.Params("xid"), user.Id)

    if err != nil {
        return err
    }

	return ctx.Status(201).JSON(post)
}
