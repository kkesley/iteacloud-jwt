package jwtforgetpassword

import (
	"reflect"
)

//GenerateForgetPasswordToken generate from map to Forget Password token
func GenerateForgetPasswordToken(context map[string]interface{}) TokenRequest {
	forgetPasswordToken := TokenRequest{}
	iterateForgetPasswordContext([]string{"IsRoot", "ClientPrefix", "ClientID", "Username", "Time"}, &forgetPasswordToken, context)
	return forgetPasswordToken
}

func iterateForgetPasswordContext(fields []string, token *TokenRequest, context map[string]interface{}) {
	for _, field := range fields {
		if context[field] != nil {
			fieldType := reflect.TypeOf(context[field]).String()
			switch field {
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
