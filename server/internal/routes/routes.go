package routes

import (
	"net/http"
    "github.com/go-chi/chi/v5"
)

func Setup() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to SwipEats!"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// r.Route("/api/v1", func(api chi.Router) {

	// });

	return r
}