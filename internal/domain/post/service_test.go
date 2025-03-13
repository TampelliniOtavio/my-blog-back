package post_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database/types"
	internalerror "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/error/internal-error"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/formatter"
	postmock "github.com/TampelliniOtavio/my-blog-back/internal/test/internalmock/post"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	service = post.ServiceImp{}
	posts   = []post.Post{}
	newPost = post.Post{
		Xid:       "randomxid",
		Post:      "New Post",
		CreatedBy: 1,
		LikeCount: 0,
		CreatedAt: formatter.CurrentTimestamp(),
		UpdatedAt: formatter.CurrentTimestamp(),
	}
	deletedPost = post.Post{
		Xid:       "randomxid",
		Post:      "New Post",
		CreatedBy: 1,
		LikeCount: 0,
		CreatedAt: formatter.CurrentTimestamp(),
		UpdatedAt: formatter.CurrentTimestamp(),
		DeletedAt: types.NewNullString(formatter.CurrentTimestamp()),
	}
	addPostBody = post.AddPostBody{
		Post: newPost.Post,
	}
)

func setup() {
	service = post.ServiceImp{}
	posts = []post.Post{}
	newPost = post.Post{
		Xid:       "randomxid",
		Post:      "New Post",
		CreatedBy: 1,
		CreatedAt: formatter.CurrentTimestamp(),
		UpdatedAt: formatter.CurrentTimestamp(),
	}
	addPostBody = post.AddPostBody{
		Post: newPost.Post,
	}
}

func newListPostsParams(limit int, offset int, userId int64) *post.ListAllPostsParams {
	return &post.ListAllPostsParams{
		AuthUserId: userId,
		Queries: &post.GetAllPostsQueries{
			Username: "",
			Limit: limit,
			Offset: offset,
		},
	}
}

func Test_GetAllPosts_should_return(t *testing.T) {
	setup()
	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetAllPosts", mock.Anything).Return(&posts, nil)
	service.Repository = repository

	post, err := service.ListAllPosts(newListPostsParams(10, 0, 0))

	assert.Nil(err)
	assert.NotNil(post)
}

func Test_GetAllPosts_should_validate_limit(t *testing.T) {
	setup()
	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetAllPosts", mock.Anything).Return(&posts, nil)
	service.Repository = repository

	_, err := service.ListAllPosts(newListPostsParams(-1, 1, 0))

	assert.NotNil(err)
	assert.Equal(err.Error(), "Limit is not valid")
}

func Test_GetAllPosts_should_validate_offset(t *testing.T) {
	setup()
	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetAllPosts", mock.Anything).Return(&posts, nil)
	service.Repository = repository

	_, err := service.ListAllPosts(newListPostsParams(1, -1, 0))

	assert.NotNil(err)
	assert.Equal(err.Error(), "Offset is not valid")
}

func Test_AddPost_should_insert(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("AddPost", mock.Anything, mock.Anything, mock.Anything).Return(&newPost, nil)
	service.Repository = repository

	addedPost, err := service.AddPost(&addPostBody, newPost.CreatedBy)

	assert.Nil(err)
	assert.Equal(addedPost.Post, addPostBody.Post)
	assert.Equal(newPost.CreatedBy, addedPost.CreatedBy)
}

func Test_AddPost_validate_post_required(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("AddPost", mock.Anything, mock.Anything, mock.Anything).Return(&newPost, nil)
	service.Repository = repository

	addedPost, err := service.AddPost(&post.AddPostBody{
		Post: "",
	}, newPost.CreatedBy)

	assert.Nil(addedPost)
	assert.NotNil(err)
	assert.Equal(err.Error(), "post is required")
}

func Test_GetPost_ShouldReturn(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetPost", mock.Anything, mock.Anything).Return(&newPost, nil)
	service.Repository = repository

	currPost, err := service.GetPost(newPost.Xid, 0)

	assert.Nil(err)
	assert.NotNil(currPost)
	assert.Equal(currPost.Xid, newPost.Xid)
}

func Test_GetPost_ShouldNotFind(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetPost", mock.Anything, mock.Anything).Return(nil, sql.ErrNoRows)
	service.Repository = repository

	currPost, err := service.GetPost("xid-not-valid", 0)

	assert.Nil(currPost)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Post Not Found")
}

func Test_GetPost_ShouldReturnGenericError(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	errorMessage := "Any Error"

	repository.On("GetPost", mock.Anything, mock.Anything).Return(nil, errors.New(errorMessage))
	service.Repository = repository

	currPost, err := service.GetPost(newPost.Xid, 0)

	assert.Nil(currPost)
	assert.NotNil(err)
	assert.Equal(err.Error(), errorMessage)
}

func Test_AddLikeToPost_Added(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetPost", mock.Anything, mock.Anything).Return(&newPost, nil)
	repository.On("AddLikeToPost", mock.Anything, mock.Anything).Return(nil)
	service.Repository = repository

	updatedPost, err := service.AddLikeToPost("randomxid", 1)

	assert.NotNil(updatedPost)
	assert.Nil(err)

	assert.Equal(newPost.LikeCount, updatedPost.LikeCount)
}

func Test_AddLikeToPost_PostNotFound(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetPost", mock.Anything, mock.Anything).Return(nil, errors.New("Not Found"))
	service.Repository = repository

	updatedPost, err := service.AddLikeToPost("randomxid", 1)

	assert.Nil(updatedPost)
	assert.NotNil(err)

	assert.Equal(err.Error(), "Post Not Found")
}

func Test_RemoveLikeFromPost_Removed(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetPost", mock.Anything, mock.Anything).Return(&newPost, nil)
	repository.On("RemoveLikeFromPost", mock.Anything, mock.Anything).Return(nil)
	service.Repository = repository

	updatedPost, err := service.RemoveLikeFromPost("randomxid", 1)

	assert.NotNil(updatedPost)
	assert.Nil(err)

	assert.Equal(newPost.LikeCount, updatedPost.LikeCount)
}

func Test_RemoveLikeToPost_PostNotFound(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetPost", mock.Anything, mock.Anything).Return(nil, errors.New("Not Found"))
	service.Repository = repository

	updatedPost, err := service.RemoveLikeFromPost("randomxid", 1)

	assert.Nil(updatedPost)
	assert.NotNil(err)

	assert.Equal(err.Error(), "Post Not Found")
}

func Test_DeletePost_Deleted(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetPost", mock.Anything, mock.Anything).Return(&deletedPost, nil)
	repository.On("DeletePost", mock.Anything, mock.Anything).Return(nil)
	service.Repository = repository

	deleted, err := service.DeletePost("randomxid", 1)

	assert.NotNil(deleted)
	assert.Nil(err)

	assert.Equal(deleted.DeletedAt.Valid, true)
	assert.True(len(deleted.DeletedAt.String) > 0)
}

func Test_DeletePost_NotFound(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetPost", mock.Anything, mock.Anything).Return(nil, errors.New("Error"))
	service.Repository = repository

	deleted, err := service.DeletePost("randomxid", 1)

	assert.Nil(deleted)
	assert.NotNil(err)

	assert.Equal(err.Error(), internalerror.NotFound("Post").Error())
}

func Test_DeletePost_InternalError(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	errorMessage := "An Error ocourred"

	repository.On("GetPost", mock.Anything, mock.Anything).Return(&deletedPost, nil)
	repository.On("DeletePost", mock.Anything, mock.Anything).Return(errors.New(errorMessage))
	service.Repository = repository

	deleted, err := service.DeletePost("randomxid", 1)

	assert.Nil(deleted)
	assert.NotNil(err)

	assert.Equal(err.Error(), errorMessage)
}
