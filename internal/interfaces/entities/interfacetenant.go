package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type TenantRepositoryInterface interface {
	Save(model *entities.Tenant) (*entities.Tenant, error)
	Search(where string) ([]entities.Tenant, error)
}
