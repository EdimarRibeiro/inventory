package database

import (
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func CreateUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (entity *UserRepository) Save(model *entities.User) (*entities.User, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *UserRepository) Search(where string) ([]entities.User, error) {
	var model []entities.User
	result := entity.DB.Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
