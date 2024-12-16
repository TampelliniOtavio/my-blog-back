package post

type Post struct {
	Xid       string `validate:"required"`
	Post      string `validate:"required"`
	CreatedBy string `validate:"required" db:"created_by"`
	CreatedAt string `validate:"required,datetime" db:"created_at"`
	UpdatedAt string `validate:"required,datetime" db:"updated_at"`
}
