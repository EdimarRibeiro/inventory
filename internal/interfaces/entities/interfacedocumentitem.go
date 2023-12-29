package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/models"
)

type DocumentItemRepositoryInterface interface {
	Save(model *entities.DocumentItem) (*entities.DocumentItem, error)
	Search(where string) ([]entities.DocumentItem, error)
	SumaryQuantity(inventoryId uint64) (*[]models.SumaryQuantityModel, error)
}
