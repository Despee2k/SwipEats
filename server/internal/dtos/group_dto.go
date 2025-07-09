package dtos

type CreateGroupRequestDto struct {
	Name			string  `json:"name" validate:"required"`
	LocationLat		float64 `json:"location_lat" validate:"required"`
	LocationLong	float64 `json:"location_long" validate:"required"`
}

type CreateGroupResponseDto struct {
	GroupCode		string  `json:"group_code"`
}