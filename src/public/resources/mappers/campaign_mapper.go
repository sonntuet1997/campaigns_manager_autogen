package mappers

import (
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
)

func ToCampaignResource(entity *entities.Campaign) *resources.Campaign {
	return &resources.Campaign{
		ID:   entity.ID,
		Code: entity.Code,
	}
}

func ToCampaignResources(arr []*entities.Campaign) []*resources.Campaign {
	result := make([]*resources.Campaign, 0, len(arr))
	for _, entity := range arr {
		result = append(result, ToCampaignResource(entity))
	}
	return result
}
