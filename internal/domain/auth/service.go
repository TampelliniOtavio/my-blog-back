package auth

import (
	"os"
	"time"

	authcontract "github.com/TampelliniOtavio/my-blog-back/internal/contract/auth-contract"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/encrypt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Service interface{
    LoginUser(body *authcontract.PostLoginBody) (string, error)
}

type ServiceImp struct{
    Repository Repository
}

func (s *ServiceImp) LoginUser(body *authcontract.PostLoginBody) (string, error) {
    user, err := s.Repository.GetByUsername(body.Username)

    if err != nil {
        return "", fiber.NewError(400, "Incorrect Username or Password")
    }

    if !encrypt.VerifyPassword(user.Password, body.Password) {
        return "", fiber.NewError(400, "Incorrect Username or Password")
    }

    claims := &jwt.MapClaims{
        "ExpiresAt": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        "IssuedAt":  jwt.NewNumericDate(time.Now()),
        "NotBefore": jwt.NewNumericDate(time.Now()),
        "data": &AuthClaims{
            Name: user.Email,
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

    if err != nil {
        return "", fiber.ErrInternalServerError
    }

    return t, nil
}
