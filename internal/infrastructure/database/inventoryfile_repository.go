package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type InventoryFileRepository struct {
	DB *gorm.DB
}

func CreateInventoryFileRepository(db *gorm.DB) *InventoryFileRepository {
	return &InventoryFileRepository{DB: db}
}

func (entity *InventoryFileRepository) Save(model *entities.InventoryFile) (*entities.InventoryFile, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *InventoryFileRepository) Search(where string) ([]entities.InventoryFile, error) {
	var model []entities.InventoryFile
	result := entity.DB.Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
