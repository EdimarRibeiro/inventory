package tests

import (
	"errors"
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidTenantCreateDataBase(t *testing.T) {
	database.Initialize(false)
	tenRepo := &database.TenantRepository{DB: database.DB}

	tens, err := tenRepo.Search("Id >= 1")
	if err != nil || len(tens) == 0 {
		ten, err := entities.NewTenant("Teste", "09066936754", 0)
		assert.NoError(t, err, nil)
		err = ten.Validate()
		assert.NoError(t, err, nil)
		ten, err = tenRepo.Save(ten)
		assert.NoError(t, err, nil)
		if ten.Id == 0 {
			err = errors.New("invalid value ID")
			assert.NoError(t, err, nil)
		}
		tens = append(tens, *ten)
	}

	if tens[0].Id == 0 {
		err = errors.New("invalid value ID")
		assert.NoError(t, err, nil)
	}
}
