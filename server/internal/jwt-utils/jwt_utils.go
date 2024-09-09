package jwtutils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your-secret-key")
var refreshKey = []byte("your-refresh-key")

type Claims struct { 
	UserID int `json:"userId"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateTokens(userId int, email string) (string, string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	refreshExpirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{UserID: userId, Email: email, RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime)}}
	refreshClaims := &Claims{UserID: userId, Email: email, RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(refreshExpirationTime)}}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}
	refreshTokenString, err := refreshToken.SignedString(refreshKey)
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}

func ParseToken(token string, claims *Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}