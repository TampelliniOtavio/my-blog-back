package database

import (
	"strings"

	"github.com/TampelliniOtavio/my-blog-back/internal/domain/user"
	databaseerror "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/error/database-error"
	internalerror "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/error/internal-error"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetByUsername(username string) (*user.User, error) {
	var user user.User
	err := r.db.QueryRowx("SELECT * FROM my_blog.users WHERE username = $1", username).StructScan(&user)

	if err != nil {
		return nil, r.handleError(err)
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(createUser *user.User) (*user.User, error) {
	var newUser user.User
	err := r.db.QueryRowx("INSERT INTO my_blog.users(xid, username, email, password) VALUES ($1, $2, $3, $4) RETURNING *", createUser.Xid, createUser.Username, createUser.Email, createUser.Password).StructScan(&newUser)

	if err != nil {
		return nil, r.handleError(err)
	}

	return &newUser, nil
}

func (r *UserRepository) handleError(err error) error {
	errMessage := err.Error()

	if strings.Index(errMessage, "duplicate key value violates unique constraint \"users_username\"") != -1 {
		return fiber.NewError(400, "Username already exists")
	}

	if strings.Index(errMessage, "duplicate key value violates unique constraint \"users_email\"") != -1 {
		return fiber.NewError(400, "Email already exists")
	}

	if databaseerror.IsNotFound(err) {
		return internalerror.NotFound("User")
	}

	return err
}
