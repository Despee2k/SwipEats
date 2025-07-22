package repositories

import (
	"github.com/SwipEats/SwipEats/server/internal/db"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"gorm.io/gorm"
)

func CreateSwipe(swipe *models.Swipe) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Create(swipe)
	if result.Error != nil {
		return result.Error // Error creating swipe
	}
	return nil // Swipe created successfully
}

func GetSwipeByID(swipeID uint) (*models.Swipe, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var swipe models.Swipe
	result := db.Conn.Where("id = ? AND deleted_at IS NULL", swipeID).First(&swipe)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Swipe not found
		}
		return nil, result.Error // Other error
	}
	return &swipe, nil
}

func GetUserSwipes(userID uint) ([]models.Swipe, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var swipes []models.Swipe
	result := db.Conn.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&swipes)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No swipes found for user
		}
		return nil, result.Error // Other error
	}
	return swipes, nil
}

func GetGroupSwipes(groupID uint) ([]models.Swipe, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var swipes []models.Swipe
	result := db.Conn.Where("group_id = ? AND deleted_at IS NULL", groupID).Find(&swipes)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No swipes found for group
		}
		return nil, result.Error // Other error
	}
	return swipes, nil
}

func GetUnfinishedGroupSwipes(groupID uint) ([]models.Swipe, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var swipes []models.Swipe
	result := db.Conn.Where("group_id = ? AND is_liked IS NULL AND deleted_at IS NULL", groupID).Find(&swipes)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No unfinished swipes found for group
		}
		return nil, result.Error // Other error
	}
	return swipes, nil
}

func GetSwipesByGroupRestaurant(groupRestaurantID uint) ([]models.Swipe, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var swipes []models.Swipe
	result := db.Conn.Where("group_restaurant_id = ? AND deleted_at IS NULL", groupRestaurantID).Find(&swipes)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No swipes found for group restaurant
		}
		return nil, result.Error // Other error
	}
	return swipes, nil
}

func GetLikeCountByGroupRestaurant(groupRestaurantID uint) (int, error) {
	if db.Conn == nil {
		return 0, gorm.ErrInvalidDB // Database connection is not established
	}

	var count int64
	result := db.Conn.Model(&models.Swipe{}).Where("group_restaurant_id = ? AND is_liked = ? AND deleted_at IS NULL", groupRestaurantID, true).Count(&count)

	if result.Error != nil {
		return 0, result.Error // Error counting swipes
	}
	return int(count), nil // Return the count of swipes
}
func GetSwipeCountByUserAndGroup(userID uint, groupID uint) (int, error) {
	if db.Conn == nil {
		return 0, gorm.ErrInvalidDB // Database connection is not established
	}

	var count int64
	result := db.Conn.Model(&models.Swipe{}).Joins("JOIN group_restaurants ON group_restaurants.id = swipes.group_restaurant_id").
		Where("swipes.user_id = ? AND group_restaurants.group_id = ? AND swipes.deleted_at IS NULL", userID, groupID).
		Count(&count)

	if result.Error != nil {
		return 0, result.Error // Error counting swipes
	}
	return int(count), nil // Return the count of swipes
}

func GetSwipeByUserAndGroupRestaurantID(userID uint, groupRestaurantID uint) (*models.Swipe, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var swipe models.Swipe
	result := db.Conn.Where("user_id = ? AND group_restaurant_id = ? AND deleted_at IS NULL", userID, groupRestaurantID).First(&swipe)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Swipe not found
		}
		return nil, result.Error // Other error
	}
	return &swipe, nil
}

func UpdateSwipe(swipe *models.Swipe) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Save(swipe)
	if result.Error != nil {
		return result.Error // Error updating swipe
	}
	return nil // Swipe updated successfully
}

func DeleteSwipe(swipe *models.Swipe) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Delete(swipe)
	if result.Error != nil {
		return result.Error // Error deleting swipe
	}
	return nil // Swipe deleted successfully
}