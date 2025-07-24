package services

import (
	"github.com/SwipEats/SwipEats/server/internal/constants"
	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/utils"
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

func GetUserRecentMatches(userID uint) ([]dtos.GroupRestaurantResponseDto, error) {
	// Call repository to get recent matches for the user
	groups, err := repositories.GetGroupsByUserID(userID)
	if err != nil {
		return nil, err
	}

	count := 0
	var matches []*models.Match
	for _, group := range groups {
		match, err := GetGroupMatch(group.ID)
		if err != nil {
			return nil, err
		}
		if match != nil {
			matches = append(matches, match)
			count++
		}
		if count >= constants.MAX_NUM_OF_RECENT_MATCHES {
			break
		}
	}

	var response []dtos.GroupRestaurantResponseDto
	for _, match := range matches {
		groupRestaurant, err := repositories.GetGroupRestaurantByID(match.RestaurantID)
		if err != nil {
			return nil, err
		}
		if groupRestaurant != nil {
			response = append(response, dtos.GroupRestaurantResponseDto{
				ID:        match.ID,
				GroupID:  match.GroupID,
				Restaurant: match.Restaurant,
				DistanceInKM: utils.DistanceInKM(
					match.Group.LocationLat,
					match.Group.LocationLong,
					match.Restaurant.LocationLat,
					match.Restaurant.LocationLong,
				),
			})
		}
	}

	return response, nil
}