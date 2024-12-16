package postendpoint

import "github.com/gofiber/fiber/v2"

func (h *Handler) getAllPosts(ctx *fiber.Ctx) error {
	posts, err := h.Service.ListAllPosts(ctx.QueryInt("limit", 0), ctx.QueryInt("offset", 20))

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(posts)
}
