package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type ParticipantRepository struct {
	DB *gorm.DB
}

func CreateParticipantRepository(db *gorm.DB) *ParticipantRepository {
	return &ParticipantRepository{DB: db}
}

func (entity *ParticipantRepository) Save(model *entities.Participant) (*entities.Participant, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *ParticipantRepository) Search(where string) ([]entities.Participant, error) {
	var model []entities.Participant
	result := entity.DB.Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

func (entity *ParticipantRepository) GetParticipantId(value string, err error) (uint64, error) {
	var model entities.Participant
	result := entity.DB.First(&model, "OriginCode = ?", value)
	if result.Error != nil {
		return 0, result.Error
	}
	return model.Id, nil
}

func (entity *ParticipantRepository) GetDocumentId(value string) (uint64, error) {
	var model entities.Participant
	var cond string = "Document = ?"
	if len(value) == 11 {
		cond = "DocumentCpf = ?"
	}
	result := entity.DB.First(&model, cond, value)
	if result.Error != nil {
		return 0, result.Error
	}
	return model.Id, nil
}
