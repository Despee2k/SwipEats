package dtos

import "github.com/SwipEats/SwipEats/server/internal/models"

type GroupRestaurantResponseDto struct {
	ID            uint   `json:"id"`
	GroupID       uint   `json:"group_id"`
	Restaurant    models.Restaurant `json:"restaurant"`
}