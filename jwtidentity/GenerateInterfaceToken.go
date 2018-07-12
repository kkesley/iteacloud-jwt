package jwtidentity

import "strings"

//GenerateInterfaceToken generate from token to map
func GenerateInterfaceToken(token TokenRequest) map[string]interface{} {
	return map[string]interface{}{
		"IsRoot":      token.IsRoot,
		"UserARN":     token.UserARN,
		"RoleARN":     token.RoleARN,
		"ClientID":    token.ClientID,
		"FirstName":   token.FirstName,
		"LastName":    token.LastName,
		"Username":    token.Username,
		"Groups":      strings.Join(token.Groups, ","),
		"Permissions": strings.Join(token.Permissions, ","),
	}
}
