package jwtidentity

import (
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
)

func TestGenerateJWT(test *testing.T) {
	request := TokenRequest{
		IsRoot:     true,
		UserARN:    "test-arn",
		ClientID:   1,
		ClientName: "test",
		FirstName:  aws.String("hello"),
		LastName:   aws.String("there"),
		Username:   "abc",
		Groups:     []string{"FLUFFY"},
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
		UserARN:    "test-arn",
		ClientID:   1,
		ClientARN:  "client-arn",
		ClientName: "test",
		FirstName:  aws.String("hello"),
		LastName:   aws.String("there"),
		Username:   "abc",
		Groups:     []string{"FLUFFY"},
	}
	jwt, err := GenerateJWTWithIssuer(request, time.Minute*1, "SuperSecretKeyOnlyForThisApplication", "pss")
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	fmt.Println(jwt)
}
