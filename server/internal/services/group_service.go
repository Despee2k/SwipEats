package services

import (
	"strings"

	"github.com/SwipEats/SwipEats/server/internal/constants"
	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/types"
	"github.com/SwipEats/SwipEats/server/internal/utils"
)

func GenerateGroupCode() (string, error) {
	var existingGroup *models.Group

	for i := 0; i < 10; i++ {
		groupCode, err := utils.GenerateGroupCode(constants.GROUP_CODE_LENGTH)

		if err != nil {
			return "", err
		}

		existingGroup, err = repositories.GetGroupByCode(groupCode)

		if err != nil {
			return "", err
		}

		if existingGroup == nil {
			return groupCode, nil
		}
	}

	return "", errors.ErrUnableToGenerateGroupCode
}

func CreateGroup(groupDto dtos.CreateGroupRequestDto, userID uint) (string, error) {
	groupCode, err := GenerateGroupCode()

	if err != nil {
		return "", err
	}

	group := &models.Group{
		Name:         groupDto.Name,
		LocationLat:  groupDto.LocationLat,
		LocationLong: groupDto.LocationLong,
		GroupCode:   strings.ToUpper(groupCode),
	}

	err = repositories.CreateGroup(group, userID)

	if err != nil {
		return "", err
	}

	// Automatically join the group after creation
	err = repositories.AddUserToGroup(
		userID,
		group.ID,
		true,
	)

	if err != nil {
		return "", err
	}

	return group.GroupCode, nil
}

func StartGroupSession(groupCode string, userID uint) error {
	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		return err
	}
	if group == nil {
		return errors.ErrGroupNotFound
	}

	member, err := repositories.GetGroupMembershipByUserIDAndGroupID(userID, group.ID)
	if err != nil {
		return err
	}
	if member == nil {
		return errors.ErrUserNotFound
	}
	if !member.IsOwner {
		return errors.ErrUserNotAuthorized
	}

	group.GroupStatus = types.GroupStatusActive

	if err := repositories.UpdateGroup(group); err != nil {
		return err
	}

	return nil
}

func EndGroupSession(groupCode string, userID uint) error {
	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		return err
	}
	if group == nil {
		return errors.ErrGroupNotFound
	}

	member, err := repositories.GetGroupMembershipByUserIDAndGroupID(userID, group.ID)
	if err != nil {
		return err
	}

	if member == nil {
		return errors.ErrUserNotFound
	}

	if !member.IsOwner {
		return errors.ErrUserNotAuthorized
	}

	group.GroupStatus = types.GroupStatusClosed

	if err := repositories.UpdateGroup(group); err != nil {
		return err
	}

	return nil
}

func CheckIfGroupIsDone(groupCode string) (bool, error) {
	restaurantCount, err := GetGroupRestaurantCountByGroupCode(groupCode)
	if err != nil {
		return false, err
	}
	if restaurantCount == 0 {
		return false, errors.ErrNoRestaurantsFound
	}

	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		return false, err
	}
	if group == nil {
		return false, errors.ErrGroupNotFound
	}

	// Add logic to check if all group members are done swiping
	members, err := repositories.GetGroupMembershipsByGroupID(group.ID)
	if err != nil {
		return false, err
	}

	for _, member := range members {
		count, err := repositories.GetSwipeCountByUserAndGroup(member.UserID, group.ID)
		if err != nil {
			return false, err
		}
		if count < restaurantCount {
			return false, nil // Not all members are done swiping
		}
	}

	return true, nil
}

func CheckIfGroupExists(groupCode string) (bool, error) {
	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		return false, err
	}
	return group != nil, nil
}