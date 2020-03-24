package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/jwt/controllers"
	"github.com/jorgeAM/jwt/middlewares"
)

// InitializeRoutes initialize routes
func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/signin", controllers.Signin).Methods("POST")
	r.HandleFunc("/welcome", middlewares.AuthMiddleware(controllers.Welcome)).Methods("GET")
	return r
}