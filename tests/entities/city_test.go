package tests

import (
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidModelCreate(t *testing.T) {
	model := entities.City{Id: 0, CityCode: "1234567"}
	assert.NoError(t, model.Validate(), nil)
}
func TestValidModelCreateError(t *testing.T) {
	model := entities.City{Id: 0, CityCode: "123457"}
	assert.Error(t, model.Validate(), nil)
}
