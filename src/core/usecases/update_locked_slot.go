package usecases

import (
	"context"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/ports"
)

type UpdateLockedSlotUsecase struct {
	Repository ports.LockedSlotRepository
}

func NewUpdateLockedSlotUsecase(
	Repository ports.LockedSlotRepository,
) *UpdateLockedSlotUsecase {
	return &UpdateLockedSlotUsecase{
		Repository: Repository,
	}
}

func (c *UpdateLockedSlotUsecase) Update(ctx context.Context, entity *entities.LockedSlot) (*entities.LockedSlot, error) {
	return c.Repository.Update(ctx, entity)
}
