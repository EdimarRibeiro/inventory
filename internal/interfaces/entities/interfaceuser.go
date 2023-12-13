package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type UserRepositoryInterface interface {
	Save(model *entities.User) (entities.User, error)
	Search(where string) []entities.User
}
