package models

import "github.com/dgrijalva/jwt-go"

// Claims struct that will be encoded to JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
