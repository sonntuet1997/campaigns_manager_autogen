package scopes

import (
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gorm.io/gorm"
)

func WithCampaignPredicates(filter *entities.CampaignFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.Code != "" {
			db.Where("code = ?", filter.Code)
		}
		return db
	}
}