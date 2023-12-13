package tests

import (
	"errors"
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidProductCreate(t *testing.T) {
	model := entities.Product{Id: 0, OriginCode: "100", TenantId: 10, Description: "Sabonete", BarCode: "00300600911", UnitId: "UN", Type: "00", GenderCode: "00"}
	assert.NoError(t, model.Validate(), nil)
}

func TestValidCreateProduct(t *testing.T) {
	model, err := entities.CreateProduct(10, "|0200|001010|COCA COLA PET 600ML C/12|7894900011609||FD|00|22021000||22||0|0300700|")
	assert.NoError(t, err, nil)

	if model.Id != 0 {
		assert.NoError(t, errors.New("invalid value field Id"))
	}
	if model.TenantId != 10 {
		assert.NoError(t, errors.New("invalid value field TenantId"))
	}
	if model.OriginCode != "001010" {
		assert.NoError(t, errors.New("invalid value field OriginCode"))
	}
	if model.Description != "COCA COLA PET 600ML C/12" {
		assert.NoError(t, errors.New("invalid value field Description"))
	}
	if model.BarCode != "7894900011609" {
		assert.NoError(t, errors.New("invalid value field BarCode"))
	}
	if model.OldOriginCode != "" {
		assert.NoError(t, errors.New("invalid value field OldOriginCode"))
	}
	if model.UnitId != "FD" {
		assert.NoError(t, errors.New("invalid value field UnitId"))
	}
	if model.Type != "00" {
		assert.NoError(t, errors.New("invalid value field Type"))
	}
	if model.NcmCode != "22021000" {
		assert.NoError(t, errors.New("invalid value field NcmCode"))
	}
	if model.ExIpi != "" {
		assert.NoError(t, errors.New("invalid value field ExIpi"))
	}
	if model.GenderCode != "22" {
		assert.NoError(t, errors.New("invalid value field GenderCode"))
	}
	if model.ServiceCode != "" {
		assert.NoError(t, errors.New("invalid value field ServiceCode"))
	}
	if model.AliqIcms != 0 {
		assert.NoError(t, errors.New("invalid value field AliqIcms"))
	}
	if model.CestCode != "0300700" {
		assert.NoError(t, errors.New("invalid value field CestCode"))
	}
}

func TestValidProductCreateError(t *testing.T) {
	model := entities.Product{Id: 0, TenantId: 10}
	assert.Error(t, model.Validate(), nil)
}
