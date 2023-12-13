package tests

import (
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidCountryCreate(t *testing.T) {
	model := entities.Country{Id: 0, CountryCode: "12345"}
	assert.NoError(t, model.Validate(), nil)
}
func TestValidCountryCreateError(t *testing.T) {
	model := entities.Country{Id: 0, CountryCode: "123"}
	assert.Error(t, model.Validate(), nil)
}
