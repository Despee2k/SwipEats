package repositories

import (
	"github.com/SwipEats/SwipEats/server/internal/db"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"gorm.io/gorm"
)

func GetGroupByCode(groupCode string) (*models.Group, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var group models.Group
	result := db.Conn.Where("group_code = ? AND deleted_at IS NULL", groupCode).First(&group)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil 
		}
		return nil, result.Error 
	}
	return &group, nil
}

func GetGroupsByUserID(userID uint) ([]models.Group, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var groups []models.Group
	result := db.Conn.Joins("JOIN group_memberships ON group_memberships.group_id = groups.id").
		Where("group_memberships.user_id = ? AND group_memberships.deleted_at IS NULL AND groups.deleted_at IS NULL", userID).
		Find(&groups).Order("groups.created_at DESC")

	if result.Error != nil {
		return nil, result.Error // Error fetching groups
	}
	return groups, nil // Return the list of groups
}

func CreateGroup(group *models.Group, userID uint) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	group.CreatedBy = userID
	result := db.Conn.Create(group)
	if result.Error != nil {
		return result.Error // Error creating group
	}
	return nil // Group created successfully
}

func UpdateGroup(group *models.Group) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Save(group)
	if result.Error != nil {
		return result.Error // Error updating group
	}
	return nil // Group updated successfully
}

func DeleteGroup(group *models.Group) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Delete(group)
	if result.Error != nil {
		return result.Error // Error deleting group
	}
	return nil // Group deleted successfully
}