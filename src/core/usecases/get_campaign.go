package usecases

import (
	"context"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/core/ports"
)

type GetCampaignUsecase struct {
	Repository ports.CampaignRepository
}

func NewGetCampaignUsecase(
	Repository ports.CampaignRepository,
) *GetCampaignUsecase {
	return &GetCampaignUsecase{
		Repository: Repository,
	}
}

func (c *GetCampaignUsecase) Get(ctx context.Context, code string) (*entities.Campaign, error) {
	return c.Repository.Get(ctx, code)
}

func (c *GetCampaignUsecase) GetAll(ctx context.Context, filter *entities.CampaignFilter) ([]*entities.Campaign, error) {
	return c.Repository.GetAll(ctx, filter)
}
