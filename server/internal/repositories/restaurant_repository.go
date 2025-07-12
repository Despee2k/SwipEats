package repositories

import (
	"github.com/SwipEats/SwipEats/server/internal/db"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"gorm.io/gorm"
)

func AddRestaurant(restaurant *models.Restaurant) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Create(restaurant)
	if result.Error != nil {
		return result.Error // Error creating restaurant
	}
	return nil // Restaurant created successfully
}

func GetRestaurantByID(restaurantID uint) (*models.Restaurant, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var restaurant models.Restaurant
	result := db.Conn.Where("id = ? AND deleted_at IS NULL", restaurantID).First(&restaurant)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Restaurant not found
		}
		return nil, result.Error // Other error
	}
	return &restaurant, nil
}

func GetNearbyRestaurants(lat float64, long float64, radius float64) ([]models.Restaurant, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var restaurants []models.Restaurant
	result := db.Conn.Where("ST_Distance_Sphere(location, ST_MakePoint(?, ?)) <= ?", long, lat, radius).Find(&restaurants)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No restaurants found
		}
		return nil, result.Error // Other error
	}
	return restaurants, nil
}

func UpdateRestaurant(restaurant *models.Restaurant) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Save(restaurant)
	if result.Error != nil {
		return result.Error // Error updating restaurant
	}
	return nil // Restaurant updated successfully
}

func DeleteRestaurant(restaurant *models.Restaurant) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Delete(restaurant)
	if result.Error != nil {
		return result.Error // Error deleting restaurant
	}

	return nil // Restaurant deleted successfully
}