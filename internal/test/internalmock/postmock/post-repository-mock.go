package postmock

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (c *RepositoryMock) GetAllPosts(limit int, offset int) (*[]post.Post, error) {
	args := c.Called(limit, offset)

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

func (c *RepositoryMock) GetPost(xid string) (*post.Post, error) {
	args := c.Called(xid)

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
