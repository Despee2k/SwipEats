package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/middlewares"
	"github.com/SwipEats/SwipEats/server/internal/services"
)

func GetRecentMatches(w http.ResponseWriter, r *http.Request) {
	var errorResponse dtos.APIErrorResponse
	var successResponse dtos.APISuccessResponse[[]dtos.GroupRestaurantResponseDto]

	w.Header().Set("Content-Type", "application/json")

	userID := r.Context().Value(middlewares.UserIDKey).(uint)

	matches, err := services.GetUserRecentMatches(userID)
	if err != nil {
		errorResponse.Message = "Failed to retrieve recent matches"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse.Data = matches
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)
}