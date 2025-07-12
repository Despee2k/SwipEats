package repositories

import (
	"github.com/SwipEats/SwipEats/server/internal/db"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"gorm.io/gorm"
)

func AddGroupRestaurant(groupRestaurant *models.GroupRestaurant) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Create(groupRestaurant)
	if result.Error != nil {
		return result.Error // Error creating group restaurant
	}

	return nil // Group restaurant created successfully
}

func GetGroupRestaurantsByGroupID(groupID uint) ([]models.GroupRestaurant, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var groupRestaurants []models.GroupRestaurant
	if err := db.Conn.Where("group_id = ? AND deleted_at IS NULL", groupID).Find(&groupRestaurants).Error; err != nil {
		return nil, err
	}

	return groupRestaurants, nil
}

func DeleteGroupRestaurant(groupRestaurant *models.GroupRestaurant) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Delete(groupRestaurant)
	if result.Error != nil {
		return result.Error // Error deleting group restaurant
	}

	return nil // Group restaurant deleted successfully
}