package database

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	internalerrors "github.com/TampelliniOtavio/my-blog-back/internal/internal-errors"
	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) GetAllPosts(limit int, offset int) (*[]post.Post, error) {
	var posts []post.Post
	err := r.db.Select(
		&posts,
		"SELECT "+
			"posts.xid, "+
			"posts.post, "+
			"posts.created_at, "+
			"posts.updated_at, "+
			"posts.like_count, "+
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
	err := r.db.QueryRowx(
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

	err := r.db.QueryRowx(
		"SELECT "+
			"posts.xid, "+
			"posts.post, "+
			"posts.created_at, "+
			"posts.updated_at, "+
			"posts.like_count, "+
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

func (r *PostRepository) AddLikeToPost(post *post.Post, userId int64) error {
	return WithTransaction(r.db, func(tx *sqlx.Tx) error {
		_, err := tx.Exec(`
		INSERT INTO
			my_blog.likes_post
			(user_id, post_xid)
		values
			($1, $2)
		`, userId, post.Xid)

		if err != nil {
			return err
		}

		_, err = tx.Exec(`
		UPDATE
			my_blog.posts
		SET
			like_count = like_count + 1
		WHERE
			xid = $1
		`, post.Xid)

		return err
	})
}

func (r *PostRepository) RemoveLikeFromPost(post *post.Post, userId int64) error {
	return WithTransaction(r.db, func(tx *sqlx.Tx) error {
		exec, err := tx.Exec(`
		DELETE FROM
			my_blog.likes_post
		WHERE
			user_id = $1 AND
			post_xid = $2
		`, userId, post.Xid)

		if rows, err := exec.RowsAffected(); rows == 0 || err != nil {
			return internalerrors.NotFound("Liked Post")
		}

		if err != nil {
			return err
		}

		_, err = tx.Exec(`
		UPDATE
			my_blog.posts
		SET
			like_count = like_count - 1
		WHERE
			xid = $1
		`, post.Xid)

		return err
	})
}

func (r *PostRepository) DeletePost(post *post.Post, userId int64) error {
	result, err := r.db.Exec(`
	UPDATE
		my_blog.posts
	SET
		deleted_at = now()
	WHERE
		xid = $1 AND
		created_by = $2 AND
		deleted_at IS NULL
	`, post.Xid, userId)

	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()

	if rows == 0 {
		return internalerrors.NotFound("Post")
	}

	return err
}
