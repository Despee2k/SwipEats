package models

import (
	"gorm.io/gorm"
	"time"
)

type GroupMembership struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement;"`
	IsOwner   bool           `json:"is_owner" gorm:"not null"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	GroupID   uint           `json:"group_id" gorm:"not null"`
	UserID    uint           `json:"user_id" gorm:"not null"`

	Group	 Group          `json:"-" gorm:"foreignKey:GroupID;not null"`
	User     User           `json:"-" gorm:"foreignKey:UserID;not null"`
}