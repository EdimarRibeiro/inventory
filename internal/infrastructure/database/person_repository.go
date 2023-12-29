package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type PersonRepository struct {
	DB *gorm.DB
}

func CreatePersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{DB: db}
}

func (entity *PersonRepository) Save(model *entities.Person) (*entities.Person, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *PersonRepository) Search(where string) ([]entities.Person, error) {
	var model []entities.Person
	result := entity.DB.Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
