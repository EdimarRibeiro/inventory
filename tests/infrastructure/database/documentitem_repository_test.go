package tests

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidDocumentItemCreateDataBase(t *testing.T) {
	database.Initialize(false)
	partRepo := &database.ParticipantRepository{DB: database.DB}
	invRepo := &database.InventoryRepository{DB: database.DB}
	tenRepo := &database.TenantRepository{DB: database.DB}
	perRepo := &database.PersonRepository{DB: database.DB}
	docRepo := &database.DocumentRepository{DB: database.DB}
	docIteRepo := &database.DocumentItemRepository{DB: database.DB}
	prodRepo := &database.ProductRepository{DB: database.DB}
	unitRepo := &database.UnitRepository{DB: database.DB}

	var tenantId uint64 = 0

	tens, err := tenRepo.Search("Id >= 1")
	if err != nil || len(tens) == 0 {
		ten, err := entities.NewTenant("Teste", "09066936754", 0)
		assert.NoError(t, err, nil)

		ten, err = tenRepo.Save(ten)
		assert.NoError(t, err, nil)
		if ten.Id == 0 {
			err = errors.New("invalid value ID")
			assert.NoError(t, err, nil)
		}

		ten.PersonId, err = GetPersonData(perRepo, ten.Id)
		assert.NoError(t, err, nil)

		ten, err = tenRepo.Save(ten)
		assert.NoError(t, err, nil)

		tens = append(tens, *ten)
	}

	if tens[0].Id == 0 {
		err = errors.New("invalid value ID")
		assert.NoError(t, err, nil)
	}
	tenantId = tens[len(tens)-1].Id

	invs, err := invRepo.Search("Inventory.Id >= 1")
	assert.NoError(t, err, nil)
	if err != nil || len(invs) == 0 {
		participantId, err := GetParticipantData(partRepo, tenantId)
		assert.NoError(t, err, nil)
		inv, err := entities.NewInventory(tenantId, participantId, "Teste")
		assert.NoError(t, err, nil)
		inv, err = invRepo.Save(inv)
		assert.NoError(t, err, nil)
		invs = append(invs, *inv)
	}
	var inventoryId uint64 = invs[len(invs)-1].Id

	docs, err := docRepo.Search("Id >= 1")
	assert.NoError(t, err, nil)

	if len(docs) == 0 {
		participantId, err := GetParticipantData(partRepo, tenantId)
		assert.NoError(t, err, nil)

		model, err := entities.NewDocument(inventoryId, participantId, "", "", "", "", "", "", "", time.Now(), time.Now(), 0, "", 0, 0, 0, "", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, false, "xml")
		assert.NoError(t, err, nil)

		doc, err := docRepo.Save(model)
		assert.NoError(t, err, nil)

		if doc.Id == 0 {
			err = errors.New("invalid id of document")
		}
		assert.NoError(t, err, nil)
		docs = append(docs, *doc)
	}
	var documentId uint64 = docs[len(docs)-1].Id

	docItens, err := docIteRepo.Search("DocumentId >= " + strconv.FormatUint(documentId, 10))
	assert.NoError(t, err, nil)

	if len(docItens) == 0 {
		produtoId, err := GetProductData(prodRepo, unitRepo, tenantId)
		assert.NoError(t, err, nil)

		model, err := entities.NewDocumentItem(inventoryId, inventoryId, "001", produtoId, "", 10, "UN", 1, 0, "0", "000", "5101", "", 0, 0, 0, 0, 0, 0, "", "00", "", 0, 0, 0, "00", 0, 0, 0, 0, 0, "00", 0, 0, 0, 0, 0, "", 0)
		assert.NoError(t, err, nil)

		docitem, err := docIteRepo.Save(model)
		assert.NoError(t, err, nil)

		if docitem.Sequency == "" {
			err = errors.New("invalid sequency of item document")
		}
		assert.NoError(t, err, nil)
		//docItens = append(docItens, *docitem)
	}
}

func GetProductData(proRep *database.ProductRepository, unitRepo *database.UnitRepository, tenantId uint64) (uint64, error) {
	units, err := unitRepo.Search("Id='UN'")

	if len(units) == 0 || err != nil {
		unit, err := entities.NewUnit(tenantId, "UN", "Unidade UN")
		if err != nil {
			return 0, err
		}
		if err == nil {
			_, err = unitRepo.Save(unit)
			if err != nil {
				return 0, err
			}
		}
	}

	prod, err := proRep.Search("Id>=1")
	if err != nil || len(prod) == 0 {
		par, err := proRep.Save(CreateProductData(tenantId))
		if err != nil {
			return 0, err
		}
		prod = append(prod, *par)
	}
	return prod[len(prod)-1].Id, nil
}

func CreateProductData(tenantId uint64) *entities.Product {
	return &entities.Product{
		TenantId:      tenantId,
		OriginCode:    "0001",
		Description:   "Teste",
		BarCode:       "",
		OldOriginCode: "",
		UnitId:        "UN",
		Type:          "00",
		NcmCode:       "",
		ExIpi:         "",
		GenderCode:    "",
		ServiceCode:   "",
		AliqIcms:      0,
		CestCode:      "",
	}
}
