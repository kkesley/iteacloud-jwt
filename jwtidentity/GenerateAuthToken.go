package jwtidentity

import (
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
)

//GenerateAuthToken generate from map to auth token
func GenerateAuthToken(authContext map[string]interface{}) TokenRequest {
	authToken := TokenRequest{}
	iterateAuthContext([]string{"IsRoot", "UserARN", "RoleARN", "ClientID", "ClientName", "FirstName", "LastName", "Username", "Groups", "Permissions", "Device"}, &authToken, authContext)
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
				if fieldType != "[]interface {}" {
					continue
				}
				token.RoleARN = make([]string, 0)
				for _, value := range context[field].([]interface{}) {
					if reflect.TypeOf(value).String() == "string" {
						token.RoleARN = append(token.RoleARN, value.(string))
					}
				}
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
				if fieldType != "[]interface {}" {
					continue
				}
				token.Groups = make([]string, 0)
				for _, value := range context[field].([]interface{}) {
					if reflect.TypeOf(value).String() == "string" {
						token.Groups = append(token.Groups, value.(string))
					}
				}
			case "Device":
				if fieldType != "string" {
					continue
				}
				token.Device = context[field].(string)
			}
		}
	}
}
