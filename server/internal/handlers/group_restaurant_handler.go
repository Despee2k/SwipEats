package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/services"
	"github.com/SwipEats/SwipEats/server/internal/types"
)

func GroupRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	var errorResponse dtos.APIErrorResponse
	var successResponse dtos.APISuccessResponse[[]dtos.GroupRestaurantResponseDto]
	encoder := json.NewEncoder(w)

	groupCode := r.URL.Query().Get("group_code")

	if groupCode == "" {
		errorResponse.Message = "Group code is required"
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(errorResponse)
		return
	}

	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		errorResponse.Message = "Error fetching group"
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(errorResponse)
		return
	}
	if group == nil {
		errorResponse.Message = "Group not found"
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(errorResponse)
		return
	}
	if group.GroupStatus == types.GroupStatusWaiting {
		errorResponse.Message = "Group is still waiting for members"
		w.WriteHeader(http.StatusForbidden)
		encoder.Encode(errorResponse)
		return
	}

	groupRestaurants, err := services.GetGroupRestaurantsByGroupCode(groupCode)
	if err != nil {
		errorResponse.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(errorResponse)
		return
	}

	successResponse.Data = groupRestaurants
	successResponse.Message = "Group restaurants fetched successfully"
	w.WriteHeader(http.StatusOK)
	encoder.Encode(successResponse)
}