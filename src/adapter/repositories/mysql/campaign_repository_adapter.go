package mysql

import (
	"context"
	"errors"
	"gitlab.com/golibs-starter/golib-data/datasource"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/adapter/repositories/mysql/mappers"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/adapter/repositories/mysql/models"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/adapter/repositories/mysql/scopes"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/constants"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/ports"
	"gorm.io/gorm"
)

type CampaignRepositoryAdapter struct {
	db                   *gorm.DB
	dataSourceProperties *datasource.Properties
}

func NewCampaignRepositoryAdapter(
	db *gorm.DB,
	dataSourceProperties *datasource.Properties,
) ports.CampaignRepository {
	return &CampaignRepositoryAdapter{
		db:                   db,
		dataSourceProperties: dataSourceProperties,
	}
}

func (c *CampaignRepositoryAdapter) GetAll(ctx context.Context, filter *entities.CampaignFilter) ([]*entities.Campaign, error) {
	var models []*models.Campaign
	tx := c.db.
		WithContext(ctx).
		Scopes(scopes.WithCampaignPredicates(filter), scopes.WithPaging(filter.PagingFilter)).
		Order("id asc").
		Find(&models)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mappers.ToCampaignEntities(models), nil
}

func (c *CampaignRepositoryAdapter) Get(ctx context.Context, code string) (*entities.Campaign, error) {
	var model *models.Campaign
	if err := c.db.
		WithContext(ctx).
		Where("code = ?", code).
		First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.ErrOrmRecordNotFound
		}
		return nil, err
	}
	return mappers.ToCampaignEntity(model), nil

}

func (c *CampaignRepositoryAdapter) Count(ctx context.Context, filter *entities.CampaignFilter) (int64, error) {
	var count int64
	tx := c.db.
		WithContext(ctx).
		Scopes(scopes.WithCampaignPredicates(filter)).Model(&models.Campaign{}).
		Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return count, nil
}

func (c *CampaignRepositoryAdapter) Create(ctx context.Context, entity *entities.Campaign) (*entities.Campaign, error) {
	model := mappers.ToCampaignModel(entity)
	if err := c.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&model).Error
	}); err != nil {
		return nil, err
	}
	return mappers.ToCampaignEntity(model), nil
}

func (c *CampaignRepositoryAdapter) Update(ctx context.Context, entity *entities.Campaign) (*entities.Campaign, error) {
	err := c.db.WithContext(ctx).Model(&models.Campaign{}).
		Where("code = ?", entity.Code).
		Updates(mappers.ToCampaignModel(entity)).Error
	return entity, err
}
