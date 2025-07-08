package services

import (
	"errors"
	"strings"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
)

func JoinGroup(groupCode string, userID uint) (*dtos.JoinGroupResponseDto, error) {
	group, err := repositories.GetGroupByCode(strings.ToUpper(groupCode))
	if err != nil {
		return nil, err
	}

	if group == nil {
		return nil, errors.New("group not found")
	}

	membership, err := repositories.GetGroupMembershipByUserIDAndGroupID(userID, group.ID)
	if err != nil {
		return nil, err
	}

	if membership != nil {
		return nil, errors.New("user is already a member of the group")
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
		return errors.New("group not found")
	}

	membership, err := repositories.GetGroupMembershipByUserIDAndGroupID(userID, group.ID)

	if err != nil {
		return err
	}

	if membership == nil {
		return errors.New("user is not a member of the group")
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
		return nil, errors.New("group not found")
	}

	confirmMembership, err := repositories.GetGroupMembershipByUserIDAndGroupID(userID, group.ID)
	if err != nil {
		return nil, err
	}

	if confirmMembership == nil {
		return nil, errors.New("user is not a member of the group")
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