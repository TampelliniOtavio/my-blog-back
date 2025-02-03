package user

type Repository interface {
	GetByUsername(username string) (*User, error)
	CreateUser(user *User) (*User, error)
}
