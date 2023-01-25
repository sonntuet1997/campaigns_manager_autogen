package usecases

import (
	"context"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/ports"
)

type CreateCampaignUsecase struct {
	Repository ports.CampaignRepository
}

func NewCreateCampaignUsecase(
	Repository ports.CampaignRepository,
) *CreateCampaignUsecase {
	return &CreateCampaignUsecase{
		Repository: Repository,
	}
}

func (c *CreateCampaignUsecase) Create(ctx context.Context, entity *entities.Campaign) (*entities.Campaign, error) {
	return c.Repository.Create(ctx, entity)
}
