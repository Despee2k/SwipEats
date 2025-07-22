package services

import (
	"github.com/SwipEats/SwipEats/server/internal/constants"
	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/utils"
)

func GenerateGroupRestaurants(groupCode string, radius int, numberOfRestaurants int) ([]dtos.GroupRestaurantResponseDto, error) {
	numberOfRestaurants = min(constants.MAX_NUM_OF_RESTAURANTS, numberOfRestaurants) // Ensure it does not exceed the maximum limit

	// Fetch the group by code
	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, errors.ErrGroupNotFound
	}

	restaurants, err := repositories.GetNearbyRestaurants(group.LocationLat, group.LocationLong, radius)
	if err != nil {
		return nil, err
	}

	if len(restaurants) < numberOfRestaurants {
		// Logic to fetch and save more restaurants if needed
		newRestaurants, err := FetchRestaurantsNearby(group.LocationLat, group.LocationLong, radius)
		if err != nil {
			return nil, err
		}
		if len(newRestaurants) > 0 {
			restaurants = append(restaurants, newRestaurants...)
		}
	}

	var groupRestaurants []models.GroupRestaurant
	for _, restaurant := range restaurants {
		groupRestaurant := &models.GroupRestaurant{
			GroupID:      group.ID,
			RestaurantID: restaurant.ID,
		}
		err = repositories.AddGroupRestaurant(groupRestaurant)
		if err != nil {
			return nil, err // Error adding group restaurant
		}

		groupRestaurant.Restaurant = restaurant // Populate the restaurant details
		groupRestaurants = append(groupRestaurants, *groupRestaurant)
	}

	var responseDtos []dtos.GroupRestaurantResponseDto
	for _, groupRestaurant := range groupRestaurants {
		responseDto := dtos.GroupRestaurantResponseDto{
			ID:        groupRestaurant.ID,
			GroupID:  groupRestaurant.GroupID,
			Restaurant: groupRestaurant.Restaurant,
			DistanceInKM: utils.DistanceInKM(
				group.LocationLat, 
				group.LocationLong, 
				groupRestaurant.Restaurant.LocationLat, 
				groupRestaurant.Restaurant.LocationLong,
			),
		}
		responseDtos = append(responseDtos, responseDto)
	}

	return responseDtos, nil
}

func GetGroupRestaurantsByGroupCode(groupCode string) ([]dtos.GroupRestaurantResponseDto, error) {
	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, errors.ErrGroupNotFound
	}

	groupRestaurants, err := repositories.GetGroupRestaurantsByGroupID(group.ID)
	if err != nil {
		return nil, err
	}

	var responseDtos []dtos.GroupRestaurantResponseDto
	for _, groupRestaurant := range groupRestaurants {
		responseDto := dtos.GroupRestaurantResponseDto{
			ID:        groupRestaurant.ID,
			GroupID:  group.ID,
			Restaurant: groupRestaurant.Restaurant,
			DistanceInKM: utils.DistanceInKM(
				group.LocationLat, 
				group.LocationLong, 
				groupRestaurant.Restaurant.LocationLat, 
				groupRestaurant.Restaurant.LocationLong,
			),
		}
		responseDtos = append(responseDtos, responseDto)
	}

	return responseDtos, nil
}

func GetGroupRestaurantCountByGroupCode(groupCode string) (int, error) {
	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		return 0, err
	}
	if group == nil {
		return 0, errors.ErrGroupNotFound
	}

	count, err := repositories.GetGroupRestaurantCountByGroupID(group.ID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetMostLikedGroupRestaurant(groupCode string) (*dtos.GroupRestaurantResponseDto, error) {
	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, errors.ErrGroupNotFound
	}

	// Call repository to get the most liked group restaurant
	groupRestaurant, err := repositories.GetMostLikedGroupRestaurant(group.ID)
	if err != nil {
		return nil, err
	}
	if groupRestaurant == nil {
		return nil, nil 
	}

	// Convert model to DTO
	responseDto := dtos.GroupRestaurantResponseDto{
		ID:        groupRestaurant.ID,
		GroupID:  	group.ID,
		RestaurantID: groupRestaurant.RestaurantID,
		Restaurant: groupRestaurant.Restaurant,
		DistanceInKM: utils.DistanceInKM(
			group.LocationLat, 
			group.LocationLong, 
			groupRestaurant.Restaurant.LocationLat, 
			groupRestaurant.Restaurant.LocationLong,
		),
	}

	return &responseDto, nil
}