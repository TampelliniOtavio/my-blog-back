package authendpoint

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/contract"
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database"
)

type Handler struct {
	Service auth.Service
	Helper  *contract.HandlerEssentials
}

func DefineRoutes(repository *database.UserRepository, helper *contract.HandlerEssentials) {
	handler := &Handler{
		Service: &auth.ServiceImp{
			Repository: repository,
		},
		Helper: helper,
	}

	handler.DefineRoutes()
}
