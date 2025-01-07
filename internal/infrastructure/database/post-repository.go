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
	err := r.DB.Select(
		&posts,
		"SELECT "+
			"posts.xid, "+
			"posts.post, "+
			"posts.created_at, "+
			"posts.updated_at, "+
			"users.username "+
			"FROM my_blog.posts AS posts "+
			"INNER JOIN my_blog.users AS users ON users.id = posts.created_by "+
			"LIMIT $1 OFFSET $2",
		limit,
		offset,
	)

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
	err := r.DB.QueryRowx(
		"with posts as ( "+
			"	insert into "+
			"	my_blog.posts (xid, post, created_by, created_at, updated_at) "+
			"	values ($1, $2, $3, $4, $5) "+
			"	returning * "+
			") select posts.xid, posts.post, posts.created_at, posts.updated_at, users.username from posts "+
			"inner join my_blog.users ON users.id = posts.created_by;",
		insertPost.Xid,
		insertPost.Post,
		insertPost.CreatedBy,
		insertPost.CreatedAt,
		insertPost.UpdatedAt,
	).StructScan(&newPost)

	if err != nil {
		return nil, err
	}

	return &newPost, nil
}

func (r *PostRepository) GetPost(xid string) (*post.Post, error) {
	var post post.Post

	err := r.DB.QueryRowx(
		"SELECT "+
			"posts.xid, "+
			"posts.post, "+
			"posts.created_at, "+
			"posts.updated_at, "+
			"users.username "+
			"FROM my_blog.posts AS posts "+
			"INNER JOIN my_blog.users AS users ON users.id = posts.created_by "+
		"WHERE posts.xid = $1",
		xid,
	).StructScan(&post)

	if err != nil {
		return nil, err
	}

	return &post, nil
}
