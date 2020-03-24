package main

import (
	"log"
	"net/http"

	"github.com/jorgeAM/jwt/handlers"
)

func main() {
	http.HandleFunc("/signin", handlers.Signin)
	http.HandleFunc("/welcome", handlers.Welcome)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
