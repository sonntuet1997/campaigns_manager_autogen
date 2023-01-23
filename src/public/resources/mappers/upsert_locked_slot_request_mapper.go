package mappers

import (
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
)

func ToLockedSlotEntity(resource *resources.UpsertLockedSlotRequest) *entities.LockedSlot {
	return &entities.LockedSlot{
		Code: resource.Code,
		Name: resource.Name,
	}
}
