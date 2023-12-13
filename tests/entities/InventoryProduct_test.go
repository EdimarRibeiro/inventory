package tests

import (
	"errors"
	"testing"
	"time"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidInventoryProductCreate(t *testing.T) {
	model := entities.InventoryProduct{ProductId: 10, InventoryId: 10, OriginCode: "100", UnitId: "UN", Quantity: 0, Value: 0}
	assert.NoError(t, model.Validate(), nil)
}

func TestValidCreateInventoryProduct(t *testing.T) {
	model, err := entities.CreateInventoryProduct(10, time.Now(), "|H010|100|UN|15,001|0,006015|7,01|0|111||||||")
	assert.NoError(t, err, nil)

	if model.InventoryId != 10 {
		assert.NoError(t, errors.New("invalid value field InventoryId"))
	}
	if model.ProductId != 10 {
		assert.NoError(t, errors.New("invalid value field DocumentId"))
	}
	if model.UnitId != "UN" {
		assert.NoError(t, errors.New("invalid value field UnitId"))
	}
	if model.Quantity != 15.001 {
		assert.NoError(t, errors.New("invalid value field Quantity"))
	}
	if model.Value != 0.006015 {
		assert.NoError(t, errors.New("invalid value field Value"))
	}
	if model.ValueTotal != 7.01 {
		assert.NoError(t, errors.New("invalid value field ValueTotal"))
	}
	if model.PossessionCode != "0" {
		assert.NoError(t, errors.New("invalid value field PossessionCode"))
	}
	if model.ParticipantId != 0 {
		assert.NoError(t, errors.New("invalid value field ParticipantId"))
	}
	if model.Complement != "" {
		assert.NoError(t, errors.New("invalid value field Complement"))
	}
	if model.AccountingCode != "" {
		assert.NoError(t, errors.New("invalid value field AccountingCode"))
	}
	if model.ValueIr != 0 {
		assert.NoError(t, errors.New("invalid value field ValueIr"))
	}
}
func TestValidInventoryProductCreateError(t *testing.T) {
	model := entities.InventoryProduct{ProductId: 0, InventoryId: 10}
	assert.Error(t, model.Validate(), nil)
}
