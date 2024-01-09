package tests

import (
	"errors"
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidParticipantCreateDataBase(t *testing.T) {
	database.Initialize(false)
	parRepo := &database.ParticipantRepository{DB: database.DB}
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

	pars, err := parRepo.Search("Id >= 1")
	if err != nil || len(pars) == 0 {
		par, err := entities.NewParticipant(tenantId, "0001", "teste", "12345678901", "", "", "1234567", "12345", "", "", "", "")
		assert.NoError(t, err, nil)
		err = par.Validate()
		assert.NoError(t, err, nil)
		par, err = parRepo.Save(par)
		assert.NoError(t, err, nil)
		if par.Id == 0 {
			err = errors.New("invalid value ID")
			assert.NoError(t, err, nil)
		}
		pars = append(pars, *par)
	}

	if pars[0].Id == 0 {
		err = errors.New("invalid value ID")
		assert.NoError(t, err, nil)
	}
}
