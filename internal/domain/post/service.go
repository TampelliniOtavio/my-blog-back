package post

import (
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	ListAllPosts(limit int, offset int) (*[]Post, error)
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
