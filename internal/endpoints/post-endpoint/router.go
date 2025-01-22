package postendpoint

import "github.com/TampelliniOtavio/my-blog-back/internal/middleware"

func (h *Handler) DefineRoutes() {
	router := h.Helper.Api.Group("/posts").Name("posts.")
	router.Get("/", h.getAllPosts).Name("list")
	router.Get("/:xid", h.getPost).Name("query")

	router.Use(middleware.Protected())
	router.Post("/", h.postAddPost).Name("add")
	router.Post("/:xid/like", h.PostAddLikeToPost).Name("like")
	router.Post("/:xid/dislike", h.PostRemoveLikeFromPost).Name("dislike")
}
