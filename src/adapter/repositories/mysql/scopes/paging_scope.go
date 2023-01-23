package scopes

import (
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gorm.io/gorm"
)

func WithPaging(paging *entities.PagingFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if paging == nil {
			return db
		}
		if len(paging.Sort) > 0 {
			db.Order(paging.Sort)
		}
		if paging.Limit > 0 {
			db.Limit(paging.Limit)
		}
		if paging.Page > 0 {
			db.Offset((paging.Page - 1) * paging.Limit)
		}
		return db
	}
}
