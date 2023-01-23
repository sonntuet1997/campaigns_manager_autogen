package usecases

import (
	"context"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/core/ports"
)

type CreateLockedSlotUsecase struct {
	Repository ports.LockedSlotRepository
}

func NewCreateLockedSlotUsecase(
	Repository ports.LockedSlotRepository,
) *CreateLockedSlotUsecase {
	return &CreateLockedSlotUsecase{
		Repository: Repository,
	}
}

func (c *CreateLockedSlotUsecase) Create(ctx context.Context, entity *entities.LockedSlot) (*entities.LockedSlot, error) {
	return c.Repository.Create(ctx, entity)
}
