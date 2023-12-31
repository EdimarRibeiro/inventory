package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type ProductRepositoryInterface interface {
	Save(model *entities.Product) (*entities.Product, error)
	GetProductId(value string, err error) (uint64, error)
	Search(where string) ([]entities.Product, error)
}
