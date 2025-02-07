package contract

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
	internalerrors "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/errors/internal-errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type HandlerEssentials struct {
	Api fiber.Router
}

func (h *HandlerEssentials) GetUserFromContext(ctx *fiber.Ctx) (*auth.AuthClaims, error) {
	user, ok := ctx.Locals("user").(*jwt.Token)

	if !ok {
		return nil, internalerrors.NotAuthorizedError
	}

	claims, ok := user.Claims.(jwt.MapClaims)

	if !ok {
		return nil, internalerrors.NotAuthorizedError
	}

	data, ok := claims["data"].(map[string]interface{})

	if !ok {
		return nil, internalerrors.NotAuthorizedError
	}

	authClaims := auth.AuthClaims{
		Xid:      data["xid"].(string),
		Id:       int64(data["id"].(float64)),
		Name:     data["name"].(string),
		Username: data["username"].(string),
	}

	return &authClaims, nil
}
