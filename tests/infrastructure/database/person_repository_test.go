package tests

import (
	"errors"
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidPersonCreateDataBase(t *testing.T) {
	database.Initialize(false)
	perRepo := &database.PersonRepository{DB: database.DB}
	tenRepo := &database.TenantRepository{DB: database.DB}

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

	pers, err := perRepo.Search("Id >= 1")
	if err != nil || len(pers) == 0 {
		per, err := entities.NewPerson(tenantId, "teste", "12345678901", "", 1, 1, "", "", "", "", "")
		assert.NoError(t, err, nil)
		err = per.Validate()
		assert.NoError(t, err, nil)
		per, err = perRepo.Save(per)
		assert.NoError(t, err, nil)
		if per.Id == 0 {
			err = errors.New("invalid value ID")
			assert.NoError(t, err, nil)
		}
		pers = append(pers, *per)
	}

	if pers[0].Id == 0 {
		err = errors.New("invalid value ID")
		assert.NoError(t, err, nil)
	}
}
