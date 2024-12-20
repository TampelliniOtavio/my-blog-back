package post

import (
	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/formatter"
	internalerrors "github.com/TampelliniOtavio/my-blog-back/internal/internal-errors"
	"github.com/rs/xid"
)

type Post struct {
	Xid       string `validate:"required"`
	Post      string `validate:"required"`
	CreatedBy int64  `validate:"required" db:"created_by"`
	CreatedAt string `validate:"required,datetime" db:"created_at"`
	UpdatedAt string `validate:"required,datetime" db:"updated_at"`
}

func NewPost(post string, createdBy int64) (*Post, error) {
	now := formatter.CurrentTimestamp()

	newPost := Post{
		Xid:       xid.New().String(),
		Post:      post,
		CreatedBy: createdBy,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := internalerrors.ValidateStruct(newPost)

	if err != nil {
		return nil, err
	}

	return &newPost, nil
}
