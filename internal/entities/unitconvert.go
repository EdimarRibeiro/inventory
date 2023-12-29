package entities

import (
	"errors"

	"github.com/EdimarRibeiro/inventory/internal/utils"
)

/*0220*/
type UnitConvert struct {
	UnitId           string  `gorm:"primaryKey;size:6"`
	Unit             Unit    `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	ProductId        uint64  `gorm:"primaryKey"`
	Product          Product `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	TenantId         uint64  `gorm:"primaryKey"`
	Tenant           Tenant  `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	ConversionFactor float64 `gorm:"type:decimal (18,6)"`
	BarCode          string  `gorm:"primaryKey;size:60"`
}

func (c *UnitConvert) Validate() error {
	if c.UnitId == "" {
		return errors.New("the unitConvert is required")
	}

	if len(c.UnitId) < 2 {
		return errors.New("the min value unitId is inválid! " + c.UnitId)
	}

	if len(c.UnitId) > 6 {
		return errors.New("the max value unitId is inválid! " + c.UnitId)
	}

	if c.ProductId == 0 {
		return errors.New("the productId is required")
	}

	if c.TenantId == 0 {
		return errors.New("the tenant is required")
	}

	if c.ConversionFactor == 0 {
		return errors.New("the conversionFactor is required")
	}

	if c.BarCode == "" {
		return errors.New("the barCode is required")
	}
	return nil
}

func NewUnitConvert(unitId string, productId uint64, tenantId uint64, conversionFactor float64, barCode string) (*UnitConvert, error) {
	model := UnitConvert{
		UnitId:           unitId,
		ProductId:        productId,
		TenantId:         tenantId,
		ConversionFactor: conversionFactor,
		BarCode:          barCode,
	}
	return NewUnitConvertEntity(model)
}

func CreateUnitConvert(tenantId uint64, productId uint64, line string) (*UnitConvert, error) {
	var err error = nil
	model := UnitConvert{}
	model.UnitId, err = utils.CopyText(line, 2)
	if err != nil {
		return nil, err
	}
	model.BarCode, err = utils.CopyText(line, 4)
	if err != nil {
		return nil, err
	}
	model.ProductId, err = productId, nil
	if err != nil {
		return nil, err
	}
	model.TenantId, err = tenantId, nil
	model.ConversionFactor, err = utils.CopyTextFloat(line, 3, 6)
	if err != nil {
		return nil, err
	}
	return NewUnitConvertEntity(model)
}

func NewUnitConvertEntity(entity UnitConvert) (*UnitConvert, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
