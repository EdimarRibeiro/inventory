package infrastructure

import (
	"errors"

	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
	"github.com/EdimarRibeiro/inventory/internal/models"
)

type SumaryQuantityDocumentItem struct {
	DocumentItem entitiesinterface.DocumentItemRepositoryInterface
}

func CreateSumaryQuantityDocumentItem(documentItem entitiesinterface.DocumentItemRepositoryInterface) *SumaryQuantityDocumentItem {
	return &SumaryQuantityDocumentItem{DocumentItem: documentItem}
}

func (calc *SumaryQuantityDocumentItem) Execute(inventoryId uint64) (*[]models.SumaryQuantityModel, error) {
	if inventoryId == 0 {
		return nil, errors.New("inventoryId is invalid")
	}

	items, err := calc.DocumentItem.SumaryQuantity(inventoryId)

	if err != nil {
		return nil, err
	}
	return items, nil
}
