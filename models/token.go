package models

// Token that represent JWT
type Token struct {
	Credentias Credentials `json:"credentials"`
	Token      string      `json:"token"`
}
