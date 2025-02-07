package user

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/encrypt"
	internalerrors "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/errors/internal-errors"
	"github.com/rs/xid"
)

type User struct {
	Id       int64  `validate:"required" json:"-"`
	Xid      string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required" json:"-"`
	Email    string `validate:"required,email"`
}

func NewUser(username string, email string, password string) (*User, error) {
	encrypted, err := encrypt.HashPassword(password)

	if err != nil {
		return nil, err
	}

	user := User{
		Id:       -1,
		Xid:      xid.New().String(),
		Username: username,
		Email:    email,
		Password: password,
	}

	err = internalerrors.ValidateStruct(user)

	if err != nil {
		return nil, err
	}

	user.Password = encrypted

	return &user, nil
}
