package jwtemail

//TokenRequest request object for generating JWT
type TokenRequest struct {
	Email        string `json:"Email,omitempty"`
	ClientID     uint64 `json:"ClientID,omitempty"`
	Username     string `json:"Username,omitempty"`
	ClientPrefix string `json:"ClientPrefix,omitempty"`
	IsRoot       bool   `json:"IsRoot"`
}
