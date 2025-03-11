package user

import "github.com/gofiber/fiber/v2"

func (h *Handler) getPostsOfUser(ctx *fiber.Ctx) error {
	loggedUser, _ := h.Helper.GetUserFromContext(ctx)

	loggedUserId := int64(0)

	if loggedUser != nil {
		loggedUserId = loggedUser.Id
	}

	posts, err := h.Service.GetPostsByUsername(
		loggedUserId,
		ctx.Params("username"),
		int64(ctx.QueryInt("limit", 20)),
		int64(ctx.QueryInt("offset", 0)),
	)

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(posts)
}
