package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateJWT(test *testing.T) {
	request := TokenRequest{
		UserARN:     "test-arn",
		ClientID:    1,
		ClientName:  "test",
		FirstName:   "hello",
		LastName:    "there",
		Username:    "abc",
		Groups:      []string{"FLUFFY"},
		Permissions: []string{"*"},
	}
	jwt, err := GenerateJWT(request, time.Minute*1, "SuperSecretKeyOnlyForThisApplication")
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	fmt.Println(jwt)
}

func TestGenerateJWTWithIssuer(test *testing.T) {
	request := TokenRequest{
		UserARN:     "test-arn",
		ClientID:    1,
		ClientName:  "test",
		FirstName:   "hello",
		LastName:    "there",
		Username:    "abc",
		Groups:      []string{"FLUFFY"},
		Permissions: []string{"*"},
	}
	jwt, err := GenerateJWTWithIssuer(request, time.Minute*1, "SuperSecretKeyOnlyForThisApplication", "pss")
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	fmt.Println(jwt)
}
