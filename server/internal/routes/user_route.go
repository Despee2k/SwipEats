package routes

import (
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func UserRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handlers.GetUserHandler)
	r.Patch("/update", handlers.UpdateUserHandler)

	return r;
}