package auth

type AuthClaims struct {
	Xid      string `json:"xid"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
