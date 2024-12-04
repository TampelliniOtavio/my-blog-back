package auth_test

import (
	"errors"
	"testing"

	authcontract "github.com/TampelliniOtavio/my-blog-back/internal/contract/auth-contract"
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
	"github.com/TampelliniOtavio/my-blog-back/internal/test/internalmock/authmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
    service = auth.ServiceImp{}
    loginBody = authcontract.PostLoginBody{
        Username: "username",
        Password: "password",
    }
)

func Test_LoginUser_should_login (t *testing.T) {
    assert := assert.New(t)

    user, _ := auth.NewUser(loginBody.Username, "email@email.com", loginBody.Password)

    repository := new(authmock.RepositoryMock)

    repository.On("GetByUsername", mock.Anything).Return(user, nil)
    service.Repository = repository

    _, err := service.LoginUser(&loginBody)

    assert.Nil(err)
}

func Test_LoginUser_should_not_login_incorrect_password (t *testing.T) {
    assert := assert.New(t)

    user, _ := auth.NewUser(loginBody.Username, "email@email.com", "OtherPassword")

    repository := new(authmock.RepositoryMock)

    repository.On("GetByUsername", mock.Anything).Return(user, nil)
    service.Repository = repository

    _, err := service.LoginUser(&loginBody)

    assert.NotNil(err)
}

func Test_LoginUser_should_not_login_user_not_found (t *testing.T) {
    assert := assert.New(t)

    repository := new(authmock.RepositoryMock)

    repository.On("GetByUsername", mock.Anything).Return(nil, errors.New("User Not Found"))
    service.Repository = repository

    _, err := service.LoginUser(&loginBody)

    assert.NotNil(err)
}
