package mappers

import (
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
)

func ToCampaignEntity(resource *resources.UpsertCampaignRequest) *entities.Campaign {
	return &entities.Campaign{
		Code: resource.Code,
	}
}
