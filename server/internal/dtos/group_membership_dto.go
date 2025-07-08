package dtos

type JoinGroupResponseDto struct {
	Message string `json:"message"`
}

type UserMembershipResponseDto struct {
	UserID    uint   `json:"user_id"`
	IsOwner   bool   `json:"is_owner"`
}