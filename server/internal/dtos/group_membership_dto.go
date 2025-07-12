package dtos

type JoinGroupResponseDto struct {
	Message string `json:"message"`
}

type UserMembershipResponseDto struct {
	UserID   uint   `json:"user_id"`
	Name string `json:"username"`
	IsOwner  bool   `json:"is_owner"`
}