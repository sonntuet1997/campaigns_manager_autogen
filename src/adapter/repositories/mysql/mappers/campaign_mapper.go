package mappers

import (
	"gitlab.com/technixo/backend/campaigns_manager_autogen/adapter/repositories/mysql/models"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/entities"
)

func ToCampaignEntity(model *models.Campaign) *entities.Campaign {
	return &entities.Campaign{
		ID:   model.ID,
		Code: model.Code,
		Name: model.Name,
	}
}

func ToCampaignEntities(arr []*models.Campaign) []*entities.Campaign {
	result := make([]*entities.Campaign, 0, len(arr))
	for _, model := range arr {
		result = append(result, ToCampaignEntity(model))
	}
	return result
}

func ToCampaignModel(entity *entities.Campaign) *models.Campaign {
	return &models.Campaign{
		ID:   entity.ID,
		Code: entity.Code,
		Name: entity.Name,
	}
}

func ToCampaignModels(arr []*entities.Campaign) []*models.Campaign {
	result := make([]*models.Campaign, 0, len(arr))
	for _, entity := range arr {
		result = append(result, ToCampaignModel(entity))
	}
	return result
}
