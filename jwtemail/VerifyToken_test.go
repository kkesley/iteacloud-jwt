package jwtemail

import (
	"fmt"
	"testing"
	"time"
)

func TestVerifier(test *testing.T) {
	request := TokenRequest{
		Email:    "kk@gmail.com",
		ClientID: 1,
	}
	jwt, err := GenerateJWT(request, time.Hour*1, "SuperSecretKeyOnlyForThisApplication")
	claims, err := VerifyToken(jwt, "SuperSecretKeyOnlyForThisApplication")
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	fmt.Println(claims)
}

func TestVerifierWithoutExpiry(test *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2FybiI6InRlc3QtYXJuIiwiY2xpZW50X2lkIjoxLCJjbGllbnRfbmFtZSI6InRlc3QiLCJmaXJzdF9uYW1lIjoiaGVsbG8iLCJsYXN0X25hbWUiOiJ0aGVyZSIsInVzZXJuYW1lIjoiYWJjIiwiZ3JvdXBzIjpbIkZMVUZGWSJdLCJwZXJtaXNzaW9ucyI6WyIqIl0sImV4cCI6MTUzMDU3NzExOSwiaWF0IjoxNTMwNTc3MDU5LCJpc3MiOiJpdGVhY2xvdWQiLCJuYmYiOjE1MzA1NzcwNTl9.TYyg61zWID-PmcHudaXAJBCvwSnudrC8KkYR-WbigXc0Ei-GWaZnJz5Ivv1sXFmc5E-k9nWj7VqHfcFHeEFX7g"
	claims, err := VerifyTokenWithoutExp(token, "SuperSecretKeyOnlyForThisApplication")
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	fmt.Println(claims)
}
