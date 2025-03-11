package database_test

import (
	"testing"

	"github.com/TampelliniOtavio/my-blog-back/internal/domain/user"
	internalerror "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/error/internal-error"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/util"
	"github.com/stretchr/testify/assert"
)

func generateRandomUser() *user.User {
	user, _ := user.NewUser(
		util.RandomString(10),
		util.RandomEmail(10, 5),
		util.RandomString(10),
	)

	return user
}

func createUser(createUser *user.User) (*user.User, error) {
	return repo.User.CreateUser(createUser)
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

	selectedUser, err := repo.User.GetByUsername(insertUser.Username)

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

	selectedUser, err := repo.User.GetByUsername(util.RandomString(5))

	assert.Nil(selectedUser)
	assert.NotNil(err)
	assert.Equal(err.Error(), internalerror.NotFound("User").Error())
}

func Test_GetPostsByUsername(t *testing.T) {
	assert := assert.New(t)

	newUser, err := createUser(generateRandomUser())
	assert.NotNil(newUser)
	assert.NoError(err)

	postLength := 10

	for range(postLength) {
		post, err := createPost(generateRandomPost(10, newUser))

		repo.Post.AddLikeToPost(post, newUser.Id)

		assert.NotNil(post)
		assert.NoError(err)
	}

	posts, err := repo.User.GetPostsByUsername(&user.GetPostsByUsernameParams{
		UserId: newUser.Id,
		Username: newUser.Username,
		Limit: int64(postLength),
		Offset: 0,
	})

	assert.NotNil(posts)
	assert.NoError(err)

	size := len(*posts)

	assert.Equal(size, postLength)

	for _, post := range(*posts) {
		assert.True(post.IsLikedByUser)
		assert.Equal(post.CreatedBy, newUser.Id)
	}
}
