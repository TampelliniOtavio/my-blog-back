package post

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/databasetypes"
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/formatter"
	internalerrors "github.com/TampelliniOtavio/my-blog-back/internal/internal-errors"
	"github.com/rs/xid"
)

type Post struct {
	Xid           string                   `validate:"required" json:"xid"`
	Post          string                   `validate:"required" json:"post"`
	CreatedBy     int64                    `validate:"required" db:"created_by" json:"-"`
	CreatedByName string                   `db:"username" json:"createdBy"`
	LikeCount     int64                    `db:"like_count" json:"likeCount"`
	CreatedAt     string                   `validate:"required,datetime" db:"created_at" json:"createdAt"`
	UpdatedAt     string                   `validate:"required,datetime" db:"updated_at" json:"updatedAt"`
	DeletedAt     databasetypes.NullString `db:"deleted_at" json:"deletedAt"`
}

type PostLike struct {
	UserId    int64  `validate:"required" json:"-" db:"user_id"`
	UserName  int64  `json:"username"`
	PostXid   string `validate:"required" json:"postXid" db:"post_xid"`
	CreatedAt string `validate:"required,datetime" db:"created_at" json:"createdAt"`
}

func NewPost(post string, createdBy int64) (*Post, error) {
	now := formatter.CurrentTimestamp()

	newPost := Post{
		Xid:       xid.New().String(),
		Post:      post,
		CreatedBy: createdBy,
		LikeCount: 0,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := internalerrors.ValidateStruct(newPost)

	if err != nil {
		return nil, err
	}

	return &newPost, nil
}

func NewPostLike(userId int64, postXid string) *PostLike {
	return &PostLike{
		UserId:    userId,
		PostXid:   postXid,
		CreatedAt: formatter.CurrentTimestamp(),
	}
}
