package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/services"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	var groupDto dtos.CreateGroupRequestDto

	if err := json.NewDecoder(r.Body).Decode(&groupDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	groupCode, err := services.CreateGroup(groupDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dtos.CreateGroupResponseDto{
		GroupCode: groupCode,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}