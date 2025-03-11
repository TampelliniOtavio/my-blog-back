package user

func (h *Handler) DefineRoutes() {
	router := h.Helper.Api.Group("/users").Name("users.")

	router.Get("/:username", h.getOneUser).Name("one-user")
	router.Get("/:username/posts", h.getPostsOfUser).Name("posts")
}
