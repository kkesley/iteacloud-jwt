package jwtforgotpassword

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT generates a jwt token in form of a string
func GenerateJWT(request TokenRequest, duration time.Duration, key string) (string, error) {
	return GenerateJWTWithIssuer(request, duration, key, "iteacloud")
}

//GenerateJWTWithIssuer generates a jwt token in form of a string with an issuer
func GenerateJWTWithIssuer(request TokenRequest, duration time.Duration, key string, issuer string) (string, error) {
	claims := Token{
		TokenRequest{
			request.UserARN,
			request.Username,
			request.ClientID,
			request.ClientPrefix,
			request.IsRoot,
			request.Time,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			Issuer:    issuer,
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(key))

	return tokenString, err
}
