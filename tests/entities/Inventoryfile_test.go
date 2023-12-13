package tests

import (
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidInventoryFileCreate(t *testing.T) {
	model := entities.InventoryFile{Id: 0, InventoryId: 10, FileName: "Jan-2023"}
	assert.NoError(t, model.Validate(), nil)
}
func TestValidInventoryFileCreateError(t *testing.T) {
	model := entities.InventoryFile{Id: 0, InventoryId: 10}
	assert.Error(t, model.Validate(), nil)
}
