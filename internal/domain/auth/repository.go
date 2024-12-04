package auth

type Repository interface{
    GetByUsername(username string) (*User, error)
}
