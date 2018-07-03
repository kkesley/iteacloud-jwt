package jwtidentity

import (
	"errors"
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

//VerifyToken verifies jwt token
func VerifyToken(tokenString string, key string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(key), nil
	})
	if err != nil {
		return jwt.MapClaims{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return jwt.MapClaims{}, errors.New("Invalid token")
}

//VerifyTokenWithoutExp verifies jwt token without expiry
func VerifyTokenWithoutExp(tokenString string, key string) (jwt.MapClaims, error) {
	var p jwt.Parser
	token, parts, err := p.ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return jwt.MapClaims{}, err
	}
	token.Signature = parts[2]
	if err = token.Method.Verify(strings.Join(parts[0:2], "."), token.Signature, []byte(key)); err != nil {
		return jwt.MapClaims{}, err
	}

	if claims := token.Claims.(jwt.MapClaims); true {
		return claims, nil
	}
	return jwt.MapClaims{}, errors.New("Invalid token")
}
