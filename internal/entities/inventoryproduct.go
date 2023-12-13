package entities

import (
	"errors"
	"time"

	"github.com/EdimarRibeiro/inventory/internal/utils"
	"gorm.io/gorm"
)

/* H005 e H010 */
type InventoryProduct struct {
	gorm.Model
	InventoryId     float64   `gorm:"primaryKey"`
	Inventory       Inventory `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	ProductId       float64   `gorm:"primaryKey"`
	Product         Product   `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	OriginCode      string    `gorm:"size:60"`
	Date            time.Time
	UnitId          string `gorm:"size:6"`
	Unit            Unit   `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	Quantity        float64
	Value           float64
	ValueTotal      float64
	PossessionCode  string `gorm:"size:1"`
	ParticipantId   float64
	Complement      string `gorm:"size:250"`
	AccountingCode  string `gorm:"size:50"`
	ValueIr         float64
	InputQuantity   float64
	OutputQuantity  float64
	BalanceQuantity float64
}

func (c *InventoryProduct) Validate() error {
	if c.InventoryId == 0 {
		return errors.New("the inventoryId is required")
	}
	if c.OriginCode == "" {
		return errors.New("the originCode is required")
	}
	if c.ProductId == 0 {
		return errors.New("the productId is required")
	}
	if c.UnitId == "" {
		return errors.New("the unitId is required")
	}
	if len(c.UnitId) < 2 {
		return errors.New("the min value unitId is inválid! " + c.UnitId)
	}
	if len(c.UnitId) > 6 {
		return errors.New("the max value unitId is inválid! " + c.UnitId)
	}
	if c.Quantity < 0 {
		return errors.New("the quantity is inválid")
	}
	if c.Value < 0 {
		return errors.New("the value is inválid")
	}
	return nil
}

func NewInventoryProduct(inventoryId float64, productId float64, originCode string, date time.Time, unitId string, quantity float64, value float64, valueTotal float64, possessionCode string, participantId float64, complement string, accountingCode string, valueIr float64) (*InventoryProduct, error) {
	model := InventoryProduct{
		InventoryId:    inventoryId,
		ProductId:      productId,
		OriginCode:     originCode,
		Date:           date,
		UnitId:         unitId,
		Quantity:       quantity,
		Value:          value,
		ValueTotal:     valueTotal,
		PossessionCode: possessionCode,
		ParticipantId:  participantId,
		Complement:     complement,
		AccountingCode: accountingCode,
		ValueIr:        valueIr,
	}
	return NewInventoryProductEntity(model)
}

func CreateInventoryProduct(inventoryId float64, date time.Time, line string) (*InventoryProduct, error) {
	var err error = nil
	inventoryProduct := InventoryProduct{}
	inventoryProduct.InventoryId, err = inventoryId, nil
	inventoryProduct.ProductId, err = GetProductId(utils.CopyText(line, 2))
	if err != nil {
		return nil, err
	}
	inventoryProduct.OriginCode, err = utils.CopyText(line, 2)
	if err != nil {
		return nil, err
	}
	inventoryProduct.Date, err = date, nil
	if err != nil {
		return nil, err
	}
	inventoryProduct.UnitId, err = utils.CopyText(line, 3)
	if err != nil {
		return nil, err
	}
	inventoryProduct.Quantity, err = utils.CopyTextFloat(line, 4, 3)
	if err != nil {
		return nil, err
	}
	inventoryProduct.Value, err = utils.CopyTextFloat(line, 5, 6)
	if err != nil {
		return nil, err
	}
	inventoryProduct.ValueTotal, err = utils.CopyTextFloat(line, 6, 2)
	if err != nil {
		return nil, err
	}
	inventoryProduct.PossessionCode, err = utils.CopyText(line, 7)
	if err != nil {
		return nil, err
	}
	inventoryProduct.ParticipantId, err = GetParticipantId(utils.CopyText(line, 8))
	if err != nil {
		return nil, err
	}
	inventoryProduct.Complement, err = utils.CopyText(line, 9)
	if err != nil {
		return nil, err
	}
	inventoryProduct.AccountingCode, err = utils.CopyText(line, 10)
	if err != nil {
		return nil, err
	}
	inventoryProduct.ValueIr, err = utils.CopyTextFloat(line, 11, 2)
	if err != nil {
		return nil, err
	}
	inventoryProduct.BalanceQuantity, err = 0, nil
	inventoryProduct.InputQuantity, err = 0, nil
	inventoryProduct.OutputQuantity, err = 0, nil

	return NewInventoryProductEntity(inventoryProduct)
}

func NewInventoryProductEntity(entity InventoryProduct) (*InventoryProduct, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (i *InventoryProduct) SetInputQuantity(quantity float64) error {
	if quantity < 0 {
		return errors.New("input quantity is invalid value")
	}
	i.InputQuantity = quantity
	return nil
}

func (i *InventoryProduct) SetOutputQuantity(quantity float64) error {
	if quantity < 0 {
		return errors.New("output quantity is invalid value")
	}
	i.InputQuantity = quantity
	return nil
}

func (i *InventoryProduct) CalculateBalanceQuantity() error {
	i.BalanceQuantity = i.Quantity + i.InputQuantity - i.OutputQuantity
	return nil
}
