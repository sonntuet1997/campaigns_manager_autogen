package services

import (
	"context"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
)

type LockedSlotServiceable interface {
	GetAllLockedSlot(ctx context.Context, filter *entities.LockedSlotFilter) ([]*resources.LockedSlot, error)
	CountLockedSlot(ctx context.Context, filter *entities.LockedSlotFilter) (int64, error)
	GetLockedSlot(ctx context.Context, code string) (*resources.LockedSlot, error)
	CreateLockedSlot(ctx context.Context, entity *entities.LockedSlot) (*resources.LockedSlot, error)
	UpdateLockedSlot(ctx context.Context, entity *entities.LockedSlot) (*resources.LockedSlot, error)
}
