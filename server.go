package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/jwt/routes"
)

func main() {
	r := routes.InitializeRoutes()
	log.Fatal(http.ListenAndServe(":8000", r))
}
