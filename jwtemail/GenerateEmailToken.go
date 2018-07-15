package jwtemail

import (
	"reflect"
)

//GenerateEmailToken generate from map to Email token
func GenerateEmailToken(emailContext map[string]interface{}) TokenRequest {
	emailToken := TokenRequest{}
	iterateEmailContext([]string{"IsRoot", "ClientPrefix", "ClientID", "Email", "Username"}, &emailToken, emailContext)
	return emailToken
}

func iterateEmailContext(fields []string, token *TokenRequest, context map[string]interface{}) {
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
			case "Email":
				if fieldType != "string" {
					continue
				}
				token.Email = context[field].(string)
			case "Username":
				if fieldType != "string" {
					continue
				}
				token.Username = context[field].(string)
			}
		}
	}
}
