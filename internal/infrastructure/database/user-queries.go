package database

import (
	"strings"

	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
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

func (r *UserRepository) GetPostsByUsername(params *user.GetPostsByUsernameParams) (*[]post.Post, error) {
	var posts []post.Post

	err := r.db.Select(
		&posts,
		`SELECT
			posts.xid,
			posts.post,
			posts.created_at,
			posts.updated_at,
			posts.like_count,
			posts.deleted_at,
			users.id as created_by,
			CASE
				WHEN
					COALESCE(
						(SELECT
							COUNT(*)
						FROM
							my_blog.likes_post AS likes
						WHERE
							likes.post_xid = posts.xid
							AND likes.user_id = $4
						)
						, 0
					) > 0
				THEN true
			ELSE
				false
			END AS is_liked_by_user,
			users.username
		FROM my_blog.posts AS posts
		INNER JOIN my_blog.users AS users ON users.id = posts.created_by
		WHERE
			posts.deleted_at IS NULL
			AND users.username = $1
		LIMIT $2 OFFSET $3`,
		params.Username,
		params.Limit,
		params.Offset,
		params.UserId,
	)

	if err != nil {
		return nil, err
	}

	if posts == nil {
		return &[]post.Post{}, nil
	}

	return &posts, nil
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
