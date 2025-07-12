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

func GetGroupRestaurantByGroupAndRestaurantID(groupID uint, restaurantID uint) (*models.GroupRestaurant, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var groupRestaurant models.GroupRestaurant
	result := db.Conn.Where("group_id = ? AND restaurant_id = ? AND deleted_at IS NULL", groupID, restaurantID).First(&groupRestaurant)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Group restaurant not found
		}
		return nil, result.Error // Other error
	}
	return &groupRestaurant, nil
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

func GetGroupRestaurantCountByGroupID(groupID uint) (int, error) {
	if db.Conn == nil {
		return 0, gorm.ErrInvalidDB // Database connection is not established
	}

	var count int64
	if err := db.Conn.Model(&models.GroupRestaurant{}).Where("group_id = ? AND deleted_at IS NULL", groupID).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func GetMostLikedGroupRestaurant(groupID uint) (*models.GroupRestaurant, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var groupRestaurant models.GroupRestaurant
	result := db.Conn.Table("swipes").Select("group_restaurant_id, COUNT(*) as like_count").
		Where("group_restaurants.group_id = ? AND is_liked = ? AND swipes.deleted_at IS NULL", groupID, true).
		Group("group_restaurant_id").
		Order("like_count DESC").
		Limit(1).
		Joins("JOIN group_restaurants ON swipes.group_restaurant_id = group_restaurants.id").
		First(&groupRestaurant)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No likes found for group restaurant
		}
		return nil, result.Error // Other error
	}
	return &groupRestaurant, nil
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