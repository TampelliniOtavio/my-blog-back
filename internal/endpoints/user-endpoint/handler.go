package userendpoint

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/contract"
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/user"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database"
)

type Handler struct {
	Service user.Service
	Helper  *contract.HandlerEssentials
}

func DefineRoutes(repository *database.UserRepository, helper *contract.HandlerEssentials) {
	handler := Handler{
		Service: &user.ServiceImp{
			Repository: repository,
		},
		Helper: helper,
	}

	handler.DefineRoutes()
}
