package tests

import (
	"errors"
	"testing"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"

	"github.com/stretchr/testify/assert"
)

func TestValidDocumentItemCreate(t *testing.T) {
	model := entities.DocumentItem{DocumentId: 10, InventoryId: 10, Sequency: "000", ProductId: 10}
	assert.NoError(t, model.Validate(), nil)
}

func TestValidDocumentItemNew(t *testing.T) {
	model, err := entities.NewDocumentItem(10, 10, "000", 10, "", 0, "", 0, 0, "", "", "", "", 0, 0, 0, 0, 0, 0, "", "", "", 0, 0, 0, "", 0, 0, 0, 0, 0, "", 0, 0, 0, 0, 0, "", 0)
	assert.NoError(t, err, nil)
	model.Validate()
}

func TestValidCreateDocumentItem(t *testing.T) {
	model, err := entities.CreateDocumentItem(10, 10, "|C170|1|001010||15|FD|801,94|0|0|010|1403|203|0|0|0|0|0|0|0|||0|0|0|50|605,58|1,86|0,00|0,00|11,26|50|605,58|8,54|0,00|0,00|51,72|||")
	assert.NoError(t, err, nil)

	if model.DocumentId != 10 {
		assert.NoError(t, errors.New("invalid value field DocumentId"))
	}
	if model.InventoryId != 10 {
		assert.NoError(t, errors.New("invalid value field InventoryId"))
	}
	if model.Sequency != "1" {
		assert.NoError(t, errors.New("invalid value field Sequency"))
	}
	if model.ProductId != 10 {
		assert.NoError(t, errors.New("invalid value field ProductId"))
	}
	if model.Complememt != "" {
		assert.NoError(t, errors.New("invalid value field Complememt"))
	}
	if model.Quantity != 15 {
		assert.NoError(t, errors.New("invalid value field Quantity"))
	}
	if model.UnitId != "FD" {
		assert.NoError(t, errors.New("invalid value field UnitId"))
	}
	if model.Value != 801.94 {
		assert.NoError(t, errors.New("invalid value field Value"))
	}
	if model.Discount != 0 {
		assert.NoError(t, errors.New("invalid value field Discount"))
	}
	if model.MovimentType != "0" {
		assert.NoError(t, errors.New("invalid value field MovimentType"))
	}
	if model.CstCode != "010" {
		assert.NoError(t, errors.New("invalid value field CstCode"))
	}
	if model.CfopCode != "1403" {
		assert.NoError(t, errors.New("invalid value field CfopCode"))
	}
	if model.OperationNatureId != "203" {
		assert.NoError(t, errors.New("invalid value field OperationNatureId"))
	}
	if model.BaseIcms != 0 {
		assert.NoError(t, errors.New("invalid value field BaseIcms"))
	}
	if model.AliquotIcms != 0 {
		assert.NoError(t, errors.New("invalid value field AliquotIcms"))
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
	if model.AliquotIcmsSt != 0 {
		assert.NoError(t, errors.New("invalid value field AliquotIcmsSt"))
	}
	if model.ApurationIpiCode != "0" {
		assert.NoError(t, errors.New("invalid value field ApurationIpiCode"))
	}
	if model.CstIpiCode != "" {
		assert.NoError(t, errors.New("invalid value field CstIpiCode"))
	}
	if model.LegalIpiCode != "" {
		assert.NoError(t, errors.New("invalid value field LegalIpiCode"))
	}
	if model.BaseIpi != 0 {
		assert.NoError(t, errors.New("invalid value field BaseIpi"))
	}
	if model.AliquotIpi != 0 {
		assert.NoError(t, errors.New("invalid value field AliquotIpi"))
	}
	if model.ValueIpi != 0 {
		assert.NoError(t, errors.New("invalid value field ValueIpi"))
	}
	if model.CstPisCode != "50" {
		assert.NoError(t, errors.New("invalid value field CstPisCode"))
	}
	if model.BasePis != 605.58 {
		assert.NoError(t, errors.New("invalid value field BasePis"))
	}
	if model.AliquotPis != 1.86 {
		assert.NoError(t, errors.New("invalid value field AliquotPis"))
	}
	if model.QuantityBasePis != 0 {
		assert.NoError(t, errors.New("invalid value field QuantityBasePis"))
	}
	if model.ValueAliquotPis != 0 {
		assert.NoError(t, errors.New("invalid value field ValueAliquotPis"))
	}
	if model.ValuePis != 11.26 {
		assert.NoError(t, errors.New("invalid value field ValuePis"))
	}
	if model.CstCofinsCode != "50" {
		assert.NoError(t, errors.New("invalid value field CstCofinsCode"))
	}
	if model.BaseCofins != 605.58 {
		assert.NoError(t, errors.New("invalid value field BaseCofins"))
	}
	if model.AliquotCofins != 8.54 {
		assert.NoError(t, errors.New("invalid value field AliquotCofins"))
	}
	if model.QuantityBaseCofins != 0 {
		assert.NoError(t, errors.New("invalid value field QuantityBaseCofins"))
	}
	if model.ValueAliquotCofins != 0 {
		assert.NoError(t, errors.New("invalid value field ValueAliquotCofins"))
	}
	if model.ValueCofins != 51.72 {
		assert.NoError(t, errors.New("invalid value field ValueCofins"))
	}
	if model.AccountingCode != "" {
		assert.NoError(t, errors.New("invalid value field AccountingCode"))
	}
	if model.RebateValue != 0 {
		assert.NoError(t, errors.New("invalid value field RebateValue"))
	}
}
