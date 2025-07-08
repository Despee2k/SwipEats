package repositories

import (
	"github.com/SwipEats/SwipEats/server/internal/db"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"gorm.io/gorm"
)

func GetUserByEmail(email string) (*models.User, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var user models.User
	result := db.Conn.Where("email = ? AND deleted_at IS NULL", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // User not found
		}
		return nil, result.Error // Other error
	}
	return &user, nil
}

func GetUserByID(userID uint) (*models.User, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var user models.User
	result := db.Conn.Where("id = ? AND deleted_at IS NULL", userID).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // User not found
		}
		return nil, result.Error // Other error
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Create(user)
	if result.Error != nil {
		return result.Error // Error creating user
	}
	return nil // User created successfully
}