package auth

import "github.com/TampelliniOtavio/my-blog-back/internal/middleware"

func (h *Handler) DefineRoutes() {
	router := h.Helper.Api.Group("/auth").Name("auth.")
	router.Post("/login", h.postLogin).Name("login")
	router.Post("/signin", h.postSignin).Name("signin")

	router.Use(middleware.Protected(&middleware.ProtectedParams{}))
	router.Get("/", h.getMyAccount).Name("my-account")
}
