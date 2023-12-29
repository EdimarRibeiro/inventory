package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type PersonRepositoryInterface interface {
	Save(model *entities.Person) (*entities.Person, error)
	Search(where string) ([]entities.Person, error)
}
