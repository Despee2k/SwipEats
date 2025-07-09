package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/middlewares"
	"github.com/SwipEats/SwipEats/server/internal/services"
	"github.com/go-chi/chi/v5"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var errorResponse dtos.APIErrorResponse
	var successResponse dtos.APISuccessResponse[dtos.UserUpdateResponseDto]

	
	var filePath string
	var err error
	
	w.Header().Set("Content-Type", "application/json")

	user := dtos.UserUpdateRequestDto{
		Name:     r.FormValue("name"),
		Password: r.FormValue("password"),
		ClearImage: r.FormValue("clear_image") == "true",
	}

	userID := r.Context().Value(middlewares.UserIDKey).(uint)

	if !user.ClearImage {
		filePath, err = services.UploadProfilePicture(r, userID)
		if err != nil && err != http.ErrMissingFile {
			errorResponse.Message = "Failed to upload profile picture: " + err.Error()
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if err := services.UpdateUser(user, userID, filePath); err != nil {
		errorResponse.Message = "Failed to update user: " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := dtos.UserUpdateResponseDto{
		Name:           user.Name,
	}

	if user.ClearImage {
		response.ProfilePicture = ""
	} else {
		response.ProfilePicture = filePath
	}

	successResponse.Message = "User updated successfully"
	successResponse.Data = response

	json.NewEncoder(w).Encode(successResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetProfilePictureHandler(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	user, err := services.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	imgPath := filepath.Clean(user.ProfilePicture)
	file, err := os.Open(imgPath)
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "image/jpeg") // or detect dynamically
	http.ServeFile(w, r, imgPath)
}