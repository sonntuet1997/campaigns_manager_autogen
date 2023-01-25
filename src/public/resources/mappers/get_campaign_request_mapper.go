package mappers

import (
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/resources"
)

func ToCampaignFilter(resource *resources.GetCampaignRequest) *entities.CampaignFilter {
	return &entities.CampaignFilter{
		Code: resource.Code,
		Name: resource.Name,
	}
}
