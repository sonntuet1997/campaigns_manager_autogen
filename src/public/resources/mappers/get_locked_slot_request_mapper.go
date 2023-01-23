package mappers

import (
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
)

func ToLockedSlotFilter(resource *resources.GetLockedSlotRequest) *entities.LockedSlotFilter {
	return &entities.LockedSlotFilter{
		Code: resource.Code,
		Name: resource.Name,
	}
}
