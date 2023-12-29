package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type CityRepository struct {
	DB *gorm.DB
}

func CreateCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{DB: db}
}

func (entity *CityRepository) Save(model *entities.City) (*entities.City, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *CityRepository) GetCityId(value string) (uint64, error) {
	var model entities.City
	result := entity.DB.First(&model, "CityCode = ?", value)
	if result.Error != nil {
		return 0, result.Error
	}
	return model.Id, nil
}