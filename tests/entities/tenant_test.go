package tests

import (
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidTenantCreate(t *testing.T) {
	model := entities.Tenant{Id: 0, Name: "Conesoft", Document: "00100600911", PersonId: 10}
	assert.NoError(t, model.Validate(), nil)
}
func TestValidTenantCreateError(t *testing.T) {
	model := entities.Tenant{Id: 0}
	assert.Error(t, model.Validate(), nil)
}
