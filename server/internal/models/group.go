package models

import (
	"time"
	"gorm.io/gorm"
)

type Group struct {
	ID 				uint           `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name 			string         `json:"name" gorm:"not null"`
	LocationLat		float64        `json:"location_lat" gorm:"not null"`
	LocationLong	float64        `json:"location_long" gorm:"not null"`
	
	CreatedAt 		time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt 		time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt 		gorm.DeletedAt  `json:"deleted_at,omitempty" gorm:"index"`

	CreatedBy 		uint           `json:"created_by" gorm:"not null"`
	User			User           `json:"-" gorm:"foreignKey:created_by;not null"`
}