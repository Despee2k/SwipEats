package dtos

type JoinGroupResponseDto struct {
	GroupCode string `json:"group_code"`
	Message   string `json:"message"`
}

type UserMembershipResponseDto struct {
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsOwner  bool   `json:"is_owner"`
}