package services

import (
	"strings"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/types"
)

func JoinGroup(groupCode string, userID uint) (*dtos.JoinGroupResponseDto, error) {
	group, err := repositories.GetGroupByCode(strings.ToUpper(groupCode))
	if err != nil {
		return nil, err
	}

	if group == nil {
		return nil, errors.ErrGroupNotFound
	}

	if group.GroupStatus != types.GroupStatusWaiting {
		return nil, errors.ErrGroupClosed
	}

	membership, err := repositories.GetGroupMembershipByUserIDAndGroupID(userID, group.ID)
	if err != nil {
		return nil, err
	}

	if membership != nil {
		return nil, errors.ErrUserAlreadyInGroup
	}

	err = repositories.AddUserToGroup(userID, group.ID, false)
	if err != nil {
		return nil, err
	}

	return &dtos.JoinGroupResponseDto{
		Message: "Successfully joined the group",
	}, nil
}

func LeaveGroup(userID uint, groupCode string) error {
	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		return err
	}

	if group == nil {
		return errors.ErrGroupNotFound
	}

	membership, err := repositories.GetGroupMembershipByUserIDAndGroupID(userID, group.ID)

	if err != nil {
		return err
	}

	if membership == nil {
		return errors.ErrUserNotInGroup
	}

	err = repositories.RemoveUserFromGroup(membership)

	if err != nil {
		return err
	}
	return nil
}

func GetGroupMembers(groupCode string, userID uint) ([]dtos.UserMembershipResponseDto, error) {
	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		return nil, err
	}

	if group == nil {
		return nil, errors.ErrGroupNotFound
	}

	confirmMembership, err := repositories.GetGroupMembershipByUserIDAndGroupID(userID, group.ID)
	if err != nil {
		return nil, err
	}

	if confirmMembership == nil {
		return nil, errors.ErrUserNotInGroup
	}

	memberships, err := repositories.GetGroupMembershipsByGroupID(group.ID)
	if err != nil {
		return nil, err
	}

	var memberDtos []dtos.UserMembershipResponseDto
	for _, membership := range memberships {
		memberDtos = append(memberDtos, dtos.UserMembershipResponseDto{
			UserID:  membership.UserID,
			IsOwner: membership.IsOwner,
		})
	}

	return memberDtos, nil
}