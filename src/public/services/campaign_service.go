package services

import (
	"context"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/core/usecases"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources/mappers"
)

type CampaignService struct {
	GetCampaignUsecase *usecases.GetCampaignUsecase
}

func NewCampaignService(
	GetCampaignUsecase *usecases.GetCampaignUsecase,
) CampaignServiceable {
	return &CampaignService{
		GetCampaignUsecase: GetCampaignUsecase,
	}
}

func (c *CampaignService) GetAllCampaign(ctx context.Context, filter *entities.CampaignFilter) ([]*resources.Campaign, error) {
	records, err := c.GetCampaignUsecase.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}
	return mappers.ToCampaignResources(records), nil
}
