package entities

import (
	"errors"
	"time"

	"github.com/EdimarRibeiro/inventory/internal/utils"
)

/* H005 e H010 */
type InventoryProduct struct {
	InventoryId     uint64    `gorm:"primaryKey" json:"inventoryId"`
	Inventory       Inventory `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	ProductId       uint64    `gorm:"primaryKey" json:"productId"`
	Product         Product   `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"product"`
	OriginCode      string    `gorm:"size:60" json:"originCode"`
	Date            time.Time `gorm:"datetime" json:"date"`
	UnitId          string    `gorm:"size:6" json:"unitId"`
	Unit            Unit      `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	Quantity        float64   `gorm:"type:decimal (18,3)" json:"quantity"`
	Value           float64   `gorm:"type:decimal (18,6)" json:"value"`
	ValueTotal      float64   `gorm:"type:decimal (18,2)" json:"valueTotal"`
	PossessionCode  string    `gorm:"size:1" json:"possessionCode"`
	ParticipantId   *uint64   `gorm:"default:null" json:"participantId"`
	Complement      string    `gorm:"size:250" json:"complement"`
	AccountingCode  string    `gorm:"size:50" json:"accountingCode"`
	ValueIr         float64   `gorm:"type:decimal (12,2)" json:"valueIr"`
	InputQuantity   float64   `gorm:"type:decimal (18,3)" json:"inputQuantity"`
	OutputQuantity  float64   `gorm:"type:decimal (18,3)" json:"outputQuantity"`
	BalanceQuantity float64   `gorm:"type:decimal (18,3)" json:"balanceQuantity"`
}

func (c *InventoryProduct) Validate() error {
	if c.InventoryId == 0 {
		return errors.New("the inventoryId is required")
	}
	if c.OriginCode == "" {
		return errors.New("the product originCode is required")
	}
	if c.ProductId == 0 {
		return errors.New("the productId is required")
	}
	if c.UnitId == "" {
		return errors.New("the unitId is required")
	}
	if len(c.UnitId) < 2 {
		return errors.New("the min value unitId is invalid! " + c.UnitId)
	}
	if len(c.UnitId) > 6 {
		return errors.New("the max value unitId is invalid! " + c.UnitId)
	}
	if c.Quantity < 0 {
		return errors.New("the quantity is invalid")
	}
	if c.Value < 0 {
		return errors.New("the value is invalid")
	}
	return nil
}

func NewInventoryProduct(inventoryId uint64, productId uint64, originCode string, date time.Time, unitId string, quantity float64, value float64, valueTotal float64, possessionCode string, participantId uint64, complement string, accountingCode string, valueIr float64) (*InventoryProduct, error) {
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
		Complement:     complement,
		AccountingCode: accountingCode,
		ValueIr:        valueIr,
		ParticipantId:  nil,
	}
	if participantId > 0 {
		model.ParticipantId = &participantId
	}
	return NewInventoryProductEntity(model)
}

func CreateInventoryProduct(inventoryId uint64, productId uint64, participantId uint64, date time.Time, line string) (*InventoryProduct, error) {
	var err error = nil
	inventoryProduct := InventoryProduct{}
	inventoryProduct.InventoryId, err = inventoryId, nil
	inventoryProduct.ProductId, err = productId, nil
	inventoryProduct.ParticipantId = nil
	if participantId > 0 {
		inventoryProduct.ParticipantId, err = &participantId, nil
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
	i.OutputQuantity = quantity
	return nil
}

func (i *InventoryProduct) CalculateBalanceQuantity() error {
	i.BalanceQuantity = i.Quantity + i.InputQuantity - i.OutputQuantity
	return nil
}
