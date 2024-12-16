package postendpoint

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/contract"
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
)

type Handler struct {
	Service post.Service
	Helper  *contract.HandlerEssentials
}
