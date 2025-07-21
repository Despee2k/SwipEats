package services

import (
	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/utils"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *dtos.UserRegisterRequestDto) (*models.User, error) {
	existingUser, err := repositories.GetUserByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.ErrEmailAlreadyInUse
	}

	if user.Password != user.ConfirmPassword {
		return nil, errors.ErrPasswordsDoNotMatch // Passwords do not match
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	return newUser, repositories.CreateUser(newUser)
}

func LoginUser(user *dtos.UserLoginRequestDto) (string, uint, error) {
	existingUser, err := repositories.GetUserByEmail(user.Email)

	if err != nil {
		return "", 0, err
	}

	if existingUser == nil {
		return "", 0, errors.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", 0, errors.ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(existingUser)
	
	if err != nil {
		return "", 0, err
	}

	return token, existingUser.ID, nil
}