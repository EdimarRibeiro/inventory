package entities

import (
	"errors"

	"github.com/EdimarRibeiro/inventory/internal/utils"
)

/*C170*/
type DocumentItem struct {
	DocumentId         uint64    `gorm:"primaryKey" json:"documentId"`
	Document           Document  `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	Sequency           string    `gorm:"primaryKey;size:3" json:"sequency"`
	InventoryId        uint64    `json:"inventoryId"`
	Inventory          Inventory `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	ProductId          uint64    `gorm:"primaryKey;size:3" json:"productId"`
	Product            Product   `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"product"`
	Complememt         string    `gorm:"size:255" json:"complememt"`
	Quantity           float64   `gorm:"type:decimal (18,5)" json:"quantity"`
	UnitId             string    `gorm:"size:6" json:"unitId"`
	Unit               Unit      `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	Value              float64   `gorm:"type:decimal (18,2)" json:"value"`
	Discount           float64   `gorm:"type:decimal (12,2)" json:"discount"`
	MovimentType       string    `gorm:"size:1" json:"movimentType"`
	CstCode            string    `gorm:"size:3" json:"cstCode"`
	CfopCode           string    `gorm:"size:4" json:"cfopCode"`
	OperationNatureId  string    `gorm:"size:10" json:"operationNatureId"`
	BaseIcms           float64   `gorm:"type:decimal (18,2)" json:"baseIcms"`
	AliquotIcms        float64   `gorm:"type:decimal (12,2)" json:"aliquotIcms"`
	ValueIcms          float64   `gorm:"type:decimal (18,2)" json:"valueIcms"`
	BaseIcmsSt         float64   `gorm:"type:decimal (18,2)" json:"baseIcmsSt"`
	ValueIcmsSt        float64   `gorm:"type:decimal (12,2)" json:"valueIcmsSt"`
	AliquotIcmsSt      float64   `gorm:"type:decimal (12,2)" json:"aliquotIcmsSt"`
	ApurationIpiCode   string    `gorm:"size:1" json:"apurationIpiCode"`
	CstIpiCode         string    `gorm:"size:2" json:"cstIpiCode"`
	LegalIpiCode       string    `gorm:"size:3" json:"legalIpiCode"`
	BaseIpi            float64   `gorm:"type:decimal (18,2)" json:"baseIpi"`
	AliquotIpi         float64   `gorm:"type:decimal (12,2)" json:"aliquotIpi"`
	ValueIpi           float64   `gorm:"type:decimal (12,2)" json:"valueIpi"`
	CstPisCode         string    `gorm:"size:2" json:"cstPisCode"`
	BasePis            float64   `gorm:"type:decimal (18,2)" json:"basePis"`
	AliquotPis         float64   `gorm:"type:decimal (12,4)" json:"aliquotPis"`
	QuantityBasePis    float64   `gorm:"type:decimal (12,3)" json:"quantityBasePis"`
	ValueAliquotPis    float64   `gorm:"type:decimal (12,4)" json:"valueAliquotPis"`
	ValuePis           float64   `gorm:"type:decimal (12,2)" json:"valuePis"`
	CstCofinsCode      string    `gorm:"size:2" json:"cstCofinsCode"`
	BaseCofins         float64   `gorm:"type:decimal (18,2)" json:"baseCofins"`
	AliquotCofins      float64   `gorm:"type:decimal (12,4)" json:"aliquotCofins"`
	QuantityBaseCofins float64   `gorm:"type:decimal (12,3)" json:"quantityBaseCofins"`
	ValueAliquotCofins float64   `gorm:"type:decimal (12,4)" json:"valueAliquotCofins"`
	ValueCofins        float64   `gorm:"type:decimal (12,2)" json:"valueCofins"`
	AccountingCode     string    `json:"accountingCode"`
	RebateValue        float64   `gorm:"type:decimal (12,2)" json:"rebateValue"`
}

func (c *DocumentItem) Validate() error {
	if c.Sequency == "" {
		return errors.New("the sequency is required")
	}
	if c.DocumentId == 0 {
		return errors.New("the documentId is invalid")
	}
	if c.InventoryId == 0 {
		return errors.New("the inventoryId is invalid")
	}
	if c.ProductId == 0 {
		return errors.New("the productId is invalid")
	}
	return nil
}

func NewDocumentItem(documentId uint64, inventoryId uint64, sequency string, productId uint64, complememt string, quantity float64, unitId string, value float64, discount float64, movimentType string, cstCode string, cfopCode string, operationNatureId string, baseIcms float64, aliquotIcms float64, valueIcms float64, baseIcmsSt float64, valueIcmsSt float64, aliquotIcmsSt float64, apurationIpiCode string, cstIpiCode string, legalIpiCode string, baseIpi float64, aliquotIpi float64, valueIpi float64, cstPisCode string, basePis float64, aliquotPis float64, quantityBasePis float64, valueAliquotPis float64, valuePis float64, cstCofinsCode string, baseCofins float64, aliquotCofins float64, quantityBaseCofins float64, valueAliquotCofins float64, valueCofins float64, accountingCode string, rebateValue float64) (*DocumentItem, error) {
	model := &DocumentItem{
		DocumentId:         documentId,
		InventoryId:        inventoryId,
		Sequency:           sequency,
		ProductId:          productId,
		Complememt:         complememt,
		Quantity:           quantity,
		UnitId:             unitId,
		Value:              value,
		Discount:           discount,
		MovimentType:       movimentType,
		CstCode:            cstCode,
		CfopCode:           cfopCode,
		OperationNatureId:  operationNatureId,
		BaseIcms:           baseIcms,
		AliquotIcms:        aliquotIcms,
		ValueIcms:          valueIcms,
		BaseIcmsSt:         baseIcmsSt,
		ValueIcmsSt:        valueIcmsSt,
		AliquotIcmsSt:      aliquotIcmsSt,
		ApurationIpiCode:   apurationIpiCode,
		CstIpiCode:         cstIpiCode,
		LegalIpiCode:       legalIpiCode,
		BaseIpi:            baseIpi,
		AliquotIpi:         aliquotIpi,
		ValueIpi:           valueIpi,
		CstPisCode:         cstPisCode,
		BasePis:            basePis,
		AliquotPis:         aliquotPis,
		QuantityBasePis:    quantityBasePis,
		ValueAliquotPis:    valueAliquotPis,
		ValuePis:           valuePis,
		CstCofinsCode:      cstCofinsCode,
		BaseCofins:         baseCofins,
		AliquotCofins:      aliquotCofins,
		QuantityBaseCofins: quantityBaseCofins,
		ValueAliquotCofins: valueAliquotCofins,
		ValueCofins:        valueCofins,
		AccountingCode:     accountingCode,
		RebateValue:        rebateValue,
	}

	err := model.Validate()
	if err != nil {
		return nil, err
	}
	return model, nil
}

func NewDocumentItemEntity(entity DocumentItem) (*DocumentItem, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func CreateDocumentItem(documentId uint64, inventoryId uint64, productId uint64, line string) (*DocumentItem, error) {
	var err error = nil
	documentItem := DocumentItem{}
	documentItem.DocumentId, err = documentId, nil
	documentItem.InventoryId, err = inventoryId, nil
	documentItem.ProductId, err = productId, nil
	documentItem.Sequency, err = utils.CopyText(line, 2)
	if err != nil {
		return nil, err
	}

	documentItem.Complememt, err = utils.CopyText(line, 4)
	if err != nil {
		return nil, err
	}
	documentItem.Quantity, err = utils.CopyTextFloat(line, 5, 5)
	if err != nil {
		return nil, err
	}
	documentItem.UnitId, err = utils.CopyText(line, 6)
	if err != nil {
		return nil, err
	}
	documentItem.Value, err = utils.CopyTextFloat(line, 7, 2)
	if err != nil {
		return nil, err
	}
	documentItem.Discount, err = utils.CopyTextFloat(line, 8, 2)
	if err != nil {
		return nil, err
	}
	documentItem.MovimentType, err = utils.CopyText(line, 9)
	if err != nil {
		return nil, err
	}
	documentItem.CstCode, err = utils.CopyText(line, 10)
	if err != nil {
		return nil, err
	}
	documentItem.CfopCode, err = utils.CopyText(line, 11)
	if err != nil {
		return nil, err
	}
	documentItem.OperationNatureId, err = utils.CopyText(line, 12)
	if err != nil {
		return nil, err
	}
	documentItem.BaseIcms, err = utils.CopyTextFloat(line, 13, 2)
	if err != nil {
		return nil, err
	}
	documentItem.AliquotIcms, err = utils.CopyTextFloat(line, 14, 2)
	if err != nil {
		return nil, err
	}
	documentItem.ValueIcms, err = utils.CopyTextFloat(line, 15, 2)
	if err != nil {
		return nil, err
	}
	documentItem.BaseIcmsSt, err = utils.CopyTextFloat(line, 16, 2)
	if err != nil {
		return nil, err
	}
	documentItem.ValueIcmsSt, err = utils.CopyTextFloat(line, 17, 2)
	if err != nil {
		return nil, err
	}
	documentItem.AliquotIcmsSt, err = utils.CopyTextFloat(line, 18, 2)
	if err != nil {
		return nil, err
	}
	documentItem.ApurationIpiCode, err = utils.CopyText(line, 19)
	if err != nil {
		return nil, err
	}
	documentItem.CstIpiCode, err = utils.CopyText(line, 20)
	if err != nil {
		return nil, err
	}
	documentItem.LegalIpiCode, err = utils.CopyText(line, 21)
	if err != nil {
		return nil, err
	}
	documentItem.BaseIpi, err = utils.CopyTextFloat(line, 22, 2)
	if err != nil {
		return nil, err
	}
	documentItem.AliquotIpi, err = utils.CopyTextFloat(line, 23, 2)
	if err != nil {
		return nil, err
	}
	documentItem.ValueIpi, err = utils.CopyTextFloat(line, 24, 2)
	if err != nil {
		return nil, err
	}
	documentItem.CstPisCode, err = utils.CopyText(line, 25)
	if err != nil {
		return nil, err
	}
	documentItem.BasePis, err = utils.CopyTextFloat(line, 26, 2)
	if err != nil {
		return nil, err
	}
	documentItem.AliquotPis, err = utils.CopyTextFloat(line, 27, 4)
	if err != nil {
		return nil, err
	}
	documentItem.QuantityBasePis, err = utils.CopyTextFloat(line, 28, 3)
	if err != nil {
		return nil, err
	}
	documentItem.ValueAliquotPis, err = utils.CopyTextFloat(line, 29, 4)
	if err != nil {
		return nil, err
	}
	documentItem.ValuePis, err = utils.CopyTextFloat(line, 30, 2)
	if err != nil {
		return nil, err
	}
	documentItem.CstCofinsCode, err = utils.CopyText(line, 31)
	if err != nil {
		return nil, err
	}
	documentItem.BaseCofins, err = utils.CopyTextFloat(line, 32, 2)
	if err != nil {
		return nil, err
	}
	documentItem.AliquotCofins, err = utils.CopyTextFloat(line, 33, 4)
	if err != nil {
		return nil, err
	}
	documentItem.QuantityBaseCofins, err = utils.CopyTextFloat(line, 34, 3)
	if err != nil {
		return nil, err
	}
	documentItem.ValueAliquotCofins, err = utils.CopyTextFloat(line, 35, 4)
	if err != nil {
		return nil, err
	}
	documentItem.ValueCofins, err = utils.CopyTextFloat(line, 36, 2)
	if err != nil {
		return nil, err
	}
	documentItem.AccountingCode, err = utils.CopyText(line, 37)
	if err != nil {
		return nil, err
	}
	documentItem.RebateValue, err = utils.CopyTextFloat(line, 38, 2)
	if err != nil {
		return nil, err
	}
	return NewDocumentItemEntity(documentItem)
}
