package postendpoint

import (
	postcontract "github.com/TampelliniOtavio/my-blog-back/internal/contract/post-contract"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) postAddPost(ctx *fiber.Ctx) error {
	auth, err := h.Helper.GetUserFromContext(ctx)

	if err != nil {
		return err
	}

	var body postcontract.PostAddPostBody
	err = ctx.BodyParser(&body)

	if err != nil {
		return err
	}

	post, err := h.Service.AddPost(&body, auth.Id)

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(post)
}
