package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jorgeAM/jwt/models"
)

// AuthMiddleware handle jwt in authorization header
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("authorization header must be sent"))
			return
		}

		bearerToken = strings.Replace(bearerToken, "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(bearerToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		claims, ok := token.Claims.(*models.Claims)

		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token"))
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
