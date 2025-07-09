package services

import (
	"errors"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *dtos.UserRegisterRequestDto) (*models.User, error) {
	existingUser, err := repositories.GetUserByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	if user.Password != user.ConfirmPassword {
		return nil, bcrypt.ErrMismatchedHashAndPassword // Passwords do not match
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

func LoginUser(user *dtos.UserLoginRequestDto) (string, error) {
	existingUser, err := repositories.GetUserByEmail(user.Email)

	if err != nil {
		return "", err
	}

	if existingUser == nil {
		return "", errors.New("invalid login credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid login credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(existingUser)
	
	if err != nil {
		return "", err
	}

	return token, nil
}