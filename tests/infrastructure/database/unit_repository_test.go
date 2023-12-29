package tests

import (
	"errors"
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidUnitCreateDataBase(t *testing.T) {
	database.Initialize(false)
	unitRepo := &database.UnitRepository{DB: database.DB}
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

	units, err := unitRepo.Search("Id = 'UN'")
	if err != nil || len(units) == 0 {
		unit, err := entities.NewUnit(tenantId, "UN", "teste UN")
		assert.NoError(t, err, nil)
		err = unit.Validate()
		assert.NoError(t, err, nil)
		unit, err = unitRepo.Save(unit)
		assert.NoError(t, err, nil)
		if unit.Id != "UN" {
			err = errors.New("invalid value ID")
			assert.NoError(t, err, nil)
		}
		units = append(units, *unit)
	}

	if units[0].Id != "UN" {
		err = errors.New("invalid value ID")
		assert.NoError(t, err, nil)
	}
}
