package scopes

import (
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gorm.io/gorm"
)

func WithLockedSlotPredicates(filter *entities.LockedSlotFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.Code != "" {
			db.Where("code = ?", filter.Code)
		}
		if filter.Name != "" {
			db.Where("name = ?", filter.Name)
		}
		return db
	}
}
