package routes

import (
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func GroupRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/{group_code}/members", handlers.GetGroupMembersHandler)

	r.Post("/create", handlers.CreateGroupHandler)
	r.Post("/{group_code}/join", handlers.JoinGroupHandler)
	r.Post("/{group_code}/leave", handlers.LeaveGroupHandler)

	return r
}