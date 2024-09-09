package auth

import (
	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router, authController *AuthController) {
	listRouter := router.PathPrefix("/api/auth").Subrouter()

	listRouter.HandleFunc("/register", authController.RegisterUser).Methods("POST")
	listRouter.HandleFunc("/login", authController.LoginUser).Methods("POST")
}