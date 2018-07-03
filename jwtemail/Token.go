package jwtemail

import jwt "github.com/dgrijalva/jwt-go"

//Token standard JWT Token for an identity
type Token struct {
	TokenRequest
	jwt.StandardClaims
}
