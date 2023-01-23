package mappers

import (
	"gitlab.com/technixo/backend/campaigns-manager/adapter/repositories/mysql/models"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
)

func ToLockedSlotEntity(model *models.LockedSlot) *entities.LockedSlot {
	return &entities.LockedSlot{
		ID:   model.ID,
		Code: model.Code,
		Name: model.Name,
	}
}

func ToLockedSlotEntities(arr []*models.LockedSlot) []*entities.LockedSlot {
	result := make([]*entities.LockedSlot, 0, len(arr))
	for _, model := range arr {
		result = append(result, ToLockedSlotEntity(model))
	}
	return result
}

func ToLockedSlotModel(entity *entities.LockedSlot) *models.LockedSlot {
	return &models.LockedSlot{
		ID:   entity.ID,
		Code: entity.Code,
		Name: entity.Name,
	}
}

func ToLockedSlotModels(arr []*entities.LockedSlot) []*models.LockedSlot {
	result := make([]*models.LockedSlot, 0, len(arr))
	for _, entity := range arr {
		result = append(result, ToLockedSlotModel(entity))
	}
	return result
}
