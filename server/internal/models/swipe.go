package models

import (
	"time"
)

type Swipe struct {
	ID        	 uint           `json:"id" gorm:"primaryKey;autoIncrement;"`
	IsLiked   	 bool           `json:"is_liked" gorm:"not null"`

	CreatedAt 	 time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`

	UserID       uint           `json:"user_id" gorm:"not null"`
	RestaurantID uint           `json:"restaurant_id" gorm:"not null"`
	GroupID      uint           `json:"group_id" gorm:"not null"`

	User         User           `json:"-" gorm:"foreignKey:UserID;not null"`
	Restaurant   Restaurant     `json:"-" gorm:"foreignKey:RestaurantID;not null"`
	Group        Group          `json:"-" gorm:"foreignKey:GroupID;not null"`
}
