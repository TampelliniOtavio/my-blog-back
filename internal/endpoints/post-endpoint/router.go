package postendpoint

import "github.com/TampelliniOtavio/my-blog-back/internal/middleware"

func (h *Handler) DefineRoutes() {
	router := h.Helper.Api.Group("/posts").Name("posts.")
	router.Get("/", h.getAllPosts).Name("list")

	router.Use(middleware.Protected())
	router.Post("/", h.postAddPost).Name("add")
}
