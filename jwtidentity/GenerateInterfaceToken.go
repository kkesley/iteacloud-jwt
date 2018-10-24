package jwtidentity

import (
	"strconv"
	"strings"
)

//GenerateInterfaceToken generate from token to map
func GenerateInterfaceToken(token TokenRequest) map[string]interface{} {
	return map[string]interface{}{
		"IsRoot":     strconv.FormatBool(token.IsRoot),
		"UserARN":    token.UserARN,
		"RoleARN":    strings.Join(token.RoleARN, ","),
		"ClientID":   token.ClientID,
		"ClientARN":  token.ClientARN,
		"ClientName": token.ClientName,
		"FirstName":  token.FirstName,
		"LastName":   token.LastName,
		"Username":   token.Username,
		"Groups":     strings.Join(token.Groups, ","),
		"Device":     token.Device,
		"IsPublic":   token.IsPublic,
	}
}
