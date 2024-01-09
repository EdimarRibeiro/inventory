package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type InventoryFileRepositoryInterface interface {
	Save(model *entities.InventoryFile) (*entities.InventoryFile, error)
	Search(where string) ([]entities.InventoryFile, error)
}
