package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type CountryRepository struct {
	DB *gorm.DB
}

func (entity *CountryRepository) Save(model *entities.Country) (entities.Country, error) {
	result := entity.DB.Save(model)
	return *model, result.Error
}

func (entity *CountryRepository) GetCountryId(value string, err error) (float64, error) {
	var model entities.Country
	result := entity.DB.First(&model, "CountryCode = ?", value)
	if result.Error != nil {
		return 0, result.Error
	}
	return model.Id, nil
}
