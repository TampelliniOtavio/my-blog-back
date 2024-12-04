package authcontract

type PostLoginBody struct{
    Username string `validate:"required"`
    Password string `validate:"required"`
}
