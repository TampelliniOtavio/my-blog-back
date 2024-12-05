package authcontract

type PostSigninBody struct{
    Email    string `validate:"required,email"`
    Username string `validate:"required"`
    Password string `validate:"required"`
}
