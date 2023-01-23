package mysql

import (
	"context"
	"errors"
	"gitlab.com/golibs-starter/golib-data/datasource"
	"gitlab.com/technixo/backend/campaigns-manager/adapter/repositories/mysql/mappers"
	"gitlab.com/technixo/backend/campaigns-manager/adapter/repositories/mysql/models"
	"gitlab.com/technixo/backend/campaigns-manager/adapter/repositories/mysql/scopes"
	"gitlab.com/technixo/backend/campaigns-manager/core/constants"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/core/ports"
	"gorm.io/gorm"
)

type LockedSlotRepositoryAdapter struct {
	db                   *gorm.DB
	dataSourceProperties *datasource.Properties
}

func NewLockedSlotRepositoryAdapter(
	db *gorm.DB,
	dataSourceProperties *datasource.Properties,
) ports.LockedSlotRepository {
	return &LockedSlotRepositoryAdapter{
		db:                   db,
		dataSourceProperties: dataSourceProperties,
	}
}

func (c *LockedSlotRepositoryAdapter) GetAll(ctx context.Context, filter *entities.LockedSlotFilter) ([]*entities.LockedSlot, error) {
	var models []*models.LockedSlot
	tx := c.db.
		WithContext(ctx).
		Scopes(scopes.WithLockedSlotPredicates(filter), scopes.WithPaging(filter.PagingFilter)).
		Order("id asc").
		Find(&models)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mappers.ToLockedSlotEntities(models), nil
}

func (c *LockedSlotRepositoryAdapter) Get(ctx context.Context, code string) (*entities.LockedSlot, error) {
	var model *models.LockedSlot
	if err := c.db.
		WithContext(ctx).
		Where("code = ?", code).
		First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.ErrOrmRecordNotFound
		}
		return nil, err
	}
	return mappers.ToLockedSlotEntity(model), nil

}

func (c *LockedSlotRepositoryAdapter) Count(ctx context.Context, filter *entities.LockedSlotFilter) (int64, error) {
	var count int64
	tx := c.db.
		WithContext(ctx).
		Scopes(scopes.WithLockedSlotPredicates(filter)).Model(&models.LockedSlot{}).
		Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return count, nil
}

func (c *LockedSlotRepositoryAdapter) Create(ctx context.Context, entity *entities.LockedSlot) (*entities.LockedSlot, error) {
	model := mappers.ToLockedSlotModel(entity)
	if err := c.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&model).Error
	}); err != nil {
		return nil, err
	}
	return mappers.ToLockedSlotEntity(model), nil
}

func (c *LockedSlotRepositoryAdapter) Update(ctx context.Context, entity *entities.LockedSlot) (*entities.LockedSlot, error) {
	err := c.db.WithContext(ctx).Model(&models.LockedSlot{}).
		Where("code = ?", entity.Code).
		Updates(mappers.ToLockedSlotModel(entity)).Error
	return entity, err
}
