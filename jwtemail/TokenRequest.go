package jwtemail

//TokenRequest request object for generating JWT
type TokenRequest struct {
	Email    string `json:"email"`
	ClientID uint64 `json:"client_id"`
}
