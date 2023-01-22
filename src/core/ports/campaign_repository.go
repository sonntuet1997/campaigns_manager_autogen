package ports

import (
	"context"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
)

type CampaignRepository interface {
	GetAll(ctx context.Context, filter *entities.CampaignFilter) ([]*entities.Campaign, error)
	Count(ctx context.Context, filter *entities.CampaignFilter) (int64, error)
	Get(ctx context.Context, code string) (*entities.Campaign, error)
	Update(ctx context.Context, entity *entities.Campaign) (*entities.Campaign, error)
	Create(ctx context.Context, entity *entities.Campaign) (*entities.Campaign, error)
}
