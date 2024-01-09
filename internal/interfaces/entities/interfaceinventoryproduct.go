package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type InventoryProductRepositoryInterface interface {
	Save(model *entities.InventoryProduct) (*entities.InventoryProduct, error)
	Search(where string) ([]entities.InventoryProduct, error)
}
