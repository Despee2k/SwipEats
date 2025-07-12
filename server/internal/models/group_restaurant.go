package models

import (
	"gorm.io/gorm"
	"time"
)

type GroupRestaurant struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement;"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	GroupID      uint           `json:"group_id" gorm:"not null"`
	RestaurantID uint           `json:"restaurant_id" gorm:"not null"`

	Group	 Group          `json:"-" gorm:"foreignKey:GroupID;not null"`
	Restaurant Restaurant `json:"-" gorm:"foreignKey:RestaurantID;not null"`
}