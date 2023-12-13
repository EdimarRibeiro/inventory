package tests

import (
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidPersonCreate(t *testing.T) {
	model := entities.Person{Id: 0, TenantId: 10, Name: "Jan-2023", Document: "02005008000199", CountryId: 10, CityId: 10}
	assert.NoError(t, model.Validate(), nil)
}
func TestValidPersonCreateError(t *testing.T) {
	model := entities.Person{Id: 0, TenantId: 10}
	assert.Error(t, model.Validate(), nil)
}
