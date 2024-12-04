package authendpoint

import "github.com/TampelliniOtavio/my-blog-back/internal/middleware"

func (h *Handler) DefineRoutes() {
    router := h.Helper.Api.Group("/auth").Name("auth.")
    router.Post("/login", h.postLogin).Name("login")
}
