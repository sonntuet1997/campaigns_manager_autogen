package mappers

import (
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/resources"
)

func ToLockedSlotEntity(resource *resources.UpsertLockedSlotRequest) *entities.LockedSlot {
	return &entities.LockedSlot{
		Code: resource.Code,
		Name: resource.Name,
	}
}
