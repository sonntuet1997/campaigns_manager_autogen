package mappers

import (
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/resources"
)

func ToCampaignEntity(resource *resources.UpsertCampaignRequest) *entities.Campaign {
	return &entities.Campaign{
		Code: resource.Code,
		Name: resource.Name,
	}
}
