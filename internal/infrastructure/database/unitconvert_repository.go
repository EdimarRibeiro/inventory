package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type UnitConvertRepository struct {
	DB *gorm.DB
}

func CreateUnitConvertRepository(db *gorm.DB) *UnitConvertRepository {
	return &UnitConvertRepository{DB: db}
}

func (entity *UnitConvertRepository) Save(model *entities.UnitConvert) (*entities.UnitConvert, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *UnitConvertRepository) Search(where string) ([]entities.UnitConvert, error) {
	var model []entities.UnitConvert
	result := entity.DB.Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
