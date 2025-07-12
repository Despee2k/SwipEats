package models

import (
	"time"

	"github.com/SwipEats/SwipEats/server/internal/types"
	"gorm.io/gorm"
)

type Group struct {
	ID 				uint           				`json:"id" gorm:"primaryKey;autoIncrement;"`
	GroupCode		string         				`json:"group_code" gorm:"not null"`
	Name 			string         				`json:"name" gorm:"not null"`
	LocationLat		float64        				`json:"location_lat" gorm:"not null"`
	LocationLong	float64        				`json:"location_long" gorm:"not null"`
	GroupStatus		types.GroupStatusEnum		`json:"group_status" gorm:"not null;default:'waiting'"` // waiting, active, closed

	CreatedAt 		time.Time      				`json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt 		time.Time      				`json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt 		gorm.DeletedAt  			`json:"deleted_at,omitempty" gorm:"index"`

	CreatedBy 		uint           				`json:"created_by" gorm:"not null"`
	User			User           				`json:"-" gorm:"foreignKey:created_by;not null"`
}