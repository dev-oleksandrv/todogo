package auth

import (
	"context"
	"net/http"

	jwtutils "github.com/dev-oleksandrv/internal/jwt-utils"
)

type contextKey string

const userIDContextKey = contextKey("userID")

func GetUserIDContextKey() contextKey {
	return userIDContextKey
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims := &jwtutils.Claims{}
		tkn, err := jwtutils.ParseToken(tokenCookie.Value, claims)
		if err != nil || !tkn.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), GetUserIDContextKey(), claims.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}