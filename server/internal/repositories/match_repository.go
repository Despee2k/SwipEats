package repositories

import (
	"github.com/SwipEats/SwipEats/server/internal/db"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"gorm.io/gorm"
)

func AddMatch(match *models.Match) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Create(match)
	if result.Error != nil {
		return result.Error // Error creating match
	}
	return nil // Match created successfully
}

func GetMatchByID(matchID uint) (*models.Match, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var match models.Match
	result := db.Conn.Where("id = ?", matchID).Preload("Group").Preload("Restaurant").First(&match)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Match not found
		}
		return nil, result.Error // Other error
	}
	return &match, nil
}

func GetMatchByGroupID(groupID uint) (*models.Match, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var match models.Match
	result := db.Conn.Where("group_id = ? AND deleted_at IS NULL", groupID).First(&match)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Match not found
		}
		return nil, result.Error // Other error
	}
	return &match, nil
}