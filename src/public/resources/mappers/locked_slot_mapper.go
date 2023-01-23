package mappers

import (
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
)

func ToLockedSlotResource(entity *entities.LockedSlot) *resources.LockedSlot {
	return &resources.LockedSlot{
		ID:   entity.ID,
		Code: entity.Code,
		Name: entity.Name,
	}
}

func ToLockedSlotResources(arr []*entities.LockedSlot) []*resources.LockedSlot {
	result := make([]*resources.LockedSlot, 0, len(arr))
	for _, entity := range arr {
		result = append(result, ToLockedSlotResource(entity))
	}
	return result
}
