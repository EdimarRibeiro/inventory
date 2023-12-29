package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type DocumentRepositoryInterface interface {
	Save(model *entities.Document) (*entities.Document, error)
	Search(where string) ([]entities.Document, error)
}
