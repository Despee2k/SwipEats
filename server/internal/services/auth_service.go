package services

import (
	"errors"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *dtos.UserRegisterRequestDto) error {
	existingUser, err := repositories.GetUserByEmail(user.Email)

	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("user already exists")
	}

	if user.Password != user.ConfirmPassword {
		return bcrypt.ErrMismatchedHashAndPassword // Passwords do not match
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	newUser := &models.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	return repositories.CreateUser(newUser)
}

func LoginUser(user *dtos.UserLoginRequestDto) (string, error) {
	existingUser, err := repositories.GetUserByEmail(user.Email)

	if err != nil {
		return "", err
	}

	if existingUser == nil {
		return "", bcrypt.ErrMismatchedHashAndPassword // User not found
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(existingUser)
	
	if err != nil {
		return "", err
	}

	return token, nil
}