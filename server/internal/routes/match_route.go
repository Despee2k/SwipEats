package routes

import (
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func MatchRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handlers.GetRecentMatches)

	return r
}