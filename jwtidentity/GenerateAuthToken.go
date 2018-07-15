package jwtidentity

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
)

//GenerateAuthToken generate from map to auth token
func GenerateAuthToken(authContext map[string]interface{}) TokenRequest {
	authToken := TokenRequest{}
	iterateAuthContext([]string{"IsRoot", "UserARN", "RoleARN", "ClientID", "ClientName", "FirstName", "LastName", "Username", "Groups", "Permissions"}, &authToken, authContext)
	return authToken
}

func iterateAuthContext(fields []string, token *TokenRequest, context map[string]interface{}) {
	for _, field := range fields {
		if context[field] != nil {
			fieldType := reflect.TypeOf(context[field]).String()
			switch field {
			case "IsRoot":
				if fieldType != "bool" {
					continue
				}
				token.IsRoot = context[field].(bool)
			case "UserARN":
				if fieldType != "string" {
					continue
				}
				token.UserARN = context[field].(string)
			case "RoleARN":
				if fieldType != "string" {
					continue
				}
				token.RoleARN = context[field].(string)
			case "ClientID":
				if fieldType != "float64" {
					continue
				}
				token.ClientID = uint64(context[field].(float64))
			case "ClientName":
				if fieldType != "string" {
					continue
				}
				token.ClientName = context[field].(string)
			case "FirstName":
				if fieldType != "string" {
					continue
				}
				token.FirstName = aws.String(context[field].(string))
			case "LastName":
				if fieldType != "string" {
					continue
				}
				token.LastName = aws.String(context[field].(string))
			case "Username":
				if fieldType != "string" {
					continue
				}
				token.Username = context[field].(string)
			case "Groups":
				if fieldType != "string" {
					continue
				}
				token.Groups = strings.Split(context[field].(string), ",")
			case "Permissions":
				if fieldType != "string" {
					continue
				}
				token.Permissions = strings.Split(context[field].(string), ",")
			}
		}
	}
}
