package postendpoint

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/contract"
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database"
)

type Handler struct {
	Service post.Service
	Helper  *contract.HandlerEssentials
}

func DefineRoutes(repository *database.PostRepository, helper *contract.HandlerEssentials) {
	handler := Handler{
		Service: &post.ServiceImp{
			Repository: repository,
		},
		Helper: helper,
	}

	handler.DefineRoutes()
}
