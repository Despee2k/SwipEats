package services

import (
	"errors"
	"strings"

	"github.com/SwipEats/SwipEats/server/internal/constants"
	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
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

	return "", errors.New("failed to generate a unique group code after multiple attempts")
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