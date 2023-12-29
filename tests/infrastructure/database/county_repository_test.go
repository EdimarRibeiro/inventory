package tests

import (
	"errors"
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidCountryCreateDataBase(t *testing.T) {
	database.Initialize(false)
	countryrep := &database.CountryRepository{DB: database.DB}

	countryId, err := countryrep.GetCountryId("00001")
	if err != nil || countryId == 0 {
		country, err := countryrep.Save(&entities.Country{CountryCode: "00001", Description: "test"})
		assert.NoError(t, err, nil)
		if country.Id == 0 {
			err = errors.New("invalid value ID")
			assert.NoError(t, err, nil)
		}
		countryId = country.Id
	}
	if countryId == 0 {
		err = errors.New("invalid value ID")
		assert.NoError(t, err, nil)
	}
}
