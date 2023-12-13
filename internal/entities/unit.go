package entities

import (
	"errors"

	"github.com/EdimarRibeiro/inventory/internal/utils"
	"gorm.io/gorm"
)

/*0190*/
type Unit struct {
	gorm.Model
	Id          string `gorm:"primaryKey;size:6"`
	TenantId    float64
	Tenant      Tenant `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	Description string `gorm:"size:50"`
}

func (c *Unit) Validate() error {
	if c.Id == "" {
		return errors.New("the unit is required")
	}
	if c.TenantId == 0 {
		return errors.New("the tenantId is required")
	}
	if len(c.Id) < 2 {
		return errors.New("the min value unit is inválid! " + c.Id)
	}

	if len(c.Id) > 6 {
		return errors.New("the max value unit is inválid! " + c.Id)
	}
	return nil
}

func NewUnit(tenantId float64, id string, description string) (*Unit, error) {
	model := Unit{
		Id:          id,
		Description: description,
	}
	return NewUnitEntity(model)
}

func CreateUnit(tenantId float64, line string) (*Unit, error) {
	var err error = nil
	unit := Unit{}
	unit.Id, err = utils.CopyText(line, 2)
	if err != nil {
		return nil, err
	}
	unit.TenantId, err = tenantId, nil
	unit.Description, err = utils.CopyText(line, 3)
	if err != nil {
		return nil, err
	}
	return NewUnitEntity(unit)
}

func NewUnitEntity(entity Unit) (*Unit, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
