package user

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/user"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (c *RepositoryMock) GetByUsername(username string) (*user.User, error) {
	args := c.Called(username)

	first := args.Get(0)
	err := args.Error(1)

	if first == nil {
		return nil, err
	}

	return first.(*user.User), err
}

func (c *RepositoryMock) CreateUser(createUser *user.User) (*user.User, error) {
	args := c.Called(createUser)

	first := args.Get(0)
	err := args.Error(1)

	if first == nil {
		return nil, err
	}

	return first.(*user.User), err
}
