package user

type GetPostsByUsernameParams struct {
	UserId   int64
	Username string
	Limit    int64
	Offset   int64
}

type Repository interface {
	GetByUsername(username string) (*User, error)
	CreateUser(user *User) (*User, error)
}
