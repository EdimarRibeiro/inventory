package tests

import (
	"errors"
	"testing"
	"time"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidDocumentCreate(t *testing.T) {
	model := entities.Document{Id: 0, InventoryId: 10}
	assert.NoError(t, model.Validate(), nil)
}

func TestValidDocumentNew(t *testing.T) {
	model, err := entities.NewDocument(10, 0, "", "", "", "", "", "", "", time.Now(), time.Now(), 0, "", 0, 0, 0, "", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, false)
	assert.NoError(t, err, nil)
	model.Validate()
}

func TestValidCreateDocument(t *testing.T) {
	model, err := entities.CreateDocument(10, "|C100|0|1|222108|55|00|010|473710|32221200074569001840550100004737101450070258|29122022|02012023|1682,23|1|0||1682,23|9|0|0|10,03|0,00|0,00|0,00|0,00|0,00|25,78|118,60|00|00|")
	assert.NoError(t, err, nil)
	if model.Id != 0 {
		assert.NoError(t, errors.New("invalid value field Id"))
	}
	if model.InventoryId != 10 {
		assert.NoError(t, errors.New("invalid value field InventoryId"))
	}
	if model.ParticipantId != 0 {
		assert.NoError(t, errors.New("invalid value field ParticipantId"))
	}
	if model.OperationId != "0" {
		assert.NoError(t, errors.New("invalid value field OperationId "+model.OperationId))
	}
	if model.EmitentTypeId != "1" {
		assert.NoError(t, errors.New("invalid value field EmitentTypeId "+model.EmitentTypeId))
	}
	if model.ModelId != "55" {
		assert.NoError(t, errors.New("invalid value field ModelId "+model.ModelId))
	}
	if model.StatusId != "00" {
		assert.NoError(t, errors.New("invalid value field StatusId "+model.StatusId))
	}
	if model.Serie != "010" {
		assert.NoError(t, errors.New("invalid value field Serie "+model.Serie))
	}
	if model.Number != "473710" {
		assert.NoError(t, errors.New("invalid value field Number "+model.Number))
	}
	if model.DocumentKey != "32221200074569001840550100004737101450070258" {
		assert.NoError(t, errors.New("invalid value field DocumentKey "+model.DocumentKey))
	}
	dateEmi, err := time.Parse("02-01-2006", "29-12-2022")
	assert.NoError(t, err)
	if model.EmitentDate != dateEmi {
		assert.NoError(t, errors.New("invalid value field EmitentDate"))
	}
	dateExit, err := time.Parse("02-01-2006", "02-01-2023")
	assert.NoError(t, err)
	if model.ExitDate != dateExit {
		assert.NoError(t, errors.New("invalid value field ExitDate"))
	}
	if model.DocumentValue != 1682.23 {
		assert.NoError(t, errors.New("invalid value field DocumentValue"))
	}
	if model.PayTypeId != "1" {
		assert.NoError(t, errors.New("invalid value field PayTypeId "+model.PayTypeId))
	}
	if model.Discount != 0 {
		assert.NoError(t, errors.New("invalid value field Discount"))
	}
	if model.Reduction != 0 {
		assert.NoError(t, errors.New("invalid value field Reduction"))
	}
	if model.ProductValue != 1682.23 {
		assert.NoError(t, errors.New("invalid value field ProductValue"))
	}
	if model.FreightType != "9" {
		assert.NoError(t, errors.New("invalid value field FreightType "+model.FreightType))
	}
	if model.FreightValue != 0 {
		assert.NoError(t, errors.New("invalid value field FreightValue"))
	}
	if model.SafeValue != 0 {
		assert.NoError(t, errors.New("invalid value field SafeValue"))
	}
	if model.ExpenseValue != 10.03 {
		assert.NoError(t, errors.New("invalid value field ExpenseValue"))
	}
	if model.BaseIcms != 0 {
		assert.NoError(t, errors.New("invalid value field BaseIcms"))
	}
	if model.ValueIcms != 0 {
		assert.NoError(t, errors.New("invalid value field ValueIcms"))
	}
	if model.BaseIcmsSt != 0 {
		assert.NoError(t, errors.New("invalid value field BaseIcmsSt"))
	}
	if model.ValueIcmsSt != 0 {
		assert.NoError(t, errors.New("invalid value field ValueIcmsSt"))
	}
	if model.ValueIpi != 0 {
		assert.NoError(t, errors.New("invalid value field ValueIpi"))
	}
	if model.ValuePis != 25.78 {
		assert.NoError(t, errors.New("invalid value field ValuePis"))
	}
	if model.ValueCofins != 118.60 {
		assert.NoError(t, errors.New("invalid value field ValueCofins"))
	}
	if model.ValuePisSt != 0 {
		assert.NoError(t, errors.New("invalid value field ValuePisSt"))
	}
	if model.ValueCofinsSt != 0 {
		assert.NoError(t, errors.New("invalid value field ValueCofinsSt"))
	}
	if model.Processed != true {
		assert.NoError(t, errors.New("invalid value field Processed"))
	}
	if model.Cloused != false {
		assert.NoError(t, errors.New("invalid value field Cloused"))
	}
	if model.Imported != false {
		assert.NoError(t, errors.New("invalid value field Imported"))
	}
}

func TestValidDocumentCreateError(t *testing.T) {
	model := entities.Document{Id: 0}
	assert.Error(t, model.Validate(), nil)
}
