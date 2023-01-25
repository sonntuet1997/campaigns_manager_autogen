package usecases

import (
	"context"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/ports"
)

type GetLockedSlotUsecase struct {
	Repository ports.LockedSlotRepository
}

func NewGetLockedSlotUsecase(
	Repository ports.LockedSlotRepository,
) *GetLockedSlotUsecase {
	return &GetLockedSlotUsecase{
		Repository: Repository,
	}
}

func (c *GetLockedSlotUsecase) Get(ctx context.Context, code string) (*entities.LockedSlot, error) {
	return c.Repository.Get(ctx, code)
}

func (c *GetLockedSlotUsecase) GetAll(ctx context.Context, filter *entities.LockedSlotFilter) ([]*entities.LockedSlot, error) {
	return c.Repository.GetAll(ctx, filter)
}

func (c *GetLockedSlotUsecase) Count(ctx context.Context, filter *entities.LockedSlotFilter) (int64, error) {
	return c.Repository.Count(ctx, filter)
}
