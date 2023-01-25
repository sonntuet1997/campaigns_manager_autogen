package scopes

import (
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
	"gorm.io/gorm"
)

func WithCampaignPredicates(filter *entities.CampaignFilter) func(db *gorm.DB) *gorm.DB {
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
