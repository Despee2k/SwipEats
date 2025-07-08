package routes

import (
	"fmt"
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Setup() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Landing page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to SwipEats!"))
	})

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Protected route example
	// This route requires a valid JWT token to access
	r.With(middlewares.JWTMiddleware).Get("/protected", func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(middlewares.UserIDKey)
		w.Write([]byte("Protected route accessed by user id: " + fmt.Sprintf("%d", userID.(uint))))
	})

	// Mounting sub-routers for different API versions
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/auth", AuthRouter())
		r.With(middlewares.JWTMiddleware).Mount("/group", GroupRouter())
	})

	return r
}