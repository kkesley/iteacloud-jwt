package jwtidentity

//TokenRequest request object for generating JWT
type TokenRequest struct {
	IsRoot     bool     `json:"IsRoot"`
	UserARN    string   `json:"UserARN"`
	RoleARN    []string `json:"RoleARN"`
	ClientID   uint64   `json:"ClientID"`
	ClientARN  string   `json:"ClientARN"`
	ClientName string   `json:"ClientName"`
	FirstName  *string  `json:"FirstName"`
	LastName   *string  `json:"LastName"`
	Username   string   `json:"Username"`
	Groups     []string `json:"Groups"`
	Device     string   `json:"Device"`
	IsPublic   bool     `json:"IsPublic"`
}
