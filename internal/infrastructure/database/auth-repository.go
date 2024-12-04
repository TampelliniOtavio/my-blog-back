package database

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/auth"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct{
    DB *sqlx.DB
}

func (r *AuthRepository) GetByUsername(username string) (*auth.User, error) {
    var user auth.User
    err := r.DB.QueryRowx("SELECT * FROM my_blog.users WHERE username = $1", username).StructScan(&user)

    if err != nil {
        return nil, err
    }

    return &user, nil
}
