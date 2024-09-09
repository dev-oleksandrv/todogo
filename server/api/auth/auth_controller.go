package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dev-oleksandrv/db"
	"github.com/dev-oleksandrv/internal/response"
)

type AuthController struct {
	authService *AuthService
}

func NewAuthController(authService *AuthService) *AuthController {
	return &AuthController{authService}
}

func (c *AuthController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userData db.User
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		log.Printf("Error while parsing request: %v", err)
		response.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "Invalid Request payload"},
		)
		return 
	}

	_, err := c.authService.RegisterUser(userData.Email, userData.Password)
	if err != nil {
		response.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "Cannot register user with such credentials"},
		)
		return 
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (c *AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userData db.User
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		log.Printf("Error while parsing request: %v", err)
		response.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "Invalid Request payload"},
		)
		return 
	}

	accessToken, refreshToken, err := c.authService.LoginUser(userData.Email, userData.Password)
	if err != nil {
		response.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "Cannot login user with such credentials"},
		)
		return 
	}

	http.SetCookie(w, &http.Cookie{Name: "access_token", Value: accessToken, Expires: time.Now().Add(15 * time.Minute)})
	http.SetCookie(w, &http.Cookie{Name: "refresh_token", Value: refreshToken, Expires: time.Now().Add(24 * time.Hour)})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}