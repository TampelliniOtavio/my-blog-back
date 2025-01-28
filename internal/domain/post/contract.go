package post

type AddPostBody struct {
	Post string `validate:"required"`
}
