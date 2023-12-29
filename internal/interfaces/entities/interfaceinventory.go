package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type InventoryRepositoryInterface interface {
	Save(model *entities.Inventory) (*entities.Inventory, error)
	Search(where string) ([]entities.Inventory, error)
}
