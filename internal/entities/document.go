package entities

import (
	"errors"
	"time"

	"github.com/EdimarRibeiro/inventory/internal/utils"
	"gorm.io/gorm"
)

/*C100*/
type Document struct {
	gorm.Model
	Id            float64   `gorm:"primaryKey;autoIncrement:true"`
	InventoryId   float64   `gorm:"index:idx_Inventory"`
	Inventory     Inventory `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	ParticipantId float64
	Participant   Participant `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	OperationId   string      `gorm:"size:1"`
	EmitentTypeId string      `gorm:"size:1"`
	ModelId       string      `gorm:"size:2"`
	StatusId      string      `gorm:"size:2"`
	Serie         string      `gorm:"size:3"`
	Number        string      `gorm:"size:9"`
	DocumentKey   string      `gorm:"size:44"`
	EmitentDate   time.Time
	ExitDate      time.Time
	DocumentValue float64
	PayTypeId     string `gorm:"size:1"`
	Discount      float64
	Reduction     float64
	ProductValue  float64
	FreightType   string `gorm:"size:1"`
	FreightValue  float64
	SafeValue     float64
	ExpenseValue  float64
	BaseIcms      float64
	ValueIcms     float64
	BaseIcmsSt    float64
	ValueIcmsSt   float64
	ValueIpi      float64
	ValuePis      float64
	ValueCofins   float64
	ValuePisSt    float64
	ValueCofinsSt float64
	Processed     bool
	Imported      bool
	Cloused       bool
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

func CreateDocument(InventoryId float64, line string) (*Document, error) {
	var err error = nil
	document := Document{}

	document.Id, err = 0, nil
	document.InventoryId, err = InventoryId, nil
	document.OperationId, err = utils.CopyText(line, 2)
	if err != nil {
		return nil, err
	}
	document.EmitentTypeId, err = utils.CopyText(line, 3)
	if err != nil {
		return nil, err
	}
	document.ParticipantId, err = GetParticipantId(utils.CopyText(line, 4))
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
	document.Processed, err = true, nil
	document.Imported, err = false, nil
	document.Cloused, err = false, nil
	return NewDocumentEntity(document)
}

func NewDocument(inventoryId float64, participantId float64, operationId string, emitentTypeId string, modelId string, statusId string, serie string, number string, documentKey string, emitentDate time.Time, exitDate time.Time, documentValue float64, payTypeId string, discount float64, reduction float64, productValue float64, freightType string, freightValue float64, safeValue float64, expenseValue float64, baseIcms float64, valueIcms float64, baseIcmsSt float64, valueIcmsSt float64, valueIpi float64, valuePis float64, valueCofins float64, valuePisSt float64, valueCofinsSt float64, imported bool) (*Document, error) {
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
