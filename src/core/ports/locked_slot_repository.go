package ports

import (
	"context"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
)

type LockedSlotRepository interface {
	GetAll(ctx context.Context, filter *entities.LockedSlotFilter) ([]*entities.LockedSlot, error)
	Count(ctx context.Context, filter *entities.LockedSlotFilter) (int64, error)
	Get(ctx context.Context, code string) (*entities.LockedSlot, error)
	Create(ctx context.Context, entity *entities.LockedSlot) (*entities.LockedSlot, error)
	Update(ctx context.Context, entity *entities.LockedSlot) (*entities.LockedSlot, error)
}
