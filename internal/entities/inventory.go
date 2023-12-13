package entities

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	Id        float64 `gorm:"primaryKey"`
	TenantId  float64
	Tenant    Tenant `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	Name      string `gorm:"size:150"`
	StartDate time.Time
	EndDate   time.Time
	Processed bool
	Cloused   bool
}

func (c *Inventory) Validate() error {
	if c.Name == "" {
		return errors.New("the name is required")
	}
	if c.TenantId == 0 {
		return errors.New("the tenantId is inválid")
	}
	return nil
}

func (c *Inventory) SetProcessed() error {
	c.Processed = true
	return nil
}

func NewInventory(tenantId float64, name string, document string) (*Inventory, error) {
	model := Inventory{
		Id:        0,
		Name:      name,
		TenantId:  tenantId,
		StartDate: time.Now(),
		Processed: false,
	}
	return NewInventoryEntity(model)
}

func NewInventoryEntity(entity Inventory) (*Inventory, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
