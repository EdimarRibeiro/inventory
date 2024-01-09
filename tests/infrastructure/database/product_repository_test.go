package tests

import (
	"errors"
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidProductCreateDataBase(t *testing.T) {
	database.Initialize(false)
	prodRepo := &database.ProductRepository{DB: database.DB}
	tenRepo := &database.TenantRepository{DB: database.DB}
	perRepo := &database.PersonRepository{DB: database.DB}

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

	prods, err := prodRepo.Search("Id >= 1")
	if err != nil || len(prods) == 0 {
		prod, err := entities.NewProduct(tenantId, "teste", "1234567890123", "", "UN", "0", "", "", "", "", 0, "", "")
		assert.NoError(t, err, nil)
		err = prod.Validate()
		assert.NoError(t, err, nil)
		prod, err = prodRepo.Save(prod)
		assert.NoError(t, err, nil)
		if prod.Id == 0 {
			err = errors.New("invalid value ID")
			assert.NoError(t, err, nil)
		}
		prods = append(prods, *prod)
	}

	if prods[0].Id == 0 {
		err = errors.New("invalid value ID")
		assert.NoError(t, err, nil)
	}
}
