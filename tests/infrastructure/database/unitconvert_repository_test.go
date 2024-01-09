package tests

import (
	"errors"
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidUnitConvertCreateDataBase(t *testing.T) {
	database.Initialize(false)
	convRepo := &database.UnitConvertRepository{DB: database.DB}
	tenRepo := &database.TenantRepository{DB: database.DB}
	perRepo := &database.PersonRepository{DB: database.DB}
	produRepo := &database.ProductRepository{DB: database.DB}
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

	convs, err := convRepo.Search("ProductId >= 1")
	if err != nil || len(convs) == 0 {
		produtoId, err := GetProductData(produRepo, unitRepo, tenantId)
		assert.NoError(t, err, nil)

		conv, err := entities.NewUnitConvert("UN", produtoId, tenantId, 1, "1234567890123")
		assert.NoError(t, err, nil)
		err = conv.Validate()
		assert.NoError(t, err, nil)
		conv, err = convRepo.Save(conv)
		assert.NoError(t, err, nil)
		if conv.ProductId == 0 {
			err = errors.New("invalid value ProductId")
			assert.NoError(t, err, nil)
		}
		convs = append(convs, *conv)
	}

	if convs[0].ProductId == 0 {
		err = errors.New("invalid value ProductId")
		assert.NoError(t, err, nil)
	}
}
