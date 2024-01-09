package tests

import (
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidUserCreate(t *testing.T) {
	model := entities.User{Id: 0, Name: "Conesoft", Login: "@", Password: "00100600911", TenantId: 10}
	assert.NoError(t, model.Validate(), nil)
}
func TestValidUserCreateError(t *testing.T) {
	model := entities.User{Id: 0}
	assert.Error(t, model.Validate(), nil)
}
