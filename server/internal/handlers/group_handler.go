package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/middlewares"
	"github.com/SwipEats/SwipEats/server/internal/services"
	"github.com/SwipEats/SwipEats/server/internal/utils"
	"github.com/go-chi/chi/v5"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	var groupDto dtos.CreateGroupRequestDto

	if err := json.NewDecoder(r.Body).Decode(&groupDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(groupDto); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middlewares.UserIDKey).(uint)

	groupCode, err := services.CreateGroup(groupDto, userID)
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

func JoinGroupHandler(w http.ResponseWriter, r *http.Request) {
	groupCode := chi.URLParam(r, "group_code")

	userID := r.Context().Value(middlewares.UserIDKey).(uint)

	response, err := services.JoinGroup(groupCode, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func LeaveGroupHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserIDKey).(uint)
	groupCode := chi.URLParam(r, "group_code")

	if groupCode == "" {
		http.Error(w, "Group code is required", http.StatusBadRequest)
		return
	}

	err := services.LeaveGroup(userID, groupCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetGroupMembersHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserIDKey).(uint)
	groupCode := chi.URLParam(r, "group_code")

	if groupCode == "" {
		http.Error(w, "Group code is required", http.StatusBadRequest)
		return
	}

	members, err := services.GetGroupMembers(groupCode, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(members); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}