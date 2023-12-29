package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type TenantRepository struct {
	DB *gorm.DB
}

func CreateTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{DB: db}
}

func (entity *TenantRepository) Save(model *entities.Tenant) (*entities.Tenant, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *TenantRepository) Search(where string) ([]entities.Tenant, error) {
	var model []entities.Tenant
	result := entity.DB.Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
