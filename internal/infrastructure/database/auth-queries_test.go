package database_test

import (
	"testing"

	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/util"
	internalerrors "github.com/TampelliniOtavio/my-blog-back/internal/internal-errors"
	"github.com/stretchr/testify/assert"
)

var (
	authRepo *database.AuthRepository
)

func TestInit(t *testing.T) {
	authRepo = repo.Auth
}

func generateRandomUser() *auth.User {
	user, _ := auth.NewUser(
		util.RandomString(10),
		util.RandomEmail(10, 5),
		util.RandomString(10),
	)

	return user
}

func createUser(user *auth.User) (*auth.User, error) {
	return authRepo.CreateUser(user)
}

func Test_createUser_Insert(t *testing.T) {
	assert := assert.New(t)

	insertUser := generateRandomUser()

	createdUser, err := createUser(insertUser)

	assert.NotNil(createdUser)
	assert.Nil(err)
	assert.NotZero(createdUser.Id)
	assert.Equal(createdUser.Xid, insertUser.Xid)
	assert.Equal(createdUser.Username, insertUser.Username)
	assert.Equal(createdUser.Password, insertUser.Password)
	assert.Equal(createdUser.Email, insertUser.Email)
}

func Test_createUser_ErrorUsernameAlreadyExists(t *testing.T) {
	assert := assert.New(t)

	insertUser := generateRandomUser()
	insertUser2 := generateRandomUser()
	insertUser2.Username = insertUser.Username

	createUser(insertUser)
	createdUser, err := createUser(insertUser2)

	assert.Nil(createdUser)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Username already exists")
}

func Test_createUser_ErrorEmailAlreadyExists(t *testing.T) {
	assert := assert.New(t)

	insertUser := generateRandomUser()
	insertUser2 := generateRandomUser()
	insertUser2.Email = insertUser.Email

	createUser(insertUser)
	createdUser, err := createUser(insertUser2)

	assert.Nil(createdUser)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Email already exists")
}

func Test_GetByUsername_FindUser(t *testing.T) {
	assert := assert.New(t)

	insertUser := generateRandomUser()

	createUser(insertUser)

	selectedUser, err := authRepo.GetByUsername(insertUser.Username)

	assert.NotNil(selectedUser)
	assert.Nil(err)
	assert.NotEqual(insertUser.Id, selectedUser.Id) // created user id is -1
	assert.Equal(insertUser.Xid, selectedUser.Xid)
	assert.Equal(insertUser.Email, selectedUser.Email)
	assert.Equal(insertUser.Username, selectedUser.Username)
	assert.Equal(insertUser.Password, selectedUser.Password)
}

func Test_GetByUsername_UserNotFound(t *testing.T) {
	assert := assert.New(t)

	selectedUser, err := authRepo.GetByUsername(util.RandomString(5))

	assert.Nil(selectedUser)
	assert.NotNil(err)
	assert.Equal(err.Error(), internalerrors.NotFound("User").Error())
}
