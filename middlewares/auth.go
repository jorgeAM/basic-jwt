package middlewares

import (
	"fmt"
	"net/http"
)

// AuthMiddleware handle jwt in authorization header
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hi from middleware")
		next.ServeHTTP(w, r)
	})
}
