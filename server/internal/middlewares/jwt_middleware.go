package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/SwipEats/SwipEats/server/internal/utils"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
)

type contextKey string

const UserIDKey contextKey = "userID"

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

		existingUser, err := repositories.GetUserByEmail(user.Email)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}


		ctx := context.WithValue(r.Context(), UserIDKey, existingUser.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}