package jwt

import jwt "github.com/dgrijalva/jwt-go"

//Token standard JWT Token for an identity
type Token struct {
	UserARN     string   `json:"user_arn"`
	ClientID    uint64   `json:"client_id"`
	ClientName  string   `json:"client_name"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Username    string   `json:"username"`
	Groups      []string `json:"groups"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}
