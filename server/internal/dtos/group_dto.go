package dtos

import "github.com/SwipEats/SwipEats/server/internal/types"

type CreateGroupRequestDto struct {
	Name			string  `json:"name" validate:"required"`
	LocationLat		float64 `json:"location_lat" validate:"required"`
	LocationLong	float64 `json:"location_long" validate:"required"`
}

type CreateGroupResponseDto struct {
	GroupCode		string  `json:"group_code"`
}

type GetGroupResponseDto struct {
	GroupCode		string  `json:"group_code"`
	Name			string  `json:"name"`
	LocationLat		float64 `json:"location_lat"`
	LocationLong	float64 `json:"location_long"`
	IsOwner			bool    `json:"is_owner"`
	GroupStatus		types.GroupStatusEnum `json:"group_status"`
	MemberCount		int     `json:"member_count"`
	CreatedAt		string  `json:"created_at"`
}

type CheckIfGroupExistsDto struct {
	Exists bool `json:"exists"`
}