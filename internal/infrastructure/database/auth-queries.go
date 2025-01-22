package database

import (
	"strings"

	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) GetByUsername(username string) (*auth.User, error) {
	var user auth.User
	err := r.db.QueryRowx("SELECT * FROM my_blog.users WHERE username = $1", username).StructScan(&user)

	if err != nil {
		return nil, handleError(err)
	}

	return &user, nil
}

func (r *AuthRepository) CreateUser(user *auth.User) (*auth.User, error) {
	var newUser auth.User
	err := r.db.QueryRowx("INSERT INTO my_blog.users(xid, username, email, password) VALUES ($1, $2, $3, $4) RETURNING *", user.Xid, user.Username, user.Email, user.Password).StructScan(&newUser)

	if err != nil {
		return nil, handleError(err)
	}

	return &newUser, nil
}

func handleError(err error) error {
	errMessage := err.Error()

	if strings.Index(errMessage, "duplicate key value violates unique constraint \"users_username\"") != -1 {
		return fiber.NewError(400, "Username already exists")
	}

	if strings.Index(errMessage, "duplicate key value violates unique constraint \"users_email\"") != -1 {
		return fiber.NewError(400, "Email already exists")
	}

	return err
}
