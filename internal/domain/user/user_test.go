package user_test

import (
	"testing"

	"github.com/TampelliniOtavio/my-blog-back/internal/domain/user"
	"github.com/stretchr/testify/assert"
)

var (
	username = "username"
	password = "password"
	email    = "email@email.com"
)

func Test_NewUser_should_create(t *testing.T) {
	assert := assert.New(t)

	user, err := user.NewUser(username, email, password)

	assert.NotNil(user)
	assert.Nil(err)
}

func Test_NewUser_username_should_be_required(t *testing.T) {
	assert := assert.New(t)

	_, err := user.NewUser("", email, password)

	assert.NotNil(err)
	assert.Equal(err.Error(), "username is required")
}

func Test_NewUser_email_should_be_required(t *testing.T) {
	assert := assert.New(t)

	_, err := user.NewUser(username, "", password)

	assert.NotNil(err)
	assert.Equal(err.Error(), "email is required")
}

func Test_NewUser_password_should_be_required(t *testing.T) {
	assert := assert.New(t)

	_, err := user.NewUser(username, email, "")

	assert.NotNil(err)
	assert.Equal(err.Error(), "password is required")
}

func Test_NewUser_should_validate_email(t *testing.T) {
	assert := assert.New(t)

	_, err := user.NewUser(username, "wrongEmail", password)

	assert.NotNil(err)
	assert.Equal(err.Error(), "email is invalid")
}
