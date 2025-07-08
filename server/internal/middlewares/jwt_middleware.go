package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/SwipEats/SwipEats/server/internal/utils"
)

type contextKey string

const UserEmailKey contextKey = "userEmail"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		user, err := utils.ValidateJWT(token)

		if err != nil {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserEmailKey, user.Email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}