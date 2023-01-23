package usecases

import (
	"context"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/core/ports"
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
