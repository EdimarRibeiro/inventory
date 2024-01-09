package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type UnitRepositoryInterface interface {
	Save(model *entities.Unit) (*entities.Unit, error)
	Search(where string) ([]entities.Unit, error)
}
