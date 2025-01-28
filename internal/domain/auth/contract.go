package auth

type PostLoginBody struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type PostSigninBody struct {
	Email    string `validate:"required,email"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}
