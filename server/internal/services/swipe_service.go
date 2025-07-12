package services

import (
	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
)

func AddSwipe(swipe dtos.AddSwipeDto, userID uint) error {
	existingSwipe, err := repositories.GetSwipeByUserAndGroupRestaurantID(userID, swipe.GroupRestaurantID)
	if err != nil {
		return err
	}
	if existingSwipe != nil {
		return errors.ErrSwipeAlreadyExists // Swipe already exists for this user and group restaurant
	}

	// Convert DTO to model
	swipeModel := &models.Swipe{
		UserID:            userID,
		GroupRestaurantID: swipe.GroupRestaurantID,
		IsLiked:           swipe.IsLiked,
	}

	// Call repository to create the swipe
	err = repositories.CreateSwipe(swipeModel)
	if err != nil {
		return err
	}

	return nil
}

func GetSwipesByGroupRestaurant(groupRestaurantID uint) ([]dtos.SwipeResponseDto, error) {
	// Call repository to get swipes by group restaurant ID
	swipes, err := repositories.GetSwipesByGroupRestaurant(groupRestaurantID)
	if err != nil {
		return nil, err
	}

	// Convert model to DTO
	var swipeDtos []dtos.SwipeResponseDto
	for _, swipe := range swipes {
		swipeDtos = append(swipeDtos, dtos.SwipeResponseDto{
			UserID:            swipe.UserID,
			GroupRestaurantID: swipe.GroupRestaurantID,
			IsLiked:           swipe.IsLiked,
		})
	}

	return swipeDtos, nil
}

func GetLikeCountByGroupRestaurant(groupRestaurantID uint) (int, error) {
	// Call repository to get swipe count by group restaurant ID
	count, err := repositories.GetLikeCountByGroupRestaurant(groupRestaurantID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetUserSwipeCount(userID uint, groupID uint) (int, error) {
	// Call repository to get swipe count by user ID
	count, err := repositories.GetSwipeCountByUserAndGroup(userID, groupID)
	if err != nil {
		return 0, err
	}

	return count, nil
}