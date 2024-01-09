package tests

import (
	"errors"
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidUnitConvertCreate(t *testing.T) {
	model := entities.UnitConvert{UnitId: "UN", ProductId: 10, TenantId: 10, ConversionFactor: 1, BarCode: "7896315125131"}
	assert.NoError(t, model.Validate(), nil)
}

func TestValidCreateUnitConvert(t *testing.T) {
	model, err := entities.CreateUnitConvert(10, 10, "|0220|FD|10,003006|7896315125131|")
	assert.NoError(t, err, nil)

	if model.UnitId != "FD" {
		assert.NoError(t, errors.New("invalid value field Id"))
	}
	if model.TenantId != 10 {
		assert.NoError(t, errors.New("invalid value field TenantId"))
	}
	if model.ProductId != 10 {
		assert.NoError(t, errors.New("invalid value field ProductId"))
	}
	if model.ConversionFactor != 10.003006 {
		assert.NoError(t, errors.New("invalid value field ConversionFactor"))
	}
	if model.BarCode != "7896315125131" {
		assert.NoError(t, errors.New("invalid value field BarCode"))
	}
}

func TestValidUnitConvertCreateError(t *testing.T) {
	model := entities.UnitConvert{UnitId: "", TenantId: 10}
	assert.Error(t, model.Validate(), nil)
}
