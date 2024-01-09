package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidInventoryProductCreateDataBase(t *testing.T) {
	database.Initialize(false)
	invRepo := &database.InventoryRepository{DB: database.DB}
	prodRepo := &database.InventoryProductRepository{DB: database.DB}
	tenRepo := &database.TenantRepository{DB: database.DB}
	perRepo := &database.PersonRepository{DB: database.DB}
	produRepo := &database.ProductRepository{DB: database.DB}
	unitRepo := &database.UnitRepository{DB: database.DB}
	partRepo := &database.ParticipantRepository{DB: database.DB}

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

	invs, err := invRepo.Search("Id >= 1")
	if err != nil || len(invs) == 0 {
		participantId, err := GetParticipantData(partRepo, tenantId)
		assert.NoError(t, err, nil)
		inv, err := entities.NewInventory(tenantId, participantId, "Dez 2023")
		assert.NoError(t, err, nil)
		err = inv.Validate()
		assert.NoError(t, err, nil)
		inv, err = invRepo.Save(inv)
		assert.NoError(t, err, nil)
		if inv.Id == 0 {
			err = errors.New("invalid value ID")
			assert.NoError(t, err, nil)
		}
		invs = append(invs, *inv)
	}

	if invs[0].Id == 0 {
		err = errors.New("invalid value ID")
		assert.NoError(t, err, nil)
	}
	var inventoryId uint64 = invs[0].Id

	prods, err := prodRepo.Search("ProductId >= 1")
	if err != nil || len(prods) == 0 {
		produtoId, err := GetProductData(produRepo, unitRepo, tenantId)
		assert.NoError(t, err, nil)

		prod, err := entities.NewInventoryProduct(inventoryId, produtoId, "0001", time.Now(), "UN", 10, 1, 10, "0", 0, "", "", 0)
		assert.NoError(t, err, nil)
		err = prod.Validate()
		assert.NoError(t, err, nil)
		prod, err = prodRepo.Save(prod)
		assert.NoError(t, err, nil)
		if prod.ProductId == 0 {
			err = errors.New("invalid value ProductId")
			assert.NoError(t, err, nil)
		}
		prods = append(prods, *prod)
	}

	if prods[0].ProductId == 0 {
		err = errors.New("invalid value ProductId")
		assert.NoError(t, err, nil)
	}
}
