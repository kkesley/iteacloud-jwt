package jwtforgotpassword

//TokenRequest request object for generating JWT
type TokenRequest struct {
	UserARN      string `json:"UserARN"`
	Username     string `json:"Username"`
	ClientID     uint64 `json:"ClientID"`
	ClientPrefix string `json:"ClientPrefix,omitempty"`
	IsRoot       bool   `json:"IsRoot"`
	Time         int64  `json:"Time"`
}
