package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type DocumentRepository struct {
	DB *gorm.DB
}

func (entity *DocumentRepository) Save(model *entities.Document) (entities.Document, error) {
	result := entity.DB.Save(model)
	return *model, result.Error
}

func (entity *DocumentRepository) Search(value string, err error) ([]entities.Document, error) {
	var model []entities.Document
	result := entity.DB.Where(&model, "DocumentCode = ?", value)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
