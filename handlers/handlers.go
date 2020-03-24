package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jorgeAM/jwt/models"
)

// Users map with valid data
var Users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

var key = os.Getenv("JWT_KEY")

// Welcome verify if jwt is valid
func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo gucci"))
}

// Signin get a jwt
func Signin(w http.ResponseWriter, r *http.Request) {
	c := models.Credentials{}
	err := json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something got wrong to parse request"))
		return
	}

	password := Users[c.Username]

	if password != c.Password {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password does not match"))
		return
	}

	claim := &models.Claims{
		Username: c.Username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "test jwt",
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something got wrong to generate jwt"))
		return
	}

	t := models.Token{
		Credentias: c,
		Token:      tokenString,
	}

	bytes, err := json.Marshal(t)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something got wrong to parse jwt to json"))
		return
	}

	w.Write(bytes)
}
