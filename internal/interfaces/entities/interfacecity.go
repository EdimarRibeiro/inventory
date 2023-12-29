package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type CityRepositoryInterface interface {
	Save(model *entities.City) (*entities.City, error)
	GetCityId(value string) (uint64, error)
}
