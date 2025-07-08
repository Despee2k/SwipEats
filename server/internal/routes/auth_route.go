package routes

import (
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func AuthRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/signup", handlers.SignupHandler)
	r.Post("/login", handlers.LoginHandler)

	return r
}