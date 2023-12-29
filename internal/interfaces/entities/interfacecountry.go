package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type CountryRepositoryInterface interface {
	Save(model *entities.Country) (*entities.Country, error)
	GetCountryId(value string) (uint64, error)
}
