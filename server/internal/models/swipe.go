package models

import (
	"time"

	"gorm.io/gorm"
)

type Swipe struct {
	ID        	 uint           `json:"id" gorm:"primaryKey;autoIncrement;"`
	IsLiked   	 bool           `json:"is_liked"`

	CreatedAt 	 time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt 	 time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt 	 gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	UserID       uint           `json:"user_id" gorm:"not null"`
	GroupRestaurantID uint `json:"group_restaurant_id" gorm:"not null"`

	User         User           `json:"-" gorm:"foreignKey:UserID;not null"`
	GroupRestaurant GroupRestaurant `json:"-" gorm:"foreignKey:GroupRestaurantID;not null"`
}
