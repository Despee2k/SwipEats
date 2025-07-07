package models

import (
	"time"
)

type Restaurant struct {
	ID           uint           `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name         string         `json:"name" gorm:"not null"`
	Address      string         `json:"address" gorm:"not null"`
	LocationLat  float64        `json:"location_lat" gorm:"not null"`
	LocationLong float64        `json:"location_long" gorm:"not null"`
	PhotoURL     string         `json:"photo_url" gorm:"not null"`
	Rating       float64        `json:"rating" gorm:"not null"`
	PriceLevel   int            `json:"price_level" gorm:"not null"`

	CreatedAt	 time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
}