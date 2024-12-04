package authendpoint

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/contract"
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
)

type Handler struct{
    Service auth.Service
    Helper *contract.HandlerEssentials
}
