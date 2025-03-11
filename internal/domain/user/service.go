package user

import "github.com/TampelliniOtavio/my-blog-back/internal/domain/post"

type Service interface {
	GetByUsername(username string) (*User, error)
	GetPostsByUsername(loggedUserId int64, username string, limit int64, offset int64) (*[]post.Post, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) GetByUsername(username string) (*User, error) {
	return s.Repository.GetByUsername(username)
}

func (s *ServiceImp) GetPostsByUsername(loggedUserId int64, username string, limit int64, offset int64) (*[]post.Post, error) {
	return s.Repository.GetPostsByUsername(&GetPostsByUsernameParams{
		UserId: loggedUserId,
		Username: username,
		Limit: limit,
		Offset: offset,
	})
}
