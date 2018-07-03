package jwt

//TokenRequest request object for generating JWT
type TokenRequest struct {
	UserARN     string   `json:"user_arn"`
	ClientID    uint64   `json:"client_id"`
	ClientName  string   `json:"client_name"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Username    string   `json:"username"`
	Groups      []string `json:"groups"`
	Permissions []string `json:"permissions"`
}
