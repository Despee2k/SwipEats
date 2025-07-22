package services

import (
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
)

func SaveMostLikedGroupRestaurant(groupRestaurantID uint) (uint, error) {
	groupRestaurant, err := repositories.GetGroupRestaurantByID(groupRestaurantID)
	if err != nil {
		return 0, err
	}
	if groupRestaurant == nil {
		return 0, errors.ErrGroupRestaurantNotFound
	}

	match := &models.Match{
		GroupID:      groupRestaurant.GroupID,
		RestaurantID: groupRestaurant.RestaurantID,
	}

	// Call repository to save the most liked group restaurant
	if err = repositories.AddMatch(match); err != nil {
		return 0, err
	}

	return match.ID, nil
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