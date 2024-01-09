package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type InventoryRepository struct {
	DB *gorm.DB
}

func CreateInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{DB: db}
}

func (entity *InventoryRepository) Save(model *entities.Inventory) (*entities.Inventory, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *InventoryRepository) Search(where string) ([]entities.Inventory, error) {
	var model []entities.Inventory
	result := entity.DB.InnerJoins("Participant").Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
