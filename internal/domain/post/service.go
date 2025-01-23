package post

import (
	"strings"

	postcontract "github.com/TampelliniOtavio/my-blog-back/internal/contract/post-contract"
	databaseerror "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database-error"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	ListAllPosts(limit int, offset int) (*[]Post, error)
	GetPost(xid string) (*Post, error)
	AddPost(body *postcontract.PostAddPostBody, createdBy int64) (*Post, error)
	AddLikeToPost(postXid string, userId int64) (*Post, error)
	RemoveLikeFromPost(postXid string, userId int64) (*Post, error)
	DeletePost(postXid string, userId int64) (*Post, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) ListAllPosts(limit int, offset int) (*[]Post, error) {
	if limit < 0 {
		return nil, fiber.NewError(400, "Limit is not valid")
	}

	if offset < 0 {
		return nil, fiber.NewError(400, "Offset is not valid")
	}
	return s.Repository.GetAllPosts(limit, offset)
}

func (s *ServiceImp) AddPost(body *postcontract.PostAddPostBody, createdBy int64) (*Post, error) {
	post, err := NewPost(body.Post, createdBy)

	if err != nil {
		return nil, err
	}

	return s.Repository.AddPost(post)
}

func (s *ServiceImp) GetPost(xid string) (*Post, error) {
	post, err := s.Repository.GetPost(xid)

	if err == nil {
		return post, nil
	}

	if err.Error() == databaseerror.NOT_FOUND {
		return nil, fiber.NewError(404, "Post Not Found")
	}

	return nil, err
}

func (s ServiceImp) AddLikeToPost(postXid string, userId int64) (*Post, error) {
	post, err := s.GetPost(postXid)

	if err != nil {
		return nil, fiber.NewError(404, "Post Not Found")
	}

	err = s.Repository.AddLikeToPost(post, userId)

	if err != nil {
		if strings.Index(err.Error(), "likes_post_one_user_per_post") > -1 {
			return post, nil
		}

		return nil, err
	}

	post.LikeCount += 1

	return post, nil
}

func (s ServiceImp) RemoveLikeFromPost(postXid string, userId int64) (*Post, error) {
	post, err := s.GetPost(postXid)

	if err != nil {
		return nil, fiber.NewError(404, "Post Not Found")
	}

	err = s.Repository.RemoveLikeFromPost(post, userId)

	if err != nil {
		if err.Error() == "Liked Post Not Found" {
			return post, nil
		}

		return nil, err
	}

	post.LikeCount -= 1

	return post, nil
}

func (s *ServiceImp) DeletePost(postXid string, userId int64) (*Post, error) {
	post, err := s.GetPost(postXid)

	if err != nil {
		return nil, internalerrors.NotFound("Post")
	}

	err = s.Repository.DeletePost(post, userId)

	if err != nil {
		return nil, err
	}

	post, _ = s.GetPost(postXid)

	return post, nil
}
