package usecases

import (
	"context"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/core/ports"
)

type UpdateCampaignUsecase struct {
	Repository ports.CampaignRepository
}

func NewUpdateCampaignUsecase(
	Repository ports.CampaignRepository,
) *UpdateCampaignUsecase {
	return &UpdateCampaignUsecase{
		Repository: Repository,
	}
}

func (c *UpdateCampaignUsecase) Update(ctx context.Context, entity *entities.Campaign) (*entities.Campaign, error) {
	return c.Repository.Update(ctx, entity)
}
