package jwtidentity

import (
	"reflect"
	"strconv"
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
			value := context[field].(string)
			fieldType := reflect.TypeOf(context[field]).String()
			if fieldType != "string" {
				continue
			}
			switch field {
			case "IsRoot":
				if isRoot, err := strconv.ParseBool(value); err == nil {
					token.IsRoot = isRoot
				}
			case "UserARN":
				token.UserARN = value
			case "RoleARN":
				token.RoleARN = value
			case "ClientID":
				if clientID, err := strconv.ParseUint(value, 10, 64); err == nil {
					token.ClientID = clientID
				}
			case "ClientName":
				token.ClientName = value
			case "FirstName":
				token.FirstName = aws.String(value)
			case "LastName":
				token.LastName = aws.String(value)
			case "Username":
				token.Username = value
			case "Groups":
				token.Groups = strings.Split(value, ",")
			case "Permissions":
				token.Permissions = strings.Split(value, ",")
			}
		}
	}
}
