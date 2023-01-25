package services

import (
	"context"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/usecases"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/resources"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/resources/mappers"
)

type CampaignService struct {
	GetCampaignUsecase    *usecases.GetCampaignUsecase
	CreateCampaignUsecase *usecases.CreateCampaignUsecase
	UpdateCampaignUsecase *usecases.UpdateCampaignUsecase
}

func NewCampaignService(
	GetCampaignUsecase *usecases.GetCampaignUsecase,
	CreateCampaignUsecase *usecases.CreateCampaignUsecase,
	UpdateCampaignUsecase *usecases.UpdateCampaignUsecase,
) CampaignServiceable {
	return &CampaignService{
		GetCampaignUsecase:    GetCampaignUsecase,
		CreateCampaignUsecase: CreateCampaignUsecase,
		UpdateCampaignUsecase: UpdateCampaignUsecase,
	}
}

func (c *CampaignService) GetAllCampaign(ctx context.Context, filter *entities.CampaignFilter) ([]*resources.Campaign, error) {
	records, err := c.GetCampaignUsecase.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}
	return mappers.ToCampaignResources(records), nil
}

func (c *CampaignService) CountCampaign(ctx context.Context, filter *entities.CampaignFilter) (int64, error) {
	return c.GetCampaignUsecase.Count(ctx, filter)
}

func (c *CampaignService) GetCampaign(ctx context.Context, code string) (*resources.Campaign, error) {
	result, err := c.GetCampaignUsecase.Get(ctx, code)
	if err != nil {
		return nil, err
	}
	return mappers.ToCampaignResource(result), nil
}

func (c *CampaignService) CreateCampaign(ctx context.Context, entity *entities.Campaign) (*resources.Campaign, error) {
	result, err := c.CreateCampaignUsecase.Create(ctx, entity)
	if err != nil {
		return nil, err
	}
	return mappers.ToCampaignResource(result), nil
}

func (c *CampaignService) UpdateCampaign(ctx context.Context, entity *entities.Campaign) (*resources.Campaign, error) {
	result, err := c.UpdateCampaignUsecase.Update(ctx, entity)
	if err != nil {
		return nil, err
	}
	return mappers.ToCampaignResource(result), nil
}
