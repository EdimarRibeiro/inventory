package tests

import (
	"errors"
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidInventoryCreateDataBase(t *testing.T) {
	database.Initialize(false)
	invRepo := &database.InventoryRepository{DB: database.DB}
	tenRepo := &database.TenantRepository{DB: database.DB}
	perRepo := &database.PersonRepository{DB: database.DB}
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
}
