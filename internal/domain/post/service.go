package post

import (
	postcontract "github.com/TampelliniOtavio/my-blog-back/internal/contract/post-contract"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	ListAllPosts(limit int, offset int) (*[]Post, error)
	AddPost(body *postcontract.PostAddPostBody, createdBy int64) (*Post, error)
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
