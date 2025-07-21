package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/middlewares"
	"github.com/SwipEats/SwipEats/server/internal/services"
	"github.com/SwipEats/SwipEats/server/internal/utils"
	"github.com/go-chi/chi/v5"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	var groupDto dtos.CreateGroupRequestDto
	var errorResponse dtos.APIErrorResponse
	var successResponse dtos.APISuccessResponse[dtos.CreateGroupResponseDto]

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&groupDto); err != nil {
		errorResponse.Message = "Invalid request body"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if err := utils.Validate.Struct(groupDto); err != nil {
		message, details := utils.GetValidationErrorDetails(err)

		errorResponse.Message = message
		errorResponse.Details = details

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	userID := r.Context().Value(middlewares.UserIDKey).(uint)

	groupCode, err := services.CreateGroup(groupDto, userID)
	if err != nil {
		switch err {
			case errors.ErrUnableToGenerateGroupCode:
				errorResponse.Message = "Failed to generate a unique group code"
				w.WriteHeader(http.StatusInternalServerError)
			case errors.ErrGroupNotFound:
				errorResponse.Message = "Group not found"
				w.WriteHeader(http.StatusNotFound)
			default:
				errorResponse.Message = "Failed to create group"
				w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := dtos.CreateGroupResponseDto{
		GroupCode: groupCode,
	}

	successResponse.Message = "Group created successfully"
	successResponse.Data = response

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successResponse)
}

func JoinGroupHandler(w http.ResponseWriter, r *http.Request) {
	var errorResponse dtos.APIErrorResponse
	var successResponse dtos.APISuccessResponse[dtos.JoinGroupResponseDto]

	groupCode := chi.URLParam(r, "group_code")

	userID := r.Context().Value(middlewares.UserIDKey).(uint)

	w.Header().Set("Content-Type", "application/json")

	response, err := services.JoinGroup(groupCode, userID)
	if err != nil {
		switch err {
			case errors.ErrGroupNotFound:
				errorResponse.Message = "Group not found"
				w.WriteHeader(http.StatusNotFound)
			case errors.ErrUserAlreadyInGroup:
				errorResponse.Message = "User is already a member of the group"
				w.WriteHeader(http.StatusConflict)
			default:
				errorResponse.Message = "Failed to join group"
				w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse.Message = "Group joined successfully"
	successResponse.Data = *response

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)
}

func GetUserGroupsHandler(w http.ResponseWriter, r *http.Request) {
	var errorResponse dtos.APIErrorResponse
	var successResponse dtos.APISuccessResponse[[]dtos.GetGroupResponseDto]

	userID := r.Context().Value(middlewares.UserIDKey).(uint)

	w.Header().Set("Content-Type", "application/json")

	groups, err := services.GetUserGroups(userID)
	if err != nil {
		errorResponse.Message = "Failed to retrieve user groups"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse.Message = "User groups retrieved successfully"
	successResponse.Data = groups

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)
}

func LeaveGroupHandler(w http.ResponseWriter, r *http.Request) {
	var errorResponse dtos.APIErrorResponse
	var successResponse dtos.APISuccessResponse[any]

	userID := r.Context().Value(middlewares.UserIDKey).(uint)
	groupCode := chi.URLParam(r, "group_code")

	if groupCode == "" {
		errorResponse.Message = "Group code is required"
		json.NewEncoder(w).Encode(errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := services.LeaveGroup(userID, groupCode)
	if err != nil {
		switch err {
			case errors.ErrGroupNotFound:
				errorResponse.Message = "Group not found"
				w.WriteHeader(http.StatusNotFound)
			case errors.ErrUserNotInGroup:
				errorResponse.Message = "User is not a member of the group"
				w.WriteHeader(http.StatusBadRequest)
			default:
				errorResponse.Message = "Failed to leave group"
				w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse.Message = "Left group successfully"

	w.WriteHeader(http.StatusNoContent)
json.NewEncoder(w).Encode(successResponse)
}

func GetGroupMembersHandler(w http.ResponseWriter, r *http.Request) {
	var errorResponse dtos.APIErrorResponse
	var successResponse dtos.APISuccessResponse[[]dtos.UserMembershipResponseDto]

	userID := r.Context().Value(middlewares.UserIDKey).(uint)
	groupCode := chi.URLParam(r, "group_code")
	
	w.Header().Set("Content-Type", "application/json")
	
	if groupCode == "" {
		errorResponse.Message = "Group code is required"
		json.NewEncoder(w).Encode(errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	members, _, err := services.GetGroupMembers(groupCode, userID)
	if err != nil {
		switch err {
			case errors.ErrGroupNotFound:
				errorResponse.Message = "Group not found"
				w.WriteHeader(http.StatusNotFound)
			case errors.ErrUserNotInGroup:
				errorResponse.Message = "User is not a member of the group"
				w.WriteHeader(http.StatusBadRequest)
			default:
				errorResponse.Message = "Failed to retrieve group members"
				w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse.Message = "Group members retrieved successfully"
	successResponse.Data = members

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)
}