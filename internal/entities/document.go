package entities

import (
	"errors"
	"time"

	"github.com/EdimarRibeiro/inventory/internal/utils"
)

/*C100*/
type Document struct {
	Id            uint64      `gorm:"primaryKey;autoIncrement:true"  json:"id"`
	InventoryId   uint64      `gorm:"index:idx_Inventory"  json:"inventoryId"`
	Inventory     Inventory   `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	ParticipantId uint64      `json:"participantId"`
	Participant   Participant `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"participant"`
	OperationId   string      `gorm:"size:1" json:"operationId"`
	EmitentTypeId string      `gorm:"size:1" json:"emitentTypeId"`
	ModelId       string      `gorm:"size:2" json:"modelId"`
	StatusId      string      `gorm:"size:2" json:"statusId"`
	Serie         string      `gorm:"size:3" json:"serie"`
	Number        string      `gorm:"size:9" json:"number"`
	DocumentKey   string      `gorm:"size:44" json:"documentKey"`
	EmitentDate   time.Time   `gorm:"datetime" json:"emitentDate"`
	ExitDate      time.Time   `gorm:"datetime" json:"exitDate"`
	DocumentValue float64     `gorm:"type:decimal (18,2)" json:"documentValue"`
	PayTypeId     string      `gorm:"size:1" json:"payTypeId"`
	Discount      float64     `gorm:"type:decimal (12,2)" json:"discount"`
	Reduction     float64     `gorm:"type:decimal (12,2)" json:"reduction"`
	ProductValue  float64     `gorm:"type:decimal (18,2)" json:"productValue"`
	FreightType   string      `gorm:"size:1" json:"freightType"`
	FreightValue  float64     `gorm:"type:decimal (12,2)" json:"freightValue"`
	SafeValue     float64     `gorm:"type:decimal (12,2)" json:"safeValue"`
	ExpenseValue  float64     `gorm:"type:decimal (12,2)" json:"expenseValue"`
	BaseIcms      float64     `gorm:"type:decimal (18,2)" json:"baseIcms"`
	ValueIcms     float64     `gorm:"type:decimal (18,2)" json:"valueIcms"`
	BaseIcmsSt    float64     `gorm:"type:decimal (18,2)" json:"baseIcmsSt"`
	ValueIcmsSt   float64     `gorm:"type:decimal (18,2)" json:"valueIcmsSt"`
	ValueIpi      float64     `gorm:"type:decimal (18,2)" json:"valueIpi"`
	ValuePis      float64     `gorm:"type:decimal (18,2)" json:"valuePis"`
	ValueCofins   float64     `gorm:"type:decimal (18,2)" json:"valueCofins"`
	ValuePisSt    float64     `gorm:"type:decimal (18,2)" json:"valuePisSt"`
	ValueCofinsSt float64     `gorm:"type:decimal (18,2)" json:"valueCofinsSt"`
	Origined      string      `gorm:"size:3" json:"origined"`
	Processed     bool        `json:"processed"`
	Imported      bool        `json:"imported"`
	Cloused       bool        `json:"cloused"`
}

func (d *Document) Validate() error {
	if d.InventoryId == 0 {
		return errors.New("inventoryId id")
	}
	return nil
}

func (c *Document) SetProcessed() error {
	c.Processed = true
	return nil
}

func (c *Document) SetImported() error {
	c.Imported = true
	return nil
}

func (c *Document) SetCloused() error {
	c.Cloused = true
	return nil
}

func CreateDocument(inventoryId uint64, participantId uint64, line string) (*Document, error) {
	var err error = nil
	document := Document{}

	document.Id, err = 0, nil
	document.InventoryId, err = inventoryId, nil
	document.ParticipantId, err = participantId, nil

	document.OperationId, err = utils.CopyText(line, 2)
	if err != nil {
		return nil, err
	}
	document.EmitentTypeId, err = utils.CopyText(line, 3)
	if err != nil {
		return nil, err
	}
	document.ModelId, err = utils.CopyText(line, 5)
	if err != nil {
		return nil, err
	}
	document.StatusId, err = utils.CopyText(line, 6)
	if err != nil {
		return nil, err
	}
	document.Serie, err = utils.CopyText(line, 7)
	if err != nil {
		return nil, err
	}
	document.Number, err = utils.CopyText(line, 8)
	if err != nil {
		return nil, err
	}
	document.DocumentKey, err = utils.CopyText(line, 9)
	if err != nil {
		return nil, err
	}
	document.EmitentDate, err = utils.CopyTextDate(line, 10, "##-##-####")
	if err != nil {
		return nil, err
	}
	document.ExitDate, err = utils.CopyTextDate(line, 11, "##-##-####")
	if err != nil {
		return nil, err
	}
	document.DocumentValue, err = utils.CopyTextFloat(line, 12, 2)
	if err != nil {
		return nil, err
	}
	document.PayTypeId, err = utils.CopyText(line, 13)
	if err != nil {
		return nil, err
	}
	document.Discount, err = utils.CopyTextFloat(line, 14, 2)
	if err != nil {
		return nil, err
	}
	document.Reduction, err = utils.CopyTextFloat(line, 15, 2)
	if err != nil {
		return nil, err
	}
	document.ProductValue, err = utils.CopyTextFloat(line, 16, 2)
	if err != nil {
		return nil, err
	}
	document.FreightType, err = utils.CopyText(line, 17)
	if err != nil {
		return nil, err
	}
	document.FreightValue, err = utils.CopyTextFloat(line, 18, 2)
	if err != nil {
		return nil, err
	}
	document.SafeValue, err = utils.CopyTextFloat(line, 19, 2)
	if err != nil {
		return nil, err
	}
	document.ExpenseValue, err = utils.CopyTextFloat(line, 20, 2)
	if err != nil {
		return nil, err
	}
	document.BaseIcms, err = utils.CopyTextFloat(line, 21, 2)
	if err != nil {
		return nil, err
	}
	document.ValueIcms, err = utils.CopyTextFloat(line, 22, 2)
	if err != nil {
		return nil, err
	}
	document.BaseIcmsSt, err = utils.CopyTextFloat(line, 23, 2)
	if err != nil {
		return nil, err
	}
	document.ValueIcmsSt, err = utils.CopyTextFloat(line, 24, 2)
	if err != nil {
		return nil, err
	}
	document.ValueIpi, err = utils.CopyTextFloat(line, 25, 2)
	if err != nil {
		return nil, err
	}
	document.ValuePis, err = utils.CopyTextFloat(line, 26, 2)
	if err != nil {
		return nil, err
	}
	document.ValueCofins, err = utils.CopyTextFloat(line, 27, 2)
	if err != nil {
		return nil, err
	}
	document.ValuePisSt, err = utils.CopyTextFloat(line, 28, 2)
	if err != nil {
		return nil, err
	}
	document.ValueCofinsSt, err = utils.CopyTextFloat(line, 29, 2)
	if err != nil {
		return nil, err
	}
	document.Origined, err = "txt", nil
	document.Processed, err = true, nil
	document.Imported, err = false, nil
	document.Cloused, err = false, nil
	return NewDocumentEntity(document)
}

func NewDocument(inventoryId uint64, participantId uint64, operationId string, emitentTypeId string, modelId string, statusId string, serie string, number string, documentKey string, emitentDate time.Time, exitDate time.Time, documentValue float64, payTypeId string, discount float64, reduction float64, productValue float64, freightType string, freightValue float64, safeValue float64, expenseValue float64, baseIcms float64, valueIcms float64, baseIcmsSt float64, valueIcmsSt float64, valueIpi float64, valuePis float64, valueCofins float64, valuePisSt float64, valueCofinsSt float64, imported bool, origined string) (*Document, error) {
	model := Document{
		Id:            0,
		InventoryId:   inventoryId,
		ParticipantId: participantId,
		OperationId:   operationId,
		EmitentTypeId: emitentTypeId,
		ModelId:       modelId,
		StatusId:      statusId,
		Serie:         serie,
		Number:        number,
		DocumentKey:   documentKey,
		EmitentDate:   emitentDate,
		ExitDate:      exitDate,
		DocumentValue: documentValue,
		PayTypeId:     payTypeId,
		Discount:      discount,
		Reduction:     reduction,
		ProductValue:  productValue,
		FreightType:   freightType,
		FreightValue:  freightValue,
		SafeValue:     safeValue,
		ExpenseValue:  expenseValue,
		BaseIcms:      baseIcms,
		ValueIcms:     valueIcms,
		BaseIcmsSt:    baseIcmsSt,
		ValueIcmsSt:   valueIcmsSt,
		ValueIpi:      valueIpi,
		ValuePis:      valuePis,
		ValueCofins:   valueCofins,
		ValuePisSt:    valuePisSt,
		ValueCofinsSt: valueCofinsSt,
		Imported:      imported,
		Origined:      origined,
	}
	return NewDocumentEntity(model)
}

func NewDocumentEntity(entity Document) (*Document, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
