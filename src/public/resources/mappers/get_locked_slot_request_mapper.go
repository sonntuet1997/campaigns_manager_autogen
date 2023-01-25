package mappers

import (
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/resources"
)

func ToLockedSlotFilter(resource *resources.GetLockedSlotRequest) *entities.LockedSlotFilter {
	return &entities.LockedSlotFilter{
		Code: resource.Code,
		Name: resource.Name,
	}
}
