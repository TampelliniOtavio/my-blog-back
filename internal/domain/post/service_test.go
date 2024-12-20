package post_test

import (
	"testing"

	postcontract "github.com/TampelliniOtavio/my-blog-back/internal/contract/post-contract"
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/formatter"
	"github.com/TampelliniOtavio/my-blog-back/internal/test/internalmock/postmock"
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
		CreatedAt: formatter.CurrentTimestamp(),
		UpdatedAt: formatter.CurrentTimestamp(),
	}
	addPostBody = postcontract.PostAddPostBody{
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
	addPostBody = postcontract.PostAddPostBody{
		Post: newPost.Post,
	}
}

func Test_GetAllPosts_should_return(t *testing.T) {
	setup()
	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetAllPosts", mock.Anything, mock.Anything).Return(&posts, nil)
	service.Repository = repository

	post, err := service.ListAllPosts(0, 1)

	assert.Nil(err)
	assert.NotNil(post)
}

func Test_GetAllPosts_should_validate_limit(t *testing.T) {
	setup()
	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetAllPosts", mock.Anything, mock.Anything).Return(&posts, nil)
	service.Repository = repository

	_, err := service.ListAllPosts(-1, 1)

	assert.NotNil(err)
	assert.Equal(err.Error(), "Limit is not valid")
}

func Test_GetAllPosts_should_validate_offset(t *testing.T) {
	setup()
	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("GetAllPosts", mock.Anything, mock.Anything).Return(&posts, nil)
	service.Repository = repository

	_, err := service.ListAllPosts(1, -1)

	assert.NotNil(err)
	assert.Equal(err.Error(), "Offset is not valid")
}

func Test_AddPost_should_insert(t *testing.T) {
	setup()

	assert := assert.New(t)

	repository := new(postmock.RepositoryMock)

	repository.On("AddPost", mock.Anything, mock.Anything).Return(&newPost, nil)
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

	repository.On("AddPost", mock.Anything, mock.Anything).Return(&newPost, nil)
	service.Repository = repository

	addedPost, err := service.AddPost(&postcontract.PostAddPostBody{
		Post: "",
	}, newPost.CreatedBy)

	assert.Nil(addedPost)
	assert.NotNil(err)
	assert.Equal(err.Error(), "post is required")
}
