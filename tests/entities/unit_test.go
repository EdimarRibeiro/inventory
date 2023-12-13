package tests

import (
	"errors"
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidUnitCreate(t *testing.T) {
	model := entities.Unit{Id: "UN", TenantId: 10, Description: "TEST"}
	assert.NoError(t, model.Validate(), nil)
}

func TestValidCreateUnit(t *testing.T) {
	model, err := entities.CreateUnit(10, "|0190|FD|UNIDADE em FD|")
	assert.NoError(t, err, nil)

	if model.Id != "FD" {
		assert.NoError(t, errors.New("invalid value field Id"))
	}
	if model.TenantId != 10 {
		assert.NoError(t, errors.New("invalid value field TenantId"))
	}
	if model.Description != "UNIDADE em FD" {
		assert.NoError(t, errors.New("invalid value field Description"))
	}
}

func TestValidUnitCreateError(t *testing.T) {
	model := entities.Unit{Id: "", TenantId: 10}
	assert.Error(t, model.Validate(), nil)
}
