package tests

import (
	"errors"
	"testing"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidCityCreateDataBase(t *testing.T) {
	database.Initialize(false)
	cityrep := &database.CityRepository{DB: database.DB}

	cityId, err := cityrep.GetCityId("0000001")
	if err != nil || cityId == 0 {
		city, err := cityrep.Save(&entities.City{CityCode: "0000001", Description: "test"})
		assert.NoError(t, err, nil)
		if city.Id == 0 {
			err = errors.New("invalid value ID")
			assert.NoError(t, err, nil)
		}
		cityId = city.Id
	}
	if cityId == 0 {
		err = errors.New("invalid value ID")
		assert.NoError(t, err, nil)
	}
}
