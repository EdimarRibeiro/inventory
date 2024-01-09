package tests

import (
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/infrastructure"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidImportFileXml(t *testing.T) {
	database.Initialize(false)
	inveRep := database.CreateInventoryFileRepository(database.DB)
	invPRep := database.CreateInventoryProductRepository(database.DB)
	unitRep := database.CreateUnitRepository(database.DB)
	uniCRep := database.CreateUnitConvertRepository(database.DB)
	partRep := database.CreateParticipantRepository(database.DB)
	prodRep := database.CreateProductRepository(database.DB)
	docuRep := database.CreateDocumentRepository(database.DB)
	docIRep := database.CreateDocumentItemRepository(database.DB)

	impt := infrastructure.CreateImportFileXml(inveRep, invPRep, unitRep, uniCRep, partRep, prodRep, docuRep, docIRep)

	invs, err := inveRep.Search("InventoryId >= 1 and FileType ='xml'")
	assert.NoError(t, err, nil)

	for i := 0; i < len(invs); i++ {
		item := invs[i]

		err = impt.Execute(&item)
		assert.NoError(t, err, nil)
	}

}
