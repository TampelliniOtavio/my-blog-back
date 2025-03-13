package post

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (c *RepositoryMock) GetAllPosts(params *post.ListAllPostsParams) (*[]post.Post, error) {
	args := c.Called(params)

	first := args.Get(0)
	err := args.Error(1)

	if first == nil {
		return nil, err
	}

	return first.(*[]post.Post), err
}

func (c *RepositoryMock) AddPost(newPost *post.Post) (*post.Post, error) {
	args := c.Called(newPost)

	first := args.Get(0)
	err := args.Error(1)

	if first == nil {
		return nil, err
	}

	return first.(*post.Post), err
}

func (c *RepositoryMock) GetPost(xid string, authUserId int64) (*post.Post, error) {
	args := c.Called(xid, authUserId)

	first := args.Get(0)
	err := args.Error(1)

	if first == nil {
		return nil, err
	}

	return first.(*post.Post), err
}

func (c *RepositoryMock) AddLikeToPost(post *post.Post, userId int64) error {
	args := c.Called(post, userId)

	return args.Error(0)
}

func (c *RepositoryMock) RemoveLikeFromPost(post *post.Post, userId int64) error {
	args := c.Called(post, userId)

	return args.Error(0)
}

func (c *RepositoryMock) DeletePost(post *post.Post, userId int64) error {
	args := c.Called(post, userId)

	return args.Error(0)
}
