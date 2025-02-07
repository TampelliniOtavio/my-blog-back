package database

import (
	"strings"

	"github.com/TampelliniOtavio/my-blog-back/internal/domain/post"
	databaseerror "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database-error"
	internalerrors "github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/errors/internal-errors"
	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) GetAllPosts(limit int, offset int, authUserId int64) (*[]post.Post, error) {
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
							AND likes.user_id = $3
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
		WHERE posts.deleted_at IS NULL
		LIMIT $1 OFFSET $2`,
		limit,
		offset,
		authUserId,
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
		`with posts as (
			insert into
			my_blog.posts (xid, post, created_by, created_at, updated_at)
			values ($1, $2, $3, $4, $5)
			returning *
		) select
			posts.xid,
			posts.post,
			posts.created_at,
			posts.updated_at,
			posts.like_count,
			posts.deleted_at,
			users.id as created_by,
			users.username
		from posts
		inner join my_blog.users ON users.id = posts.created_by`,
		insertPost.Xid,
		insertPost.Post,
		insertPost.CreatedBy,
		insertPost.CreatedAt,
		insertPost.UpdatedAt,
	).StructScan(&newPost)

	if err != nil {
		return nil, r.handleError(err)
	}

	return &newPost, nil
}

func (r *PostRepository) GetPost(xid string, authUserId int64) (*post.Post, error) {
	var post post.Post

	err := r.db.QueryRowx(
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
							and likes.user_id = $2
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
		WHERE posts.xid = $1`,
		xid,
		authUserId,
	).StructScan(&post)

	if err != nil {
		if databaseerror.IsNotFound(err) {
			return nil, internalerrors.NotFound("Post")
		}

		return nil, err
	}

	return &post, nil
}

func (r *PostRepository) AddLikeToPost(post *post.Post, userId int64) error {
	return r.handleError(WithTransaction(r.db, func(tx *sqlx.Tx) error {
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

		exec, err := tx.Exec(`
		UPDATE
			my_blog.posts
		SET
			like_count = like_count + 1
		WHERE
			xid = $1 AND
			deleted_at IS NULL
		`, post.Xid)

		if rows, err := exec.RowsAffected(); rows == 0 || err != nil {
			return internalerrors.NotFound("Liked Post")
		}

		return err
	}))
}

func (r *PostRepository) RemoveLikeFromPost(post *post.Post, userId int64) error {
	return r.handleError(WithTransaction(r.db, func(tx *sqlx.Tx) error {
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

		exec, err = tx.Exec(`
		UPDATE
			my_blog.posts
		SET
			like_count = like_count - 1
		WHERE
			xid = $1 AND
			deleted_at IS NULL
		`, post.Xid)

		if rows, err := exec.RowsAffected(); rows == 0 || err != nil {
			return internalerrors.NotFound("Liked Post")
		}

		return err
	}))
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
		return r.handleError(err)
	}
	rows, err := result.RowsAffected()

	if rows == 0 {
		return internalerrors.NotFound("Post")
	}

	return r.handleError(err)
}

func (p *PostRepository) handleError(err error) error {
	if err == nil {
		return err
	}

	if strings.Index(err.Error(), "violates foreign key constraint \"posts_users_fk\"") > -1 {
		return internalerrors.NotFound("User")
	}

	if strings.Index(err.Error(), "violates unique constraint \"likes_post_one_user_per_post\"") > -1 {
		return internalerrors.BadRequest("User Already Liked the post")
	}

	return err
}
