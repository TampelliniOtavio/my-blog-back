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

func (r *PostRepository) AddPost(insertPost *post.Post) (*post.Post, error) {
	var newPost post.Post
	err := r.DB.QueryRowx("INSERT INTO my_blog.posts(xid, post, created_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *", insertPost.Xid, insertPost.Post, insertPost.CreatedBy, insertPost.CreatedAt, insertPost.UpdatedAt).StructScan(&newPost)

	if err != nil {
		return nil, err
	}

	return &newPost, nil
}
