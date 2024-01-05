package tests

import (
	"errors"
	"strconv"
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/infrastructure"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidCalculateSumaryQuantityDataBase(t *testing.T) {
	database.Initialize(false)

	invRepo := database.CreateInventoryRepository(database.DB)
	prodRepo := database.CreateInventoryProductRepository(database.DB)
	docItemRepo := database.CreateDocumentItemRepository(database.DB)

	inves, err := invRepo.Search("Inventory.Id>= 1 and Cloused=0")
	assert.NoError(t, err, nil)

	calc := infrastructure.CreateCalculateBalanceQuantityData(prodRepo, docItemRepo)

	err = calc.Execute(inves[0].Id)
	assert.NoError(t, err, nil)

	resps, err := prodRepo.Search("InventoryId >=" + strconv.FormatUint(inves[0].Id, 10))
	assert.NoError(t, err, nil)

	for i := 0; i < len(resps); i++ {
		item := resps[i]

		if item.Quantity < 0 {
			err = errors.New("invalid by zero inicial Quantity")
		}
		assert.NoError(t, err, nil)

		if item.InputQuantity < 0 {
			err = errors.New("invalid by zero Input Quantity")
		}
		assert.NoError(t, err, nil)

		if item.OutputQuantity < 0 {
			err = errors.New("invalid by zero Output Quantity")
		}
		assert.NoError(t, err, nil)
	}
}
