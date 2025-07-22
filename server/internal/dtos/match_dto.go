package dtos

type MatchResponseDto struct {
	ID                 uint   `json:"id"`
	GroupRestaurantID uint   `json:"group_restaurant_id"`
	GroupRestaurant   GroupRestaurantResponseDto `json:"group_restaurant"`
}