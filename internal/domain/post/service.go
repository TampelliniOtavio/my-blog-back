package post

import (
	databaseerror "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/error/database-error"
	internalerror "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/error/internal-error"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	ListAllPosts(params *ListAllPostsParams) (*[]Post, error)
	GetPost(xid string, authUserId int64) (*Post, error)
	AddPost(body *AddPostBody, createdBy int64) (*Post, error)
	AddLikeToPost(postXid string, userId int64) (*Post, error)
	RemoveLikeFromPost(postXid string, userId int64) (*Post, error)
	DeletePost(postXid string, userId int64) (*Post, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) ListAllPosts(params *ListAllPostsParams) (*[]Post, error) {
	if params.Queries.Limit < 0 {
		return nil, fiber.NewError(400, "Limit is not valid")
	}

	if params.Queries.Offset < 0 {
		return nil, fiber.NewError(400, "Offset is not valid")
	}

	return s.Repository.GetAllPosts(params)
}

func (s *ServiceImp) AddPost(body *AddPostBody, createdBy int64) (*Post, error) {
	post, err := NewPost(body.Post, createdBy)

	if err != nil {
		return nil, err
	}

	return s.Repository.AddPost(post)
}

func (s *ServiceImp) GetPost(xid string, authUserId int64) (*Post, error) {
	post, err := s.Repository.GetPost(xid, authUserId)

	if err == nil {
		return post, nil
	}

	if databaseerror.IsNotFound(err) {
		return nil, internalerror.NotFound("Post")
	}

	return nil, err
}

func (s *ServiceImp) AddLikeToPost(postXid string, userId int64) (*Post, error) {
	post, err := s.GetPost(postXid, userId)

	if err != nil {
		return nil, internalerror.NotFound("Post")
	}

	err = s.Repository.AddLikeToPost(post, userId)

	if err != nil {
		if err.Error() == "User Already Liked the post" {
			return post, nil
		}

		return nil, err
	}

	post, _ = s.GetPost(postXid, userId)

	return post, nil
}

func (s *ServiceImp) RemoveLikeFromPost(postXid string, userId int64) (*Post, error) {
	post, err := s.GetPost(postXid, userId)

	if err != nil {
		return nil, internalerror.NotFound("Post")
	}

	err = s.Repository.RemoveLikeFromPost(post, userId)

	if err != nil {
		return nil, err
	}

	post, _ = s.GetPost(postXid, userId)

	return post, nil
}

func (s *ServiceImp) DeletePost(postXid string, userId int64) (*Post, error) {
	post, err := s.GetPost(postXid, userId)

	if err != nil {
		return nil, internalerror.NotFound("Post")
	}

	err = s.Repository.DeletePost(post, userId)

	if err != nil {
		return nil, err
	}

	post, _ = s.GetPost(postXid, userId)

	return post, nil
}
