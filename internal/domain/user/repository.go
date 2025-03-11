package user

import "github.com/TampelliniOtavio/my-blog-back/internal/domain/post"

type GetPostsByUsernameParams struct {
	UserId   int64
	Username string
	Limit    int64
	Offset   int64
}

type Repository interface {
	GetByUsername(username string) (*User, error)
	GetPostsByUsername(params *GetPostsByUsernameParams) (*[]post.Post, error)
	CreateUser(user *User) (*User, error)
}
