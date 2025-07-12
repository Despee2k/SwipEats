package models

import (
	"time"
)

type Match struct {
	ID        		uint           `json:"id" gorm:"primaryKey;autoIncrement;"`

	CreatedAt 		time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`

	GroupRestaurantID uint `json:"group_restaurant_id" gorm:"not null"`

	GroupRestaurant GroupRestaurant `json:"-" gorm:"foreignKey:GroupRestaurantID;references:ID"`
}