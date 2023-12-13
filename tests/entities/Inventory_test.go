package tests

import (
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidInventoryCreate(t *testing.T) {
	model := entities.Inventory{Id: 0, TenantId: 10, Name: "Jan-2023"}
	assert.NoError(t, model.Validate(), nil)
}
func TestValidInventoryCreateError(t *testing.T) {
	model := entities.Inventory{Id: 0, TenantId: 10}
	assert.Error(t, model.Validate(), nil)
}
