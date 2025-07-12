package services

import (
	"github.com/SwipEats/SwipEats/server/internal/models"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
)

func SaveMostLikedGroupRestaurant(groupRestaurantID uint) error {
	match := &models.Match{
		GroupRestaurantID: groupRestaurantID,
	}

	// Call repository to save the most liked group restaurant
	err := repositories.AddMatch(match)
	if err != nil {
		return err
	}

	return nil
}

func GetGroupMatch(groupID uint) (*models.Match, error) {
	// Call repository to get match by group ID
	match, err := repositories.GetMatchByGroupID(groupID)
	if err != nil {
		return nil, err
	}

	if match == nil {
		return nil, nil // No match found for the group
	}

	return match, nil
}