package tests

import (
	"errors"
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/infrastructure"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/EdimarRibeiro/inventory/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestValidSumaryQuantityDocumentItem(t *testing.T) {
	database.Initialize(false)
	prodRepo := database.CreateInventoryProductRepository(database.DB)
	docItemRepo := database.CreateDocumentItemRepository(database.DB)

	resps, err := prodRepo.Search("InventoryId >= 1")
	assert.NoError(t, err, nil)

	sum := infrastructure.CreateSumaryQuantityDocumentItem(docItemRepo)
	sumaries, err := sum.Execute(resps[0].InventoryId)
	assert.NoError(t, err, nil)

	for i := 0; i < len(resps); i++ {
		sumary := models.LocateSumary(*sumaries, resps[i].ProductId)
		if sumary != nil {
			if sumary.SumaryQuantityInput != 0 {
				err = errors.New("SumaryQuantityInput not is zero")
			}
			assert.NoError(t, err, nil)

			if sumary.SumaryQuantityOutput != 10 {
				err = errors.New("SumaryQuantityOutput not is 10")
			}
			assert.NoError(t, err, nil)
		}
	}
}
