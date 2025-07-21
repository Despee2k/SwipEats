package services

import (
	"strings"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/types"
	"gorm.io/gorm"
)

func JoinGroup(groupCode string, userID uint) (*dtos.JoinGroupResponseDto, error) {
	group, err := repositories.GetGroupByCode(strings.ToUpper(groupCode))
	if err != nil {
		return nil, err
	}

	if group == nil {
		return nil, errors.ErrGroupNotFound
	}

	if group.GroupStatus == types.GroupStatusClosed {
		return nil, errors.ErrGroupClosed
	}

	membership, err := repositories.GetGroupMembershipByUserIDAndGroupIDWithDeleted(userID, group.ID)
	if err != nil {
		return nil, err
	}

	// Reuse the membership if it exists and was previously deleted
	if membership != nil {
		membership.DeletedAt = gorm.DeletedAt{} // Restore the membership if it was previously deleted
		err = repositories.UpdateGroupMembership(membership)
		if err != nil {
			return nil, err
		}
		return &dtos.JoinGroupResponseDto{
			Message: "Successfully rejoined the group",
		}, nil
	}

	err = repositories.AddUserToGroup(userID, group.ID, false)
	if err != nil {
		return nil, err
	}

	return &dtos.JoinGroupResponseDto{
		GroupCode: group.GroupCode,
		Message: "Successfully joined the group",
	}, nil
}

func GetUserGroups(userID uint) ([]dtos.GetGroupResponseDto, error) {
	groups, err := repositories.GetGroupsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var groupDtos []dtos.GetGroupResponseDto
	for _, group := range groups {
		count, err := repositories.GetMemberCountByGroupID(group.ID)
		if err != nil {
			return nil, err
		}

		if group.GroupStatus == types.GroupStatusClosed {
			continue // Skip closed groups
		}
		
		groupDtos = append(groupDtos, dtos.GetGroupResponseDto{
			GroupCode:   group.GroupCode,
			Name:        group.Name,
			LocationLat: group.LocationLat,
			LocationLong: group.LocationLong,
			IsOwner:    group.CreatedBy == userID,
			GroupStatus:     group.GroupStatus,
			MemberCount: int(count),
			CreatedAt: group.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return groupDtos, nil
}

func LeaveGroup(userID uint, groupCode string) error {
	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil {
		return err
	}

	if group == nil {
		return errors.ErrGroupNotFound
	}

	if group.GroupStatus == types.GroupStatusClosed {
		return errors.ErrGroupClosed
	}

	membership, err := repositories.GetGroupMembershipByUserIDAndGroupIDWithDeleted(userID, group.ID)

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
		user, err := repositories.GetUserByID(membership.UserID)
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, errors.ErrUserNotFound
		}

		memberDtos = append(memberDtos, dtos.UserMembershipResponseDto{
			UserID:  membership.UserID,
			Name: user.Name,
			Email: user.Email,
			IsOwner: membership.IsOwner,
		})
	}

	return memberDtos, nil
}