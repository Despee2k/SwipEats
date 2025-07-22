package models

import (
	"time"
)

type Match struct {
	ID        		uint           `json:"id" gorm:"primaryKey;autoIncrement;"`

	CreatedAt 		time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`

	GroupID 		uint           `json:"group_id" gorm:"not null;index;"`
	RestaurantID  uint           `json:"restaurant_id" gorm:"not null;index;"`

	Group      Group		  `json:"group" gorm:"foreignKey:GroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Restaurant  Restaurant	  `json:"restaurant" gorm:"foreignKey:RestaurantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}