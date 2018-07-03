package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT generates a jwt token in form of a string
func GenerateJWT(request TokenRequest, duration time.Duration, key string) (string, error) {
	claims := Token{
		request.UserARN,
		request.ClientID,
		request.ClientName,
		request.FirstName,
		request.LastName,
		request.Username,
		request.Groups,
		request.Permissions,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			Issuer:    "iteacloud",
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(key))

	return tokenString, err
}
