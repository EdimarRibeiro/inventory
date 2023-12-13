package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type UnitConvertRepositoryInterface interface {
	Save(model *entities.UnitConvert) (entities.UnitConvert, error)
	Search(where string) []entities.UnitConvert
}
