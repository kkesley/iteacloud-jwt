package jwtforgetpassword

//TokenRequest request object for generating JWT
type TokenRequest struct {
	Username     string `json:"Username"`
	ClientID     uint64 `json:"ClientID"`
	ClientPrefix string `json:"ClientPrefix,omitempty"`
	IsRoot       bool   `json:"IsRoot"`
}
