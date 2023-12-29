package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type DocumentRepository struct {
	DB *gorm.DB
}

func CreateDocumentRepository(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{DB: db}
}

func (entity *DocumentRepository) Save(model *entities.Document) (*entities.Document, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *DocumentRepository) Search(where string) ([]entities.Document, error) {
	var model []entities.Document
	result := entity.DB.Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
