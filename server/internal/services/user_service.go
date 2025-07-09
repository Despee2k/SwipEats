package services

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/SwipEats/SwipEats/server/internal/constants"
	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

func UploadProfilePicture(r *http.Request, userID uint) (string, error) {
	err := r.ParseMultipartForm(constants.MAX_FILE_SIZE) // Limit your max input length!
	if err != nil {
		return "", err
	}

	file, _, err := r.FormFile("profile_picture")
	if err != nil {
		return "", err // Return error if file upload fails
	}
	defer file.Close()

	os.MkdirAll(constants.UPLOAD_DIR, os.ModePerm)

	// Save the uploaded file to the server
	dstPath := filepath.Join(constants.UPLOAD_DIR, "profile_picture_"+fmt.Sprintf("%d", userID)+"_" +fmt.Sprintf("%d", time.Now().Unix())+".jpg")
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err // Return error if file creation fails
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", errors.New("could not save file: " + err.Error())
	}

	return dstPath, nil
}

func UpdateUser(user dtos.UserUpdateRequestDto, userID uint, filePath string) error {
	existingUser, err := repositories.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Update the existing user with the new values
	existingUser.Name = user.Name

	if filePath != "" || user.ClearImage {
		if existingUser.ProfilePicture != "" {
			absPath, _ := filepath.Abs(existingUser.ProfilePicture)
			err = os.Remove(absPath) // Remove old profile picture if it exists
			if err != nil {
				return errors.New("could not remove old profile picture: " + err.Error())
			}
		}

		if user.ClearImage {
			existingUser.ProfilePicture = "" // Clear profile picture if requested
		} else {
			existingUser.ProfilePicture = filePath // Update with new profile picture path
		}
	}

	if user.Password != "" {
		bcryptHashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err // Return error if password hashing fails
		}

		existingUser.Password = string(bcryptHashedPassword)
	}

	// Save the updated user back to the database
	err = repositories.UpdateUser(existingUser)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(userID uint) (*dtos.UserResponseDto, error) {
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	response := &dtos.UserResponseDto{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
	}

	return response, nil
}

func GetUserByEmail(email string) (*dtos.UserResponseDto, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	response := &dtos.UserResponseDto{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
	}

	return response, nil
}