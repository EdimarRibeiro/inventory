package entities

import (
	"errors"
	"time"
)

type Inventory struct {
	Id        uint64 `gorm:"primaryKey"`
	TenantId  uint64
	Tenant    Tenant    `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	Name      string    `gorm:"size:150"`
	StartDate time.Time `gorm:"datetime"`
	EndDate   time.Time `gorm:"datetime"`
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

func NewInventory(tenantId uint64, name string) (*Inventory, error) {
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
