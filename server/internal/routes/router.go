package routes

import (
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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to SwipEats!"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.With(middlewares.JWTMiddleware).Get("/protected", func(w http.ResponseWriter, r *http.Request) {
		userEmail := r.Context().Value(middlewares.UserEmailKey)
		w.Write([]byte("Protected route accessed by user email: " + userEmail.(string)))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/auth", AuthRouter())
	})

	return r
}