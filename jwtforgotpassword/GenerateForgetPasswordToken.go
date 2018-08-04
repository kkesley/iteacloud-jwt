package jwtforgotpassword

import (
	"reflect"
)

//GenerateForgotPasswordToken generate from map to Forgot Password token
func GenerateForgotPasswordToken(context map[string]interface{}) TokenRequest {
	forgotPasswordToken := TokenRequest{}
	iterateForgotPasswordContext([]string{"IsRoot", "ClientPrefix", "ClientID", "Username", "Time", "UserARN"}, &forgotPasswordToken, context)
	return forgotPasswordToken
}

func iterateForgotPasswordContext(fields []string, token *TokenRequest, context map[string]interface{}) {
	for _, field := range fields {
		if context[field] != nil {
			fieldType := reflect.TypeOf(context[field]).String()
			switch field {
			case "UserARN":
				if fieldType != "string" {
					continue
				}
				token.UserARN = context[field].(string)
			case "IsRoot":
				if fieldType != "bool" {
					continue
				}
				token.IsRoot = context[field].(bool)
			case "ClientPrefix":
				if fieldType != "string" {
					continue
				}
				token.ClientPrefix = context[field].(string)
			case "ClientID":
				if fieldType != "float64" {
					continue
				}
				token.ClientID = uint64(context[field].(float64))
			case "Username":
				if fieldType != "string" {
					continue
				}
				token.Username = context[field].(string)
			case "Time":
				if fieldType != "float64" {
					continue
				}
				token.Time = int64(context[field].(float64))
			}
		}
	}
}
