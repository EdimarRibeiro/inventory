package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func CreateProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (entity *ProductRepository) Save(model *entities.Product) (*entities.Product, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *ProductRepository) Search(where string) ([]entities.Product, error) {
	var model []entities.Product
	result := entity.DB.Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

func (entity *ProductRepository) GetProductId(value string, err error) (uint64, error) {
	var model entities.Product
	result := entity.DB.First(&model, "OriginCode = ?", value)
	if result.Error != nil {
		return 0, result.Error
	}
	return model.Id, nil
}

func (entity *ProductRepository) GetProductIdBarCode(value string, err error) (uint64, error) {
	var model entities.Product
	result := entity.DB.First(&model, "BarCode = ?", value)
	if result.Error != nil {
		return 0, result.Error
	}
	return model.Id, nil
}
