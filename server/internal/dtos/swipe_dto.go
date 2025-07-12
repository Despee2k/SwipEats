package dtos

type AddSwipeDto struct {
	IsLiked          bool   `json:"is_liked"`
	GroupRestaurantID uint   `json:"group_restaurant_id"`
}

type SwipeResponseDto struct {
	UserID            uint   `json:"user_id"`
	GroupRestaurantID uint   `json:"group_restaurant_id"`
	IsLiked           bool   `json:"is_liked"`
}