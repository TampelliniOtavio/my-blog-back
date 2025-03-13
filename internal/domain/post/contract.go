package post

import "github.com/gofiber/fiber/v2"

type AddPostBody struct {
	Post string `validate:"required"`
}

type GetAllPostsQueries struct {
	Username string
	Limit    int
	Offset   int
}

type ListAllPostsParams struct {
	AuthUserId int64
	Queries    *GetAllPostsQueries
}

func NewGetAllPostsParams(userId int64, ctx *fiber.Ctx) *ListAllPostsParams {
	return &ListAllPostsParams{
		AuthUserId: userId,
		Queries: &GetAllPostsQueries{
			Limit:    ctx.QueryInt("limit", 20),
			Offset:   ctx.QueryInt("offset", 0),
			Username: ctx.Query("username"),
		},
	}
}
