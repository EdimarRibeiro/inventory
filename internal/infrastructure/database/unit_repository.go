package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type UnitRepository struct {
	DB *gorm.DB
}

func CreateUnitRepository(db *gorm.DB) *UnitRepository {
	return &UnitRepository{DB: db}
}

func (entity *UnitRepository) Save(model *entities.Unit) (*entities.Unit, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *UnitRepository) Search(where string) ([]entities.Unit, error) {
	var model []entities.Unit
	result := entity.DB.Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
