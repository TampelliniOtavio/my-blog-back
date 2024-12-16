package postendpoint

func (h *Handler) DefineRoutes() {
	router := h.Helper.Api.Group("/posts").Name("posts.")
	router.Get("/", h.getAllPosts).Name("list")
}
