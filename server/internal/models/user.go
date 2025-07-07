package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID	   			uint   			`json:"id" gorm:"primaryKey;autoIncrement;"`
	Name 			string 			`json:"name" gorm:"not null"`
	Email 			string 			`json:"email" gorm:"unique;not null"`
	Password 		string 			`json:"password" gorm:"not null"`
	ProfilePicture 	string 			`json:"profile_picture"`
	
	CreatedAt 		time.Time 		`json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt 		time.Time 		`json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt 		gorm.DeletedAt 	`json:"deleted_at,omitempty" gorm:"index"`
}