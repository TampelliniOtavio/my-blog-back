package postcontract

type PostAddPostBody struct {
	Post string `validate:"required"`
}
