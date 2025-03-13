package post

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) getAllPosts(ctx *fiber.Ctx) error {
	userId := int64(0)

	if user, _ := h.Helper.GetUserFromContext(ctx); user != nil {
		userId = user.Id
	}

	posts, err := h.Service.ListAllPosts(post.NewGetAllPostsParams(userId, ctx))

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(posts)
}
