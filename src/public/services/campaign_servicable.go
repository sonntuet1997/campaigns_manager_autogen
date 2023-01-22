package services

import (
	"context"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
)

type CampaignServiceable interface {
	GetAllCampaign(ctx context.Context, filter *entities.CampaignFilter) ([]*resources.Campaign, error)
	CountCampaign(ctx context.Context, filter *entities.CampaignFilter) (int64, error)
	GetCampaign(ctx context.Context, code string) (*resources.Campaign, error)
	CreateCampaign(ctx context.Context, entity *entities.Campaign) (*resources.Campaign, error)
	UpdateCampaign(ctx context.Context, entity *entities.Campaign) (*resources.Campaign, error)
}
