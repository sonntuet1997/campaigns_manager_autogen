package services

import (
	"context"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/core/usecases"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources/mappers"
)

type LockedSlotService struct {
	GetLockedSlotUsecase    *usecases.GetLockedSlotUsecase
	CreateLockedSlotUsecase *usecases.CreateLockedSlotUsecase
	UpdateLockedSlotUsecase *usecases.UpdateLockedSlotUsecase
}

func NewLockedSlotService(
	GetLockedSlotUsecase *usecases.GetLockedSlotUsecase,
	CreateLockedSlotUsecase *usecases.CreateLockedSlotUsecase,
	UpdateLockedSlotUsecase *usecases.UpdateLockedSlotUsecase,
) LockedSlotServiceable {
	return &LockedSlotService{
		GetLockedSlotUsecase:    GetLockedSlotUsecase,
		CreateLockedSlotUsecase: CreateLockedSlotUsecase,
		UpdateLockedSlotUsecase: UpdateLockedSlotUsecase,
	}
}

func (c *LockedSlotService) GetAllLockedSlot(ctx context.Context, filter *entities.LockedSlotFilter) ([]*resources.LockedSlot, error) {
	records, err := c.GetLockedSlotUsecase.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}
	return mappers.ToLockedSlotResources(records), nil
}

func (c *LockedSlotService) CountLockedSlot(ctx context.Context, filter *entities.LockedSlotFilter) (int64, error) {
	return c.GetLockedSlotUsecase.Count(ctx, filter)
}

func (c *LockedSlotService) GetLockedSlot(ctx context.Context, code string) (*resources.LockedSlot, error) {
	result, err := c.GetLockedSlotUsecase.Get(ctx, code)
	if err != nil {
		return nil, err
	}
	return mappers.ToLockedSlotResource(result), nil
}

func (c *LockedSlotService) CreateLockedSlot(ctx context.Context, entity *entities.LockedSlot) (*resources.LockedSlot, error) {
	result, err := c.CreateLockedSlotUsecase.Create(ctx, entity)
	if err != nil {
		return nil, err
	}
	return mappers.ToLockedSlotResource(result), nil
}

func (c *LockedSlotService) UpdateLockedSlot(ctx context.Context, entity *entities.LockedSlot) (*resources.LockedSlot, error) {
	result, err := c.UpdateLockedSlotUsecase.Update(ctx, entity)
	if err != nil {
		return nil, err
	}
	return mappers.ToLockedSlotResource(result), nil
}
