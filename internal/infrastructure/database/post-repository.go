package database

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	DB *sqlx.DB
}

func (r *PostRepository) GetAllPosts(limit int, offset int) (*[]post.Post, error) {
	var posts []post.Post
	err := r.DB.Select(&posts, "SELECT * FROM my_blog.posts LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		return nil, err
	}

	if posts == nil {
		return &[]post.Post{}, nil
	}

	return &posts, nil
}
