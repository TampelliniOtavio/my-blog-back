package database_test

import (
	"testing"

	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/user"
	internalerror "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/error/internal-error"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/util"
	"github.com/stretchr/testify/assert"
)

func generateRandomPost(postSize int, user *user.User) *post.Post {
	post, _ := post.NewPost(util.RandomString(postSize), user.Id)
	return post
}

func createPost(post *post.Post) (*post.Post, error) {
	return repo.Post.AddPost(post)
}

func Test_CreatePost_Inserted(t *testing.T) {
	assert := assert.New(t)

	user, _ := createUser(generateRandomUser())

	post := generateRandomPost(20, user)

	inserted, err := createPost(post)

	assert.NotNil(post)
	assert.Nil(err)

	assert.Equal(post.Xid, inserted.Xid)
	assert.Equal(post.Post, inserted.Post)
	assert.Equal(post.CreatedBy, inserted.CreatedBy)
	assert.Equal(post.LikeCount, inserted.LikeCount)
	assert.Equal(post.CreatedAt, inserted.CreatedAt)
	assert.Equal(post.UpdatedAt, inserted.UpdatedAt)
	assert.Equal(post.DeletedAt.String, inserted.DeletedAt.String)
}

func Test_CreatePost_UserNotFound(t *testing.T) {
	assert := assert.New(t)

	user := generateRandomUser() // not created

	post := generateRandomPost(10, user)

	inserted, err := createPost(post)

	assert.Nil(inserted)
	assert.NotNil(err)

	assert.Equal(err.Error(), "User Not Found")
}

func Test_GetAllPosts_List(t *testing.T) {
	assert := assert.New(t)
	quantity := 10

	user, _ := createUser(generateRandomUser())

	for range(quantity) {
		post, err := createPost(generateRandomPost(1, user))
		assert.NotNil(post)
		assert.Nil(err)
	}

	users, err := repo.Post.GetAllPosts(quantity/2, quantity/2, 0)

	assert.NotNil(users)
	assert.Nil(err)

	assert.Equal(quantity/2, len(*users))
}

func Test_AddLikeToPost_AddLike(t *testing.T) {
	assert := assert.New(t)

	user, _ := createUser(generateRandomUser())

	post, _ := createPost(generateRandomPost(1, user))

	err := repo.Post.AddLikeToPost(post, user.Id)
	likedPost, _ := repo.Post.GetPost(post.Xid, user.Id)

	assert.NotNil(likedPost)
	assert.Nil(err)

	assert.Equal(post.Xid, likedPost.Xid)
	assert.Equal(post.LikeCount+1, likedPost.LikeCount)
	assert.True(likedPost.IsLikedByUser)
}

func Test_AddLikeToPost_AddLikeMultipleNotIncrease(t *testing.T) {
	assert := assert.New(t)

	user, _ := createUser(generateRandomUser())

	post, _ := createPost(generateRandomPost(1, user))

	err := repo.Post.AddLikeToPost(post, user.Id)
	assert.Nil(err)
	err = repo.Post.AddLikeToPost(post, user.Id)

	assert.NotNil(err)

	assert.Equal(err.Error(), internalerror.BadRequest("User Already Liked the post").Error())
}

func Test_RemoveLikeFromPost(t *testing.T) {
	assert := assert.New(t)

	user, _ := createUser(generateRandomUser())

	post, _ := createPost(generateRandomPost(1, user))

	repo.Post.AddLikeToPost(post, user.Id)
	likedPost, _ := repo.Post.GetPost(post.Xid, user.Id)

	assert.Equal(post.Xid, likedPost.Xid)
	assert.Equal(post.LikeCount+1, likedPost.LikeCount)
	assert.True(likedPost.IsLikedByUser)

	err := repo.Post.RemoveLikeFromPost(post, user.Id)
	assert.Nil(err)

	dislikedPost, _ := repo.Post.GetPost(post.Xid, user.Id)

	assert.Equal(post.Xid, dislikedPost.Xid)
	assert.Equal(post.LikeCount, dislikedPost.LikeCount)
	assert.False(dislikedPost.IsLikedByUser)
}

func Test_RemoveLikeFromPost_RemoveLikeNotDecrease(t *testing.T) {
	assert := assert.New(t)

	user, _ := createUser(generateRandomUser())

	post, _ := createPost(generateRandomPost(1, user))

	err := repo.Post.RemoveLikeFromPost(post, user.Id)
	assert.NotNil(err)
	assert.Equal(err.Error(), internalerror.NotFound("Liked Post").Error())
}

func Test_GetPost_Valid(t *testing.T) {
	assert := assert.New(t)

	user, _ := createUser(generateRandomUser())

	post, _ := createPost(generateRandomPost(1, user))

	selectedPost, err := repo.Post.GetPost(post.Xid, 0)

	assert.NotNil(selectedPost)
	assert.Nil(err)

	assert.Equal(post.Xid, selectedPost.Xid)
	assert.Equal(post.Post, selectedPost.Post)
	assert.Equal(post.CreatedBy, selectedPost.CreatedBy)
	assert.Equal(post.CreatedAt, selectedPost.CreatedAt)
	assert.Equal(post.UpdatedAt, selectedPost.UpdatedAt)
	assert.Equal(post.LikeCount, selectedPost.LikeCount)
}

func Test_GetPost_NotFound(t *testing.T) {
	assert := assert.New(t)

	user, _ := createUser(generateRandomUser())

	post := generateRandomPost(1, user) // Not Created

	selectedPost, err := repo.Post.GetPost(post.Xid, 0)

	assert.Nil(selectedPost)
	assert.NotNil(err)

	assert.Equal(err.Error(), internalerror.NotFound("Post").Error())
}

func Test_DeletePost_Deleted(t *testing.T) {
	assert := assert.New(t)

	user, _ := createUser(generateRandomUser())

	post, _ := createPost(generateRandomPost(1, user))

	err := repo.Post.DeletePost(post, user.Id)
	assert.Nil(err)

	deletedPost, err := repo.Post.GetPost(post.Xid, 0)

	assert.NotNil(deletedPost)
	assert.NotZero(len(deletedPost.DeletedAt.String))
	assert.True(deletedPost.DeletedAt.Valid)
}

func Test_DeletePost_NotFound(t *testing.T) {
	assert := assert.New(t)

	user, _ := createUser(generateRandomUser())

	post, _ := createPost(generateRandomPost(1, user))

	err := repo.Post.DeletePost(post, user.Id)
	assert.Nil(err)
	err = repo.Post.DeletePost(post, user.Id)
	assert.NotNil(err)

	assert.Equal(err.Error(), internalerror.NotFound("Post").Error())
}
