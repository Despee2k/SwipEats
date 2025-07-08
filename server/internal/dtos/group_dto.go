package dtos

type CreateGroupRequestDto struct {
	Name			string  `json:"name" binding:"required"`
	LocationLat		float64 `json:"location_lat" binding:"required"`
	LocationLong	float64 `json:"location_long" binding:"required"`
}

type CreateGroupResponseDto struct {
	GroupCode		string  `json:"group_code"`
}