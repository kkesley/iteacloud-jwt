package jwtemail

//TokenRequest request object for generating JWT
type TokenRequest struct {
	Email        string `json:"email,omitempty"`
	ClientID     uint64 `json:"client_id,omitempty"`
	Username     string `json:"username,omitempty"`
	ClientPrefix string `json:"client_prefix,omitempty"`
	IsRoot       bool   `json:"is_root"`
}
