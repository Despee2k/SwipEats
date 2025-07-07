package models

import (
	"time"
)

type Match struct {
	ID        		uint           `json:"id" gorm:"primaryKey;autoIncrement;"`

	CreatedAt 		time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`

	GroupID   		uint           `json:"group_id" gorm:"not null"`
	RestaurantID 	uint           `json:"restaurant_id" gorm:"not null"`

	Group	   		Group          `json:"-" gorm:"foreignKey:GroupID;not null"`
	Restaurant  	Restaurant     `json:"-" gorm:"foreignKey:RestaurantID;not null"`
}