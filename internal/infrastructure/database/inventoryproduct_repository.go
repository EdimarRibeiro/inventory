package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type InventoryProductRepository struct {
	DB *gorm.DB
}

func CreateInventoryProductRepository(db *gorm.DB) *InventoryProductRepository {
	return &InventoryProductRepository{DB: db}
}

func (entity *InventoryProductRepository) Save(model *entities.InventoryProduct) (*entities.InventoryProduct, error) {
	var result *gorm.DB
	if model.ParticipantId != nil && *model.ParticipantId == 0 {
		model.ParticipantId = nil
	}
	result = entity.DB.Save(&model)
	return model, result.Error
}

func (entity *InventoryProductRepository) Search(where string) ([]entities.InventoryProduct, error) {
	var model []entities.InventoryProduct
	result := entity.DB.InnerJoins("Product").Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
