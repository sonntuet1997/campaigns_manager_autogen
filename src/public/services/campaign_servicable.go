package services

import (
	"context"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
)

type CampaignServiceable interface {
	GetAllCampaign(ctx context.Context, filter *entities.CampaignFilter) ([]*resources.Campaign, error)
	CountCampaign(ctx context.Context, filter *entities.CampaignFilter) (*resources.Campaign, error)
	GetCampaign(ctx context.Context, filter *entities.CampaignFilter) (*resources.Campaign, error)
	CreateCampaign(ctx context.Context, filter *entities.CampaignFilter) ([]*resources.Campaign, error)
	UpdateCampaign(ctx context.Context, filter *entities.CampaignFilter) ([]*resources.Campaign, error)
}
