package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/services"
	"github.com/SwipEats/SwipEats/server/internal/utils"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user dtos.UserRegisterRequestDto
	var errorResponse dtos.APIErrorResponse
	var successResponse dtos.APISuccessResponse[dtos.UserResponseDto]

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errorResponse.Message = "Invalid request body"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		message, details := utils.GetValidationErrorDetails(err)

		errorResponse.Message = message
		errorResponse.Details = details

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	new_user, err := services.RegisterUser(&user)
	if err != nil {
		switch err {
			case errors.ErrEmailAlreadyInUse:
				errorResponse.Message = "Email already in use"
				errorResponse.Details = map[string]string{"email": "Already in use."}
				w.WriteHeader(http.StatusConflict)
			case errors.ErrPasswordsDoNotMatch:
				errorResponse.Message = "Passwords do not match"
				errorResponse.Details = map[string]string{
					"password": "Do not match.", 
					"confirm_password": "Do not match.",
				}
				w.WriteHeader(http.StatusBadRequest)
			default:
				errorResponse.Message = "Failed to register user"
				w.WriteHeader(http.StatusInternalServerError)
		}
		
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse.Message = "User registered successfully"
	successResponse.Data = dtos.UserResponseDto{
		ID:    new_user.ID,
		Email: new_user.Email,
	}

	json.NewEncoder(w).Encode(successResponse)
	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user dtos.UserLoginRequestDto
	var errorResponse dtos.APIErrorResponse
	var successResponse dtos.APISuccessResponse[dtos.UserLoginResponseDto]

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errorResponse.Message = "Invalid request body"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		message, details := utils.GetValidationErrorDetails(err)

		errorResponse.Message = message
		errorResponse.Details = details

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	token, userID, err := services.LoginUser(&user)
	if err != nil {

		switch err {
			case errors.ErrInvalidCredentials:
				errorResponse.Message = "Invalid login credentials"
				errorResponse.Details = map[string]string{
					"user_id":  "invalid",
					"email":    "invalid",
					"password": "invalid",
				}
			default:
				errorResponse.Message = "Failed to login user"
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := dtos.UserLoginResponseDto{
		UserID: userID,
		Email:  user.Email,
		Token:  token,
	}

	successResponse.Message = "Login successful"
	successResponse.Data = response

	json.NewEncoder(w).Encode(successResponse)
	w.WriteHeader(http.StatusOK)
}