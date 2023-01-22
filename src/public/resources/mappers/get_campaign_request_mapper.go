package mappers

import (
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
)

func ToCampaignFilter(resource *resources.GetCampaignRequest) *entities.CampaignFilter {
	return &entities.CampaignFilter{}
}
