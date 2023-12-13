package tests

import (
	"errors"
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidParticipantCreate(t *testing.T) {
	model := entities.Participant{Id: 0, OriginCode: "100", TenantId: 10, Name: "Joe", DocumentCpf: "00300600911", CountryId: 10, CityId: 10}
	assert.NoError(t, model.Validate(), nil)
}

func TestValidCreateParticipant(t *testing.T) {
	model, err := entities.CreateParticipant(10, "|0150|221865|COMERCIAL CEREALISTA PRETTI LTDA|1058|28534402000357||082077860|3201506||AV SILVIO AVIDOS,1192|0||SAO SILVANO|")
	assert.NoError(t, err, nil)

	if model.Id != 0 {
		assert.NoError(t, errors.New("invalid value field Id"))
	}
	if model.TenantId != 10 {
		assert.NoError(t, errors.New("invalid value field TenantId"))
	}
	if model.OriginCode != "221865" {
		assert.NoError(t, errors.New("invalid value field OriginCode"))
	}
	if model.Name != "COMERCIAL CEREALISTA PRETTI LTDA" {
		assert.NoError(t, errors.New("invalid value field Name"))
	}
	if model.Document != "28534402000357" {
		assert.NoError(t, errors.New("invalid value field Document"))
	}
	if model.DocumentCpf != "" {
		assert.NoError(t, errors.New("invalid value field DocumentCpf"))
	}
	if model.Registration != "082077860" {
		assert.NoError(t, errors.New("invalid value field Registration"))
	}
	if model.CountryId != 1058 {
		assert.NoError(t, errors.New("invalid value field CountryId"))
	}
	if model.CityId != 3201506 {
		assert.NoError(t, errors.New("invalid value field CityId"))
	}
	if model.Suframa != "" {
		assert.NoError(t, errors.New("invalid value field Suframa"))
	}
	if model.Street != "AV SILVIO AVIDOS,1192" {
		assert.NoError(t, errors.New("invalid value field Street"))
	}
	if model.Number != "0" {
		assert.NoError(t, errors.New("invalid value field Number"))
	}
	if model.Complememt != "" {
		assert.NoError(t, errors.New("invalid value field Complememt"))
	}
	if model.Neighborhood != "SAO SILVANO" {
		assert.NoError(t, errors.New("invalid value field Neighborhood"))
	}
}

func TestValidParticipantCreateError(t *testing.T) {
	model := entities.Participant{Id: 0, TenantId: 10}
	assert.Error(t, model.Validate(), nil)
}
