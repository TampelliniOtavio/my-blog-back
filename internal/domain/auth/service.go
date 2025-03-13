package auth

import (
	"os"
	"time"

	"github.com/TampelliniOtavio/my-blog-back/internal/domain/user"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/encrypt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	LoginUser(body *PostLoginBody) (string, error)
	CreateUser(body *PostSigninBody) (*user.User, error)
}

type ServiceImp struct {
	Repository user.Repository
}

func (s *ServiceImp) LoginUser(body *PostLoginBody) (string, error) {
	user, err := s.Repository.GetByUsername(body.Username)

	if err != nil {
		return "", fiber.NewError(400, "Incorrect Username or Password")
	}

	if !encrypt.VerifyPassword(body.Password, user.Password) {
		return "", fiber.NewError(400, "Incorrect Username or Password")
	}

	claims := &jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
		"NotBefore": jwt.NewNumericDate(time.Now()),
		"data": &AuthClaims{
			Xid:      user.Xid,
			Id:       user.Id,
			Name:     user.Email,
			Username: user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", fiber.ErrInternalServerError
	}

	return t, nil
}

func (s *ServiceImp) CreateUser(body *PostSigninBody) (*user.User, error) {
	user, err := user.NewUser(body.Username, body.Email, body.Password)

	if err != nil {
		return nil, err
	}

	return s.Repository.CreateUser(user)
}
