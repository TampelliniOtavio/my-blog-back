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
    signinBody = authcontract.PostSigninBody{
        Username: "username",
        Email: "email@email.com",
        Password: "password",
    }
)

func setup() {
    service = auth.ServiceImp{}
    loginBody = authcontract.PostLoginBody{
        Username: "username",
        Password: "password",
    }
    signinBody = authcontract.PostSigninBody{
        Username: "username",
        Email: "email@email.com",
        Password: "password",
    }
}

func Test_LoginUser_should_login (t *testing.T) {
    setup()
    assert := assert.New(t)

    user, _ := auth.NewUser(loginBody.Username, "email@email.com", loginBody.Password)

    repository := new(authmock.RepositoryMock)

    repository.On("GetByUsername", mock.Anything).Return(user, nil)
    service.Repository = repository

    _, err := service.LoginUser(&loginBody)

    assert.Nil(err)
}

func Test_LoginUser_should_not_login_incorrect_password (t *testing.T) {
    setup()
    assert := assert.New(t)

    user, _ := auth.NewUser(loginBody.Username, "email@email.com", "OtherPassword")

    repository := new(authmock.RepositoryMock)

    repository.On("GetByUsername", mock.Anything).Return(user, nil)
    service.Repository = repository

    _, err := service.LoginUser(&loginBody)

    assert.NotNil(err)
}

func Test_LoginUser_should_not_login_user_not_found (t *testing.T) {
    setup()
    assert := assert.New(t)

    repository := new(authmock.RepositoryMock)

    repository.On("GetByUsername", mock.Anything).Return(nil, errors.New("User Not Found"))
    service.Repository = repository

    _, err := service.LoginUser(&loginBody)

    assert.NotNil(err)
}

func Test_CreateUser_should_create(t *testing.T) {
    setup()
    assert := assert.New(t)

    repository := new(authmock.RepositoryMock)

    user, _ := auth.NewUser(signinBody.Username, signinBody.Email, signinBody.Password)

    repository.On("CreateUser", mock.Anything).Return(user, nil)
    service.Repository = repository

    user, err := service.CreateUser(&signinBody)

    assert.Nil(err)
    assert.NotNil(user)
    assert.NotNil(user.Id)
    assert.NotNil(user.Xid)
    assert.Equal(user.Username, signinBody.Username)
    assert.Equal(user.Email, signinBody.Email)
    assert.NotEqual(user.Password, signinBody.Password, "Password Should be Encrypted")
}
