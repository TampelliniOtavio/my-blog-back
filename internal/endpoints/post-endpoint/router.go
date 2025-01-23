package postendpoint

import "github.com/TampelliniOtavio/my-blog-back/internal/middleware"

func (h *Handler) DefineRoutes() {
	router := h.Helper.Api.Group("/posts").Name("posts.")
	router.Get("/", h.getAllPosts).Name("list")
	router.Get("/:xid", h.getPost).Name("query")

	router.Use(middleware.Protected())
	router.Post("/", h.postAddPost).Name("add")
	router.Delete("/:xid", h.deletePost).Name("delete")
	router.Post("/:xid/like", h.postAddLikeToPost).Name("like")
	router.Post("/:xid/dislike", h.postRemoveLikeFromPost).Name("dislike")
}
