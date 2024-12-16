package post_test

import (
	"testing"

	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	"github.com/TampelliniOtavio/my-blog-back/internal/test/internalmock/postmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	service = post.ServiceImp{}
	posts   = []post.Post{}
)

func setup() {
	service = post.ServiceImp{}
	posts = []post.Post{}
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
