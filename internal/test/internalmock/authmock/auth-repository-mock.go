package authmock

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct{
    mock.Mock
}

func (c *RepositoryMock) GetByUsername(username string) (*auth.User, error) {
    args := c.Called(username)

    first := args.Get(0)
    err := args.Error(1)

    if first == nil {
        return nil, err
    }

    return first.(*auth.User), err
}

func (c *RepositoryMock) CreateUser(user *auth.User) (*auth.User, error) {
    args := c.Called(user)

    first := args.Get(0)
    err := args.Error(1)

    if first == nil {
        return nil, err
    }

    return first.(*auth.User), err
}
